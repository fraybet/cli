// Package exec signs and broadcasts an agent's economic actions to a chain. It
// is the bridge from prepared calldata (internal/txbuild) to a live node: the
// agent self-custodies its key (internal/wallet), so it signs its own txs here —
// the platform never signs. Used by the demo runtime against a local anvil.
package exec

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/fraybet/cli/bets"
	bef "github.com/fraybet/cli/bindings/betescrowfactory"
	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/txbuild"
)

// Executor broadcasts signed transactions to a chain.
type Executor struct {
	eth     *ethclient.Client
	chainID *big.Int
}

// Dial connects to an RPC endpoint (e.g. http://localhost:8545 for anvil).
func Dial(ctx context.Context, rpcURL string) (*Executor, error) {
	eth, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("exec: dial %s: %w", rpcURL, err)
	}
	id, err := eth.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec: chain id: %w", err)
	}
	return &Executor{eth: eth, chainID: id}, nil
}

// Close releases the connection.
func (e *Executor) Close() { e.eth.Close() }

// SendETH transfers native ETH (used to fund fresh agent wallets with gas money
// from a pre-funded anvil account).
func (e *Executor) SendETH(ctx context.Context, signerHex string, to core.Address, amount *big.Int) error {
	key, _, err := keyAndAddr(signerHex)
	if err != nil {
		return err
	}
	_, err = e.sendTx(ctx, key, to, amount, nil)
	return err
}

// MintUSDC mints dev USDC to an address (MockUSDC only). signerHex pays gas.
func (e *Executor) MintUSDC(ctx context.Context, signerHex string, token, to core.Address, amount *big.Int) error {
	key, from, err := keyAndAddr(signerHex)
	if err != nil {
		return err
	}
	tx, err := txbuild.MintUSDC(from, token, to, amount)
	if err != nil {
		return err
	}
	_, err = e.sendTx(ctx, key, token, big.NewInt(0), tx.Data)
	return err
}

// Approve lets `spender` pull `amount` of `token` from the signer.
func (e *Executor) Approve(ctx context.Context, signerHex string, token, spender core.Address, amount *big.Int) error {
	key, from, err := keyAndAddr(signerHex)
	if err != nil {
		return err
	}
	tx, err := txbuild.Approve(from, token, spender, amount)
	if err != nil {
		return err
	}
	_, err = e.sendTx(ctx, key, token, big.NewInt(0), tx.Data)
	return err
}

// CreateBet sends factory.create(terms) and returns the new escrow address
// (parsed from the BetCreated event).
func (e *Executor) CreateBet(ctx context.Context, signerHex string, factory core.Address, d bets.Draft) (core.Address, error) {
	key, from, err := keyAndAddr(signerHex)
	if err != nil {
		return core.Address{}, err
	}
	tx, err := txbuild.CreateBet(from, factory, d)
	if err != nil {
		return core.Address{}, err
	}
	rcpt, err := e.sendTx(ctx, key, factory, big.NewInt(0), tx.Data)
	if err != nil {
		return core.Address{}, err
	}
	f, err := bef.NewBetEscrowFactory(toEth(factory), e.eth)
	if err != nil {
		return core.Address{}, err
	}
	for _, lg := range rcpt.Logs {
		if ev, perr := f.ParseBetCreated(*lg); perr == nil {
			return fromEth(ev.Escrow), nil
		}
	}
	return core.Address{}, fmt.Errorf("exec: no BetCreated event in receipt")
}

// Fund funds the signer's side of a bet escrow.
func (e *Executor) Fund(ctx context.Context, signerHex string, escrow core.Address) error {
	key, from, err := keyAndAddr(signerHex)
	if err != nil {
		return err
	}
	tx, err := txbuild.Fund(from, escrow)
	if err != nil {
		return err
	}
	_, err = e.sendTx(ctx, key, escrow, big.NewInt(0), tx.Data)
	return err
}

// BalanceOfUSDC reads an ERC-20 balance.
func (e *Executor) BalanceOfUSDC(ctx context.Context, token, who core.Address) (*big.Int, error) {
	parsed, err := abi.JSON(strings.NewReader(erc20BalanceABI))
	if err != nil {
		return nil, err
	}
	data, err := parsed.Pack("balanceOf", toEth(who))
	if err != nil {
		return nil, err
	}
	to := toEth(token)
	out, err := e.eth.CallContract(ctx, ethereum.CallMsg{To: &to, Data: data}, nil)
	if err != nil {
		return nil, fmt.Errorf("exec: balanceOf: %w", err)
	}
	vals, err := parsed.Unpack("balanceOf", out)
	if err != nil {
		return nil, err
	}
	return vals[0].(*big.Int), nil
}

