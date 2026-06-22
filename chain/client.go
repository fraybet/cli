package chain

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/fraybet/cli/core"
)

// Client wraps a go-ethereum RPC client for the read + simulate operations the
// platform needs. It never signs or broadcasts on a user's behalf.
type Client struct {
	profile Profile
	eth     *ethclient.Client
}

// Dial connects to the profile's RPC endpoint.
func Dial(ctx context.Context, p Profile) (*Client, error) {
	if p.RPCURL == "" {
		return nil, fmt.Errorf("chain: profile %q has no RPC URL", p.Name)
	}
	ec, err := ethclient.DialContext(ctx, p.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("chain: dial %s: %w", p.RPCURL, err)
	}
	return &Client{profile: p, eth: ec}, nil
}

// Close releases the RPC connection.
func (c *Client) Close() { c.eth.Close() }

// Profile returns the client's profile.
func (c *Client) Profile() Profile { return c.profile }

// BalanceAt returns the native (ETH) balance of an address at the latest block.
func (c *Client) BalanceAt(ctx context.Context, a core.Address) (*big.Int, error) {
	return c.eth.BalanceAt(ctx, toETH(a), nil)
}

// PendingNonce returns the next nonce for an address.
func (c *Client) PendingNonce(ctx context.Context, a core.Address) (uint64, error) {
	return c.eth.PendingNonceAt(ctx, toETH(a))
}

// EstimateGas estimates gas for an unsigned tx.
func (c *Client) EstimateGas(ctx context.Context, tx UnsignedTx) (uint64, error) {
	return c.eth.EstimateGas(ctx, c.callMsg(tx))
}

// Simulate runs eth_call against the latest block and returns the raw return
// data — a dry-run of a state-changing call before the user signs it.
func (c *Client) Simulate(ctx context.Context, tx UnsignedTx) ([]byte, error) {
	return c.eth.CallContract(ctx, c.callMsg(tx), nil)
}

func (c *Client) callMsg(tx UnsignedTx) ethereum.CallMsg {
	to := toETH(tx.To)
	return ethereum.CallMsg{From: toETH(tx.From), To: &to, Value: tx.Value, Data: tx.Data}
}

func toETH(a core.Address) ethcommon.Address { return ethcommon.BytesToAddress(a[:]) }

func fromETH(a ethcommon.Address) core.Address {
	var out core.Address
	copy(out[:], a.Bytes())
	return out
}
