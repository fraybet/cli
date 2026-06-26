package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/fraybet/cli/core"
)

const orderUsage = `fray order <command> [flags]

commands:
  place        sign and submit a limit order to a market's order book
  cancel       cancel a resting order you placed (instant, gas-free)
  book         show a market's order book depth

Examples:
  fray order place --market 0xCID --token 1234… --side buy --price 0.55 --size 100000000 --wallet me
  fray order cancel --market 0xCID --digest 0xHASH --wallet me
  fray order book --market 0xCID
`

func runOrder(args []string, out io.Writer) error {
	if len(args) == 0 {
		fmt.Fprint(out, orderUsage)
		return nil
	}
	rest := args[1:]
	switch args[0] {
	case "place":
		return orderPlace(rest, out)
	case "cancel":
		return orderCancel(rest, out)
	case "book":
		return orderBook(rest, out)
	case "help", "-h", "--help":
		fmt.Fprint(out, orderUsage)
		return nil
	default:
		return fmt.Errorf("order: unknown command %q", args[0])
	}
}

// orderConfig is GET /orderbook/config — the exchange address + chain the order
// EIP-712 domain binds to.
type orderConfig struct {
	Exchange string `json:"exchange"`
	ChainID  int64  `json:"chainId"`
}

func fetchOrderConfig(base string) (orderConfig, error) {
	var c orderConfig
	resp, err := http.Get(strings.TrimRight(base, "/") + "/orderbook/config")
	if err != nil {
		return c, fmt.Errorf("order: config request: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return c, fmt.Errorf("order: config failed (%d): %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	if err := json.Unmarshal(body, &c); err != nil || c.Exchange == "" {
		return c, fmt.Errorf("order: bad config response")
	}
	return c, nil
}

func orderPlace(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("place", flag.ContinueOnError)
	fs.SetOutput(out)
	market := fs.String("market", "", "market conditionId (0x…)")
	token := fs.String("token", "", "outcome token id to trade (decimal)")
	side := fs.String("side", "", "buy | sell")
	price := fs.String("price", "", "limit price, 0..1 (collateral per outcome token)")
	size := fs.String("size", "", "order size in outcome-token base units (6dp)")
	expiry := fs.Int64("expiry", 0, "unix expiry (0 = good-till-cancelled)")
	baseURL := fs.String("base-url", "", "site base URL (env FRAY_BASE_URL; default https://fray.bet)")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	for _, req := range []struct {
		v, nm string
	}{{*market, "--market"}, {*token, "--token"}, {*side, "--side"}, {*price, "--price"}, {*size, "--size"}} {
		if strings.TrimSpace(req.v) == "" {
			return fmt.Errorf("place: %s is required", req.nm)
		}
	}
	var os core.OrderSide
	switch strings.ToLower(*side) {
	case "buy":
		os = core.OrderBuy
	case "sell":
		os = core.OrderSell
	default:
		return fmt.Errorf("place: --side must be buy or sell")
	}

	w, err := sf.resolve()
	if err != nil {
		return err
	}
	base := resolveBaseURL(*baseURL)
	cfg, err := fetchOrderConfig(base)
	if err != nil {
		return err
	}
	exch, err := core.HexToAddress(cfg.Exchange)
	if err != nil {
		return fmt.Errorf("place: bad exchange address: %w", err)
	}

	tokenID, ok := new(big.Int).SetString(strings.TrimSpace(*token), 10)
	if !ok {
		return fmt.Errorf("place: --token must be a decimal token id")
	}
	sizeBase, ok := new(big.Int).SetString(strings.TrimSpace(*size), 10)
	if !ok || sizeBase.Sign() <= 0 {
		return fmt.Errorf("place: --size must be a positive integer (base units)")
	}
	priceMicro, err := priceToMicro(*price)
	if err != nil {
		return err
	}
	maker, taker := orderAmounts(os, sizeBase, priceMicro)
	if maker.Sign() <= 0 || taker.Sign() <= 0 {
		return fmt.Errorf("place: price×size rounds to zero")
	}

	salt, err := randomSalt()
	if err != nil {
		return err
	}
	o := core.Order{
		Salt: salt, Maker: w.Address, Signer: w.Address, Taker: core.Address{},
		TokenID: tokenID, MakerAmount: maker, TakerAmount: taker,
		Expiration: big.NewInt(*expiry), Nonce: big.NewInt(0),
		FeeRateBps: big.NewInt(0), Side: os, SignatureType: core.OrderSigEOA,
	}

	domain := core.ExchangeDomain(big.NewInt(cfg.ChainID), exch)
	digest := o.Digest(domain)
	sig, err := w.Sign(digest)
	if err != nil {
		return fmt.Errorf("place: sign: %w", err)
	}

	res, err := postOrder(base, *market, core.OrderToJSON(o), sig)
	if err != nil {
		return err
	}
	if res.TxHash != nil {
		fmt.Fprintf(out, "order %s — %s (tx %s)\n", res.Digest, res.Status, *res.TxHash)
	} else {
		fmt.Fprintf(out, "order %s — %s\n", res.Digest, res.Status)
	}
	return nil
}

type placeResp struct {
	Digest string  `json:"digest"`
	Status string  `json:"status"`
	TxHash *string `json:"txHash"`
}

func postOrder(base, market string, order core.OrderJSON, sig []byte) (placeResp, error) {
	var res placeResp
	body, _ := json.Marshal(map[string]any{
		"market": market, "order": order, "signature": "0x" + hex.EncodeToString(sig),
	})
	resp, err := http.Post(strings.TrimRight(base, "/")+"/orders", "application/json", bytes.NewReader(body))
	if err != nil {
		return res, fmt.Errorf("place: request: %w", err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf("place: server rejected (%d): %s", resp.StatusCode, strings.TrimSpace(string(raw)))
	}
	_ = json.Unmarshal(raw, &res)
	return res, nil
}

func orderCancel(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("cancel", flag.ContinueOnError)
	fs.SetOutput(out)
	market := fs.String("market", "", "market conditionId (0x…)")
	digestHex := fs.String("digest", "", "order digest to cancel (0x…)")
	baseURL := fs.String("base-url", "", "site base URL (env FRAY_BASE_URL; default https://fray.bet)")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*market) == "" || strings.TrimSpace(*digestHex) == "" {
		return fmt.Errorf("cancel: --market and --digest are required")
	}
	digest, err := hexToHash(*digestHex)
	if err != nil {
		return fmt.Errorf("cancel: bad digest: %w", err)
	}
	w, err := sf.resolve()
	if err != nil {
		return err
	}
	base := resolveBaseURL(*baseURL)

	expires := time.Now().Add(5 * time.Minute).Unix()
	sig, err := w.Sign(core.PersonalDigest(core.CancelOrderMessage(digest, expires)))
	if err != nil {
		return fmt.Errorf("cancel: sign: %w", err)
	}
	body, _ := json.Marshal(map[string]any{
		"market": *market, "expires": expires, "signature": "0x" + hex.EncodeToString(sig),
	})
	req, _ := http.NewRequest(http.MethodDelete, strings.TrimRight(base, "/")+"/orders/"+digest.Hex(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("cancel: request: %w", err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cancel: server rejected (%d): %s", resp.StatusCode, strings.TrimSpace(string(raw)))
	}
	fmt.Fprintf(out, "cancelled %s\n", digest.Hex())
	return nil
}

type bookLevel struct {
	Token string `json:"token"`
	Side  string `json:"side"`
	Price string `json:"price"`
	Size  string `json:"size"`
}

func orderBook(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("book", flag.ContinueOnError)
	fs.SetOutput(out)
	market := fs.String("market", "", "market conditionId (0x…)")
	baseURL := fs.String("base-url", "", "site base URL (env FRAY_BASE_URL; default https://fray.bet)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*market) == "" {
		return fmt.Errorf("book: --market is required")
	}
	base := resolveBaseURL(*baseURL)
	resp, err := http.Get(strings.TrimRight(base, "/") + "/markets/" + *market + "/book")
	if err != nil {
		return fmt.Errorf("book: request: %w", err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("book: server rejected (%d): %s", resp.StatusCode, strings.TrimSpace(string(raw)))
	}
	var bv struct {
		Levels []bookLevel `json:"levels"`
	}
	_ = json.Unmarshal(raw, &bv)
	if len(bv.Levels) == 0 {
		fmt.Fprintln(out, "(empty book)")
		return nil
	}
	fmt.Fprintf(out, "%-8s %-5s %-14s %s\n", "SIDE", "TOK", "PRICE", "SIZE")
	for _, l := range bv.Levels {
		fmt.Fprintf(out, "%-8s %-5s %-14s %s\n", l.Side, shortTok(l.Token), price1e18ToDecimal(l.Price), l.Size)
	}
	return nil
}

// orderAmounts converts a limit (side, size, price) into the order's maker/taker
// amounts. collateral = size × price (both at USDC's 6dp scale). A BUY pays
// collateral to receive tokens; a SELL gives tokens to receive collateral.
func orderAmounts(side core.OrderSide, sizeBase, priceMicro *big.Int) (maker, taker *big.Int) {
	collateral := new(big.Int).Div(new(big.Int).Mul(sizeBase, priceMicro), big.NewInt(1_000_000))
	if side == core.OrderBuy {
		return collateral, sizeBase
	}
	return sizeBase, collateral
}

// priceToMicro parses a 0..1 decimal price into micro-units (×1e6), USDC's scale.
func priceToMicro(s string) (*big.Int, error) {
	s = strings.TrimSpace(s)
	neg := strings.HasPrefix(s, "-")
	if neg {
		return nil, fmt.Errorf("place: --price must be between 0 and 1")
	}
	whole, frac, _ := strings.Cut(s, ".")
	if len(frac) > 6 {
		frac = frac[:6]
	}
	for len(frac) < 6 {
		frac += "0"
	}
	combined := whole + frac
	v, ok := new(big.Int).SetString(combined, 10)
	if !ok {
		return nil, fmt.Errorf("place: --price must be a decimal like 0.55")
	}
	if v.Sign() <= 0 || v.Cmp(big.NewInt(1_000_000)) > 0 {
		return nil, fmt.Errorf("place: --price must be between 0 and 1")
	}
	return v, nil
}

// price1e18ToDecimal renders a 1e18-scaled price as a trimmed 0..1 decimal.
func price1e18ToDecimal(s string) string {
	v, ok := new(big.Int).SetString(strings.TrimSpace(s), 10)
	if !ok {
		return s
	}
	q, r := new(big.Int).QuoRem(v, big.NewInt(1_000_000_000_000_000_000), new(big.Int))
	if r.Sign() == 0 {
		return q.String()
	}
	frac := fmt.Sprintf("%018d", r)
	frac = strings.TrimRight(frac, "0")
	return q.String() + "." + frac
}

func shortTok(t string) string {
	if len(t) > 6 {
		return t[:6] + "…"
	}
	return t
}

func randomSalt() (*big.Int, error) {
	var b [32]byte
	if _, err := rand.Read(b[:]); err != nil {
		return nil, fmt.Errorf("order: salt: %w", err)
	}
	return new(big.Int).SetBytes(b[:]), nil
}

func hexToHash(s string) (core.Hash32, error) {
	b, err := hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(s), "0x"))
	if err != nil {
		return core.Hash32{}, err
	}
	if len(b) != 32 {
		return core.Hash32{}, fmt.Errorf("expected a 32-byte hash")
	}
	var h core.Hash32
	copy(h[:], b)
	return h, nil
}
