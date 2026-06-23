package main

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/txbuild"
)

const agentUsage = `fray agent <command> [flags]

commands:
  register     register an agent (required to post public/arbitered bets).
               Approve the fee + bond to the registry FIRST:
                 fray bet approve --token <USDC> --spender <registry> --amount <fee+bond> | fray tx send …
  deactivate   deactivate an agent and refund its bond (the bond claw-back)

All commands print an UNSIGNED transaction — pipe to `+"`fray tx send`"+` to broadcast.
`

func runAgent(args []string, out io.Writer) error {
	if len(args) == 0 {
		fmt.Fprint(out, agentUsage)
		return fmt.Errorf("agent: missing command")
	}
	cmd, rest := args[0], args[1:]
	switch cmd {
	case "register":
		return agentRegister(rest, out)
	case "deactivate":
		return agentDeactivate(rest, out)
	case "help", "-h", "--help":
		fmt.Fprint(out, agentUsage)
		return nil
	default:
		fmt.Fprint(out, agentUsage)
		return fmt.Errorf("agent: unknown command %q", cmd)
	}
}

func agentRegister(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("register", flag.ContinueOnError)
	fs.SetOutput(out)
	from := fs.String("from", "", "tx sender / principal (pays fee + bond)")
	wallet := fs.String("wallet", "", "agent wallet to register (default: --from)")
	signer := fs.String("signer", "", "signing key for message auth (default: --wallet)")
	registry := fs.String("registry", defaultRegistry, "AgentRegistry address")
	policy := fs.String("policy-hash", "", "policy hash (bytes32; default zero)")
	metadata := fs.String("metadata-hash", "", "metadata hash (bytes32; default zero)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	fromA, err := parseAddr(*from, "from")
	if err != nil {
		return err
	}
	walletStr := *wallet
	if strings.TrimSpace(walletStr) == "" {
		walletStr = *from
	}
	walletA, err := parseAddr(walletStr, "wallet")
	if err != nil {
		return err
	}
	signerStr := *signer
	if strings.TrimSpace(signerStr) == "" {
		signerStr = walletStr
	}
	signerA, err := parseAddr(signerStr, "signer")
	if err != nil {
		return err
	}
	registryA, err := parseAddr(*registry, "registry")
	if err != nil {
		return err
	}
	policyH, err := optHash32(*policy)
	if err != nil {
		return err
	}
	metaH, err := optHash32(*metadata)
	if err != nil {
		return err
	}
	tx, err := txbuild.AgentRegister(fromA, registryA, walletA, signerA, policyH, metaH)
	if err != nil {
		return err
	}
	return writeUnsignedTx(out, "register agent", tx, "")
}

func agentDeactivate(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("deactivate", flag.ContinueOnError)
	fs.SetOutput(out)
	from := fs.String("from", "", "tx sender (the agent owner)")
	wallet := fs.String("wallet", "", "agent wallet to deactivate (default: --from)")
	registry := fs.String("registry", defaultRegistry, "AgentRegistry address")
	if err := fs.Parse(args); err != nil {
		return err
	}
	fromA, err := parseAddr(*from, "from")
	if err != nil {
		return err
	}
	walletStr := *wallet
	if strings.TrimSpace(walletStr) == "" {
		walletStr = *from
	}
	walletA, err := parseAddr(walletStr, "wallet")
	if err != nil {
		return err
	}
	registryA, err := parseAddr(*registry, "registry")
	if err != nil {
		return err
	}
	tx, err := txbuild.AgentDeactivate(fromA, registryA, walletA)
	if err != nil {
		return err
	}
	return writeUnsignedTx(out, "deactivate agent (refund bond)", tx, "")
}

// optHash32 parses a bytes32 hex, defaulting to the zero hash when empty.
func optHash32(s string) (core.Hash32, error) {
	if strings.TrimSpace(s) == "" {
		return core.Hash32{}, nil
	}
	return core.HexToHash32(s)
}
