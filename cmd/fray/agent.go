package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

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
  set-name     claim a unique @handle for your wallet (signs an ownership
               challenge; shown across the site in place of your address)

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
	case "set-name":
		return agentSetName(rest, out)
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

// agentSetName claims a unique @handle for the wallet. Unlike the tx commands it
// touches no chain: it fetches a sign-in challenge from the site, signs it
// (personal_sign, proving wallet ownership), and POSTs the handle. Uniqueness and
// the registration gate are enforced server-side.
func agentSetName(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("set-name", flag.ContinueOnError)
	fs.SetOutput(out)
	name := fs.String("name", "", "the @handle to claim (1–15 chars: letters, digits, underscore)")
	baseURL := fs.String("base-url", "", "site base URL (env FRAY_BASE_URL; default https://fray.bet)")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*name) == "" {
		return fmt.Errorf("set-name: --name is required")
	}
	w, err := sf.resolve()
	if err != nil {
		return err
	}
	base := resolveBaseURL(*baseURL)

	// The signed message binds the handle to this address, so the signature IS the
	// handle→address claim — the server cannot bind a different handle than the one
	// we signed. No on-chain tx; replay is bounded by `expires`.
	handle := core.NormalizeHandle(*name)
	expires := time.Now().Add(5 * time.Minute).Unix()
	msg := core.HandleClaimMessage(handle, w.Address, expires)
	sig, err := w.Sign(core.PersonalDigest(msg))
	if err != nil {
		return fmt.Errorf("set-name: sign: %w", err)
	}
	claimed, err := postSetName(base, w.Address, handle, expires, sig)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "claimed @%s for %s\n", claimed, w.Address.Hex())
	return nil
}

// postSetName submits the signed handle claim and returns the stored handle.
func postSetName(base string, addr core.Address, name string, expires int64, sig []byte) (string, error) {
	body, _ := json.Marshal(map[string]any{
		"address":   addr.Hex(),
		"name":      name,
		"expires":   expires,
		"signature": "0x" + hex.EncodeToString(sig),
	})
	resp, err := http.Post(strings.TrimRight(base, "/")+"/agent/name", "application/json", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("set-name: request: %w", err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("set-name: server rejected (%d): %s", resp.StatusCode, strings.TrimSpace(string(raw)))
	}
	var res struct {
		Name string `json:"name"`
	}
	_ = json.Unmarshal(raw, &res)
	if res.Name == "" {
		res.Name = strings.TrimPrefix(strings.TrimSpace(name), "@")
	}
	return res.Name, nil
}
