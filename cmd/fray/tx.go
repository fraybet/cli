package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/exec"
)

var zeroAddr core.Address

const txUsage = `fray tx <command> [flags]

commands:
  send   sign an unsigned tx (piped from a bet/wallet command) with a keystore
         wallet and broadcast it to a chain

Example — fund your side of a bet end to end:
  fray bet fund --from $ADDR --escrow $ESCROW \
    | fray tx send --wallet my-agent --rpc https://mainnet.base.org

The CLI signs locally with your keystore key and broadcasts. Non-custodial:
your key never leaves your machine.
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
	walletName := fs.String("wallet", "", "keystore wallet name to sign with")
	keystore := fs.String("keystore", "", "keystore dir (default ~/.agentbet/wallets)")
	rpc := fs.String("rpc", "", "RPC URL (required)")
	file := fs.String("file", "", "unsigned-tx JSON file (default: read stdin)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*walletName) == "" {
		return fmt.Errorf("--wallet is required")
	}
	if strings.TrimSpace(*rpc) == "" {
		return fmt.Errorf("--rpc is required")
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
		return fmt.Errorf("tx: input has no destination — pipe the output of a `bet`/`wallet` command")
	}

	st, err := openStore(*keystore)
	if err != nil {
		return err
	}
	w, err := st.Load(*walletName)
	if err != nil {
		return err
	}
	// Guard against signing a tx built for a different sender.
	if env.UnsignedTx.From != zeroAddr && env.UnsignedTx.From != w.Address {
		return fmt.Errorf("tx: unsigned tx 'from' %s != wallet %q (%s); rebuild with --from %s",
			env.UnsignedTx.From.Hex(), *walletName, w.Address.Hex(), w.Address.Hex())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	ex, err := exec.Dial(ctx, *rpc)
	if err != nil {
		return err
	}
	defer ex.Close()

	fmt.Fprintf(os.Stderr, "broadcasting %q from %s …\n", env.Action, w.Address.Hex())
	res, err := ex.SendRaw(ctx, w.PrivHex(), env.UnsignedTx.To, env.UnsignedTx.Value, env.UnsignedTx.Data)
	if err != nil {
		return err
	}
	result := map[string]string{
		"action": env.Action,
		"from":   w.Address.Hex(),
		"to":     env.UnsignedTx.To.Hex(),
		"txHash": res.TxHash.Hex(),
		"status": "mined",
	}
	// On a create, hand back the new escrow so the next step (approve/fund) has it.
	if res.Escrow != zeroAddr {
		result["escrow"] = res.Escrow.Hex()
	}
	return writeJSON(out, result)
}

func readInput(file string) ([]byte, error) {
	if strings.TrimSpace(file) == "" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(file)
}
