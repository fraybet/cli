// Package txbuild turns domain actions into UNSIGNED transactions (calldata) by
// packing contract calls through the abigen-generated ABIs. This is the bridge
// from internal/bets + the contract bindings to internal/chain.UnsignedTx — the
// non-custodial pattern: the CLI/MCP hand back unsigned txs for the agent to
// sign, the platform never signs or broadcasts.
package txbuild

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/fraybet/cli/bets"
	"github.com/fraybet/cli/bindings/betescrow"
	"github.com/fraybet/cli/bindings/betescrowfactory"
	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
)

func ethAddr(a core.Address) ethcommon.Address { return ethcommon.BytesToAddress(a[:]) }

// CreateBet builds the unsigned tx that deploys a bet via the factory.
func CreateBet(from, factory core.Address, d bets.Draft) (chain.UnsignedTx, error) {
	parsed, err := betescrowfactory.BetEscrowFactoryMetaData.GetAbi()
	if err != nil {
		return chain.UnsignedTx{}, fmt.Errorf("txbuild: factory abi: %w", err)
	}
	th := d.TermsHash()
	terms := betescrowfactory.BetEscrowTerms{
		YesAgent:        ethAddr(d.Terms.YesAgent),
		NoAgent:         ethAddr(d.Terms.NoAgent),
		Arbiter:         ethAddr(d.Terms.Arbiter),
		Token:           ethAddr(d.Terms.CollateralToken),
		YesStake:        d.Terms.YesStake,
		NoStake:         d.Terms.NoStake,
		ClaimDeadline:   d.Terms.ClaimDeadline,
		ChallengeWindow: d.Terms.ChallengeWindow,
		TermsHash:       [32]byte(th),
		Visibility:      uint8(d.Visibility),
	}
	data, err := parsed.Pack("create", terms)
	if err != nil {
		return chain.UnsignedTx{}, fmt.Errorf("txbuild: pack create: %w", err)
	}
	return chain.BuildCall(from, factory, nil, data), nil
}

// Approve builds an ERC-20 approve tx (the agent must approve the escrow to pull
// its stake before fund()).
func Approve(from, token, spender core.Address, amount *big.Int) (chain.UnsignedTx, error) {
	parsed, err := abi.JSON(strings.NewReader(erc20ApproveABI))
	if err != nil {
		return chain.UnsignedTx{}, fmt.Errorf("txbuild: erc20 abi: %w", err)
	}
	data, err := parsed.Pack("approve", ethAddr(spender), amount)
	if err != nil {
		return chain.UnsignedTx{}, fmt.Errorf("txbuild: pack approve: %w", err)
	}
	return chain.BuildCall(from, token, nil, data), nil
}

// Fund builds the unsigned tx to fund the caller's side of a bet.
func Fund(from, escrow core.Address) (chain.UnsignedTx, error) {
	return packEscrow(from, escrow, "fund")
}

// Claim builds the unsigned tx to claim an outcome (opens the challenge window).
func Claim(from, escrow core.Address, outcome bets.Outcome, evidence core.Hash32) (chain.UnsignedTx, error) {
	return packEscrow(from, escrow, "claim", uint8(outcome), [32]byte(evidence))
}

// Challenge builds the unsigned tx to challenge a claim.
func Challenge(from, escrow core.Address) (chain.UnsignedTx, error) {
	return packEscrow(from, escrow, "challenge")
}

// Finalize builds the unsigned tx to finalize an unchallenged claim.
func Finalize(from, escrow core.Address) (chain.UnsignedTx, error) {
	return packEscrow(from, escrow, "finalize")
}

// VoidUnclaimed builds the unsigned tx to refund an unclaimed, expired bet.
func VoidUnclaimed(from, escrow core.Address) (chain.UnsignedTx, error) {
	return packEscrow(from, escrow, "voidUnclaimed")
}

func packEscrow(from, escrow core.Address, method string, args ...any) (chain.UnsignedTx, error) {
	parsed, err := betescrow.BetEscrowMetaData.GetAbi()
	if err != nil {
		return chain.UnsignedTx{}, fmt.Errorf("txbuild: escrow abi: %w", err)
	}
	data, err := parsed.Pack(method, args...)
	if err != nil {
		return chain.UnsignedTx{}, fmt.Errorf("txbuild: pack %s: %w", method, err)
	}
	return chain.BuildCall(from, escrow, nil, data), nil
}

// MintUSDC builds a MockUSDC mint tx (dev only — for funding agent wallets on a
// local anvil). Real USDC has no open mint.
func MintUSDC(from, token, to core.Address, amount *big.Int) (chain.UnsignedTx, error) {
	parsed, err := abi.JSON(strings.NewReader(mockMintABI))
	if err != nil {
		return chain.UnsignedTx{}, fmt.Errorf("txbuild: mint abi: %w", err)
	}
	data, err := parsed.Pack("mint", ethAddr(to), amount)
	if err != nil {
		return chain.UnsignedTx{}, fmt.Errorf("txbuild: pack mint: %w", err)
	}
	return chain.BuildCall(from, token, nil, data), nil
}

const erc20ApproveABI = `[{"name":"approve","type":"function","stateMutability":"nonpayable",` +
	`"inputs":[{"name":"spender","type":"address"},{"name":"amount","type":"uint256"}],` +
	`"outputs":[{"name":"","type":"bool"}]}]`

const mockMintABI = `[{"name":"mint","type":"function","stateMutability":"nonpayable",` +
	`"inputs":[{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],` +
	`"outputs":[]}]`
