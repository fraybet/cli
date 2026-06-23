package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/exec"
	"github.com/fraybet/cli/wallet"
)

// signerFlags are shared by every transaction command: which wallet to act as
// (default: the configured default wallet), which HD account, the RPC, and
// whether to print the unsigned tx instead of broadcasting it.
type signerFlags struct {
	wallet   *string
	account  *uint
	keystore *string
	rpc      *string
	unsigned *bool
}

func addSignerFlags(fs *flag.FlagSet) *signerFlags {
	return &signerFlags{
		wallet:   fs.String("wallet", "", "wallet to act as (default: the default wallet)"),
		account:  fs.Uint("account", 0, "HD account index (default 0)"),
		keystore: fs.String("keystore", "", "keystore dir (default ~/.agentbet/wallets)"),
		rpc:      fs.String("rpc", "", "RPC URL (default: Base mainnet)"),
		unsigned: fs.Bool("unsigned", false, "print the unsigned tx instead of signing + broadcasting"),
	}
}

func (sf *signerFlags) rpcURL() string {
	if r := strings.TrimSpace(*sf.rpc); r != "" {
		return r
	}
	return defaultRPCURL
}

// resolve loads the wallet to act as: --wallet if given, else the default wallet,
// at the chosen HD account index.
func (sf *signerFlags) resolve() (*wallet.Wallet, error) {
	st, err := openStore(*sf.keystore)
	if err != nil {
		return nil, err
	}
	name := strings.TrimSpace(*sf.wallet)
	if name == "" {
		name, err = st.Default()
		if err != nil {
			return nil, err
		}
		if name == "" {
			return nil, fmt.Errorf("no wallet selected: pass --wallet, or set a default with `fray wallet use <name>`")
		}
	}
	return st.LoadAccount(name, uint32(*sf.account))
}

// run resolves the signer, builds the tx from its address, then signs +
// broadcasts it — or prints the unsigned tx with --unsigned.
func (sf *signerFlags) run(out io.Writer, action, termsHash string, build func(from core.Address) (chain.UnsignedTx, error)) error {
	w, err := sf.resolve()
	if err != nil {
		return err
	}
	tx, err := build(w.Address)
	if err != nil {
		return err
	}
	if *sf.unsigned {
		return writeUnsignedTx(out, action, tx, termsHash)
	}
	return broadcastTx(out, action, termsHash, tx, w, sf.rpcURL())
}

// emit broadcasts a pre-built tx with the resolved wallet (or prints it unsigned
// with --unsigned). Use when the command needs the wallet/tx before transmit.
func (sf *signerFlags) emit(out io.Writer, action, termsHash string, tx chain.UnsignedTx, w *wallet.Wallet) error {
	if *sf.unsigned {
		return writeUnsignedTx(out, action, tx, termsHash)
	}
	return broadcastTx(out, action, termsHash, tx, w, sf.rpcURL())
}

// broadcastTx signs tx with the wallet's key and sends it, printing the result.
// Non-custodial: the key is the agent's own, held locally; nothing leaves the box.
func broadcastTx(out io.Writer, action, termsHash string, tx chain.UnsignedTx, w *wallet.Wallet, rpc string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	ex, err := exec.Dial(ctx, rpc)
	if err != nil {
		return err
	}
	defer ex.Close()
	fmt.Fprintf(os.Stderr, "broadcasting %q as %s (%s) …\n", action, w.Name, w.Address.Hex())
	res, err := ex.SendRaw(ctx, w.PrivHex(), tx.To, tx.Value, tx.Data)
	if err != nil {
		return err
	}
	result := map[string]string{
		"action": action,
		"from":   w.Address.Hex(),
		"to":     tx.To.Hex(),
		"txHash": res.TxHash.Hex(),
		"status": "mined",
	}
	if termsHash != "" {
		result["termsHash"] = termsHash
	}
	if res.Escrow != zeroAddr {
		result["escrow"] = res.Escrow.Hex()
	}
	return writeJSON(out, result)
}
