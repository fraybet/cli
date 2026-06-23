package main

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/txbuild"
)

const agentUsage = `fray agent <command> [flags]

commands:
  register     register an agent (required to post public/arbitered bets).
               Approve the fee + bond to the STORAGE vault (the funds custodian)
               FIRST: fray bet approve --spender <storage> --amount <fee+bond>
  deactivate   deactivate an agent and refund its bond (the bond claw-back).
               Blocked while the agent has bond reserved against open arbitered bets.
  sweep        sweep accrued protocol fees to the revenue wallet (anyone may call)

tx commands sign with your default wallet and broadcast to Base mainnet
(override: --wallet / --account / --rpc; print unsigned with --unsigned).
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
	case "sweep":
		return agentSweep(rest, out)
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
	agent := fs.String("agent", "", "agent wallet to register (default: the signing wallet)")
	signer := fs.String("signer", "", "signing key for message auth (default: --agent)")
	registry := fs.String("registry", defaultRegistry, "AgentRegistry address")
	policy := fs.String("policy-hash", "", "policy hash (bytes32; default zero)")
	metadata := fs.String("metadata-hash", "", "metadata hash (bytes32; default zero)")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
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
	return sf.run(out, "register agent", "", func(from core.Address) (chain.UnsignedTx, error) {
		agentA := from
		if strings.TrimSpace(*agent) != "" {
			if agentA, err = parseAddr(*agent, "agent"); err != nil {
				return chain.UnsignedTx{}, err
			}
		}
		signerA := agentA
		if strings.TrimSpace(*signer) != "" {
			if signerA, err = parseAddr(*signer, "signer"); err != nil {
				return chain.UnsignedTx{}, err
			}
		}
		return txbuild.AgentRegister(from, registryA, agentA, signerA, policyH, metaH)
	})
}

func agentDeactivate(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("deactivate", flag.ContinueOnError)
	fs.SetOutput(out)
	agent := fs.String("agent", "", "agent wallet to deactivate (default: the signing wallet)")
	registry := fs.String("registry", defaultRegistry, "AgentRegistry address")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	registryA, err := parseAddr(*registry, "registry")
	if err != nil {
		return err
	}
	return sf.run(out, "deactivate agent (refund bond)", "", func(from core.Address) (chain.UnsignedTx, error) {
		agentA := from
		if strings.TrimSpace(*agent) != "" {
			if agentA, err = parseAddr(*agent, "agent"); err != nil {
				return chain.UnsignedTx{}, err
			}
		}
		return txbuild.AgentDeactivate(from, registryA, agentA)
	})
}

func agentSweep(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("sweep", flag.ContinueOnError)
	fs.SetOutput(out)
	registry := fs.String("registry", defaultRegistry, "AgentRegistry address")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	registryA, err := parseAddr(*registry, "registry")
	if err != nil {
		return err
	}
	return sf.run(out, "sweep fees", "", func(from core.Address) (chain.UnsignedTx, error) {
		return txbuild.AgentSweep(from, registryA)
	})
}

// optHash32 parses a bytes32 hex, defaulting to the zero hash when empty.
func optHash32(s string) (core.Hash32, error) {
	if strings.TrimSpace(s) == "" {
		return core.Hash32{}, nil
	}
	return core.HexToHash32(s)
}
