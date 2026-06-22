// Package chain is the thin layer between Go and Base: RPC profiles, a client
// wrapper (balances, gas estimation, eth_call simulation), an unsigned-tx /
// calldata builder, and EIP-712 signing helpers. Consistent with the
// non-custodial rule, the builder produces UNSIGNED transactions; signing is a
// separate, explicit step the caller controls (the MCP server never signs).
package chain

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/fraybet/cli/core"
)

// Profile is the connection + contract context for one environment
// (e.g. base-sepolia, base-mainnet).
type Profile struct {
	Name      string
	ChainID   *big.Int
	RPCURL    string
	Contracts map[string]core.Address
}

// Contract resolves a logical contract name (e.g. "BetEscrowFactory") to its
// deployed address for this profile.
func (p Profile) Contract(name string) (core.Address, error) {
	a, ok := p.Contracts[name]
	if !ok {
		return core.Address{}, fmt.Errorf("chain: contract %q not configured for profile %q", name, p.Name)
	}
	return a, nil
}

// contractsFile is the on-disk shape of a per-profile contracts JSON
// (design §17.3 contracts_file = "./contracts.base.json").
type contractsFile struct {
	ChainID   int64             `json:"chainId"`
	RPCURL    string            `json:"rpcUrl"`
	Contracts map[string]string `json:"contracts"`
}

// LoadProfile parses a contracts JSON file into a Profile.
func LoadProfile(name string, data []byte) (Profile, error) {
	var cf contractsFile
	if err := json.Unmarshal(data, &cf); err != nil {
		return Profile{}, fmt.Errorf("chain: parse contracts file: %w", err)
	}
	if cf.ChainID == 0 {
		return Profile{}, fmt.Errorf("chain: chainId is required")
	}
	contracts := make(map[string]core.Address, len(cf.Contracts))
	for n, hexAddr := range cf.Contracts {
		a, err := core.HexToAddress(hexAddr)
		if err != nil {
			return Profile{}, fmt.Errorf("chain: contract %q: %w", n, err)
		}
		contracts[n] = a
	}
	return Profile{
		Name:      name,
		ChainID:   big.NewInt(cf.ChainID),
		RPCURL:    cf.RPCURL,
		Contracts: contracts,
	}, nil
}