// SendResult is the outcome of a broadcast. Escrow is set (non-zero) only when
// the receipt carried a BetCreated event — i.e. the tx created a new escrow.
type SendResult struct {
	TxHash core.Hash32
	Escrow core.Address
}

// SendRaw signs and broadcasts an arbitrary call (to/value/data, e.g. from
// internal/txbuild) with the agent's own key, returning the mined tx hash (and
// the new escrow address if the tx created one). This is the generic "execute
// this unsigned transaction" path the CLI uses — the agent self-custodies, so
// it signs locally; the platform never sees the key.
func (e *Executor) SendRaw(ctx context.Context, signerHex string, to core.Address, value *big.Int, data []byte) (SendResult, error) {
	key, _, err := keyAndAddr(signerHex)
	if err != nil {
		return SendResult{}, err
	}
	if value == nil {
		value = big.NewInt(0)
	}
	rcpt, err := e.sendTx(ctx, key, to, value, data)
	if err != nil {
		return SendResult{}, err
	}
	res := SendResult{}
	copy(res.TxHash[:], rcpt.TxHash.Bytes())
	// If this tx created a bet, surface the escrow address so the caller can
	// fund/claim it next. Parsing is best-effort; non-create txs simply skip.
	if f, ferr := bef.NewBetEscrowFactory(toEth(to), e.eth); ferr == nil {
		for _, lg := range rcpt.Logs {
			if ev, perr := f.ParseBetCreated(*lg); perr == nil {
				res.Escrow = fromEth(ev.Escrow)
				break
			}
		}
	}
	return res, nil
}

// sendTx signs and broadcasts a tx, waiting for a successful receipt.
func (e *Executor) sendTx(ctx context.Context, key *ecdsa.PrivateKey, to core.Address, value *big.Int, data []byte) (*types.Receipt, error) {
	from := ethcrypto.PubkeyToAddress(key.PublicKey)
	toAddr := toEth(to)
	nonce, err := e.eth.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, fmt.Errorf("exec: nonce: %w", err)
	}
	gasPrice, err := e.eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec: gas price: %w", err)
	}
	gas, err := e.eth.EstimateGas(ctx, ethereum.CallMsg{From: from, To: &toAddr, Value: value, Data: data})
	if err != nil {
		return nil, fmt.Errorf("exec: estimate gas: %w", err)
	}
	tx := types.NewTx(&types.LegacyTx{Nonce: nonce, GasPrice: gasPrice, Gas: gas, To: &toAddr, Value: value, Data: data})
	signed, err := types.SignTx(tx, types.LatestSignerForChainID(e.chainID), key)
	if err != nil {
		return nil, fmt.Errorf("exec: sign: %w", err)
	}
	if err := e.eth.SendTransaction(ctx, signed); err != nil {
		return nil, fmt.Errorf("exec: send: %w", err)
	}
	rcpt, err := bind.WaitMined(ctx, e.eth, signed)
	if err != nil {
		return nil, fmt.Errorf("exec: wait mined: %w", err)
	}
	if rcpt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("exec: tx %s reverted", signed.Hash().Hex())
	}
	return rcpt, nil
}

func keyAndAddr(hexKey string) (*ecdsa.PrivateKey, core.Address, error) {
	key, err := ethcrypto.HexToECDSA(strings.TrimPrefix(strings.TrimSpace(hexKey), "0x"))
	if err != nil {
		return nil, core.Address{}, fmt.Errorf("exec: bad key: %w", err)
	}
	return key, fromEth(ethcrypto.PubkeyToAddress(key.PublicKey)), nil
}

func toEth(a core.Address) ethcommon.Address { return ethcommon.BytesToAddress(a[:]) }

func fromEth(a ethcommon.Address) core.Address {
	var o core.Address
	copy(o[:], a.Bytes())
	return o
}

const erc20BalanceABI = `[{"name":"balanceOf","type":"function","stateMutability":"view",` +
	`"inputs":[{"name":"a","type":"address"}],"outputs":[{"name":"","type":"uint256"}]}]`
