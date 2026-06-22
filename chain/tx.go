package chain

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/fraybet/cli/core"
)

// UnsignedTx is a prepared-but-unsigned transaction. This is exactly what the
// CLI and MCP server hand back to a wallet/agent to sign and broadcast — the
// platform never signs or holds keys (design §16.4). All numeric fields
// serialize as 0x-hex so the payload is wallet-ready.
type UnsignedTx struct {
	From    core.Address
	To      core.Address
	Value   *big.Int // wei
	Data    []byte   // calldata
	Gas     uint64
	Nonce   uint64
	ChainID *big.Int
}

// BuildCall constructs an unsigned contract call. Gas/Nonce are filled in later
// by the client (EstimateGas / pending nonce) or by the signer's wallet.
func BuildCall(from, to core.Address, value *big.Int, data []byte) UnsignedTx {
	if value == nil {
		value = big.NewInt(0)
	}
	return UnsignedTx{From: from, To: to, Value: value, Data: data}
}

// txJSON is the wire form: every field 0x-hex so a wallet can consume it.
type txJSON struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Value   string `json:"value"`
	Data    string `json:"data"`
	Gas     string `json:"gas"`
	Nonce   string `json:"nonce"`
	ChainID string `json:"chainId"`
}

func hexBig(x *big.Int) string {
	if x == nil {
		x = big.NewInt(0)
	}
	return "0x" + x.Text(16)
}

func hexUint(x uint64) string { return "0x" + new(big.Int).SetUint64(x).Text(16) }

func (t UnsignedTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(txJSON{
		From:    t.From.Hex(),
		To:      t.To.Hex(),
		Value:   hexBig(t.Value),
		Data:    "0x" + hex.EncodeToString(t.Data),
		Gas:     hexUint(t.Gas),
		Nonce:   hexUint(t.Nonce),
		ChainID: hexBig(t.ChainID),
	})
}

// UnmarshalJSON parses the wire form back into an UnsignedTx, so an agent can
// round-trip the payload it receives before signing.
func (t *UnsignedTx) UnmarshalJSON(data []byte) error {
	var j txJSON
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	from, err := core.HexToAddress(j.From)
	if err != nil {
		return fmt.Errorf("chain: tx.from: %w", err)
	}
	to, err := core.HexToAddress(j.To)
	if err != nil {
		return fmt.Errorf("chain: tx.to: %w", err)
	}
	value, err := parseHexBig(j.Value)
	if err != nil {
		return fmt.Errorf("chain: tx.value: %w", err)
	}
	chainID, err := parseHexBig(j.ChainID)
	if err != nil {
		return fmt.Errorf("chain: tx.chainId: %w", err)
	}
	gas, err := parseHexUint(j.Gas)
	if err != nil {
		return fmt.Errorf("chain: tx.gas: %w", err)
	}
	nonce, err := parseHexUint(j.Nonce)
	if err != nil {
		return fmt.Errorf("chain: tx.nonce: %w", err)
	}
	raw, err := parseHexBytes(j.Data)
	if err != nil {
		return fmt.Errorf("chain: tx.data: %w", err)
	}
	*t = UnsignedTx{From: from, To: to, Value: value, Data: raw, Gas: gas, Nonce: nonce, ChainID: chainID}
	return nil
}

func parseHexBig(s string) (*big.Int, error) {
	s = strings.TrimPrefix(strings.TrimSpace(s), "0x")
	if s == "" {
		return big.NewInt(0), nil
	}
	n, ok := new(big.Int).SetString(s, 16)
	if !ok {
		return nil, fmt.Errorf("invalid hex quantity %q", s)
	}
	return n, nil
}

func parseHexUint(s string) (uint64, error) {
	n, err := parseHexBig(s)
	if err != nil {
		return 0, err
	}
	return n.Uint64(), nil
}

func parseHexBytes(s string) ([]byte, error) {
	s = strings.TrimPrefix(strings.TrimSpace(s), "0x")
	if s == "" {
		return nil, nil
	}
	return hex.DecodeString(s)
}
