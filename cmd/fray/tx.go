package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
)

var zeroAddr core.Address

const txUsage = `fray tx <command> [flags]

commands:
  send   sign an unsigned tx (from a bet/agent --unsigned command) and broadcast

Most commands now sign + broadcast directly, so you rarely need this. Use it to
broadcast a tx that was built --unsigned (e.g. on another machine):
  fray bet fund --escrow $ESCROW --unsigned | fray tx send

Signs with your default wallet (override: --wallet/--account) on Base mainnet
(override: --rpc). Non-custodial: your key never leaves your machine.
`

func runTx(args []string, out io.Writer) error {
	if len(args) == 0 {
		fmt.Fprint(out, txUsage)
		return fmt.Errorf("tx: missing command")
	}
	cmd, rest := args[0], args[1:]
	switch cmd {
	case "send":
		return txSend(rest, out)
	case "help", "-h", "--help":
		fmt.Fprint(out, txUsage)
		return nil
	default:
		fmt.Fprint(out, txUsage)
		return fmt.Errorf("tx: unknown command %q", cmd)
	}
}

// unsignedEnvelope is the CLI's tx output shape (writeUnsignedTx / txEnvelope).
type unsignedEnvelope struct {
	Action     string           `json:"action"`
	TermsHash  string           `json:"termsHash"`
	UnsignedTx chain.UnsignedTx `json:"unsignedTx"`
}

func txSend(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("send", flag.ContinueOnError)
	fs.SetOutput(out)
	walletName := fs.String("wallet", "", "wallet to sign with (default: the default wallet)")
	account := fs.Uint("account", 0, "HD account index")
	keystore := fs.String("keystore", "", "keystore dir (default ~/.agentbet/wallets)")
	rpc := fs.String("rpc", "", "RPC URL (default: Base mainnet)")
	file := fs.String("file", "", "unsigned-tx JSON file (default: read stdin)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	raw, err := readInput(*file)
	if err != nil {
		return err
	}
	var env unsignedEnvelope
	if err := json.Unmarshal(raw, &env); err != nil {
		return fmt.Errorf("tx: parse unsigned tx: %w", err)
	}
	if env.UnsignedTx.To == zeroAddr {
		return fmt.Errorf("tx: input has no destination — pipe the output of a `bet`/`agent` command")
	}

	st, err := openStore(*keystore)
	if err != nil {
		return err
	}
	name := strings.TrimSpace(*walletName)
	if name == "" {
		if name, err = st.Default(); err != nil {
			return err
		}
		if name == "" {
			return fmt.Errorf("no wallet: pass --wallet or set a default (fray wallet use <name>)")
		}
	}
	w, err := st.LoadAccount(name, uint32(*account))
	if err != nil {
		return err
	}
	// Guard against signing a tx built for a different sender.
	if env.UnsignedTx.From != zeroAddr && env.UnsignedTx.From != w.Address {
		return fmt.Errorf("tx: unsigned tx 'from' %s != wallet %s (%s)",
			env.UnsignedTx.From.Hex(), name, w.Address.Hex())
	}
	rpcURL := strings.TrimSpace(*rpc)
	if rpcURL == "" {
		rpcURL = defaultRPCURL
	}
	return broadcastTx(out, env.Action, env.TermsHash, env.UnsignedTx, w, rpcURL)
}

func readInput(file string) ([]byte, error) {
	if strings.TrimSpace(file) == "" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(file)
}
