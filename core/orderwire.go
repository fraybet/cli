package core

import (
	"fmt"
	"math/big"
	"strings"
)

// OrderJSON is the wire representation of an Order: big integers as decimal
// strings, addresses as 0x-hex, side/signatureType as small ints. Shared by the
// CLI (which builds + signs orders) and the API (which parses them), so both
// agree on the exact field set the signature covers.
type OrderJSON struct {
	Salt          string `json:"salt"`
	Maker         string `json:"maker"`
	Signer        string `json:"signer"`
	Taker         string `json:"taker"`
	TokenID       string `json:"tokenId"`
	MakerAmount   string `json:"makerAmount"`
	TakerAmount   string `json:"takerAmount"`
	Expiration    string `json:"expiration"`
	Nonce         string `json:"nonce"`
	FeeRateBps    string `json:"feeRateBps"`
	Side          uint8  `json:"side"`
	SignatureType uint8  `json:"signatureType"`
}

func decBig(s string) (*big.Int, error) {
	x, ok := new(big.Int).SetString(strings.TrimSpace(s), 10)
	if !ok {
		return nil, fmt.Errorf("core: not a decimal integer: %q", s)
	}
	return x, nil
}

func wireAddr(s string) (Address, error) {
	if strings.TrimSpace(s) == "" {
		return Address{}, nil // empty => zero address (public order taker)
	}
	return HexToAddress(s)
}

// ToOrder converts the wire form to an Order.
func (j OrderJSON) ToOrder() (Order, error) {
	var o Order
	var err error
	if o.Maker, err = wireAddr(j.Maker); err != nil {
		return o, fmt.Errorf("maker: %w", err)
	}
	if o.Signer, err = wireAddr(j.Signer); err != nil {
		return o, fmt.Errorf("signer: %w", err)
	}
	if o.Taker, err = wireAddr(j.Taker); err != nil {
		return o, fmt.Errorf("taker: %w", err)
	}
	for _, f := range []struct {
		dst **big.Int
		src string
		nm  string
	}{
		{&o.Salt, j.Salt, "salt"},
		{&o.TokenID, j.TokenID, "tokenId"},
		{&o.MakerAmount, j.MakerAmount, "makerAmount"},
		{&o.TakerAmount, j.TakerAmount, "takerAmount"},
		{&o.Expiration, j.Expiration, "expiration"},
		{&o.Nonce, j.Nonce, "nonce"},
		{&o.FeeRateBps, j.FeeRateBps, "feeRateBps"},
	} {
		v, perr := decBig(f.src)
		if perr != nil {
			return o, fmt.Errorf("%s: %w", f.nm, perr)
		}
		*f.dst = v
	}
	o.Side = OrderSide(j.Side)
	o.SignatureType = OrderSignatureType(j.SignatureType)
	return o, nil
}

// OrderToJSON converts an Order to its wire form.
func OrderToJSON(o Order) OrderJSON {
	return OrderJSON{
		Salt:          o.Salt.String(),
		Maker:         o.Maker.Hex(),
		Signer:        o.Signer.Hex(),
		Taker:         o.Taker.Hex(),
		TokenID:       o.TokenID.String(),
		MakerAmount:   o.MakerAmount.String(),
		TakerAmount:   o.TakerAmount.String(),
		Expiration:    o.Expiration.String(),
		Nonce:         o.Nonce.String(),
		FeeRateBps:    o.FeeRateBps.String(),
		Side:          uint8(o.Side),
		SignatureType: uint8(o.SignatureType),
	}
}
