// Command fray-mcp is a Model Context Protocol server that exposes the
// Fray toolkit over stdio, so an MCP-capable agent (e.g. Claude) can draft
// bets, mint shareable challenge links, accept them, and build the on-chain
// transactions for the full bilateral-bet lifecycle.
//
// It is non-custodial by construction: every transaction tool returns an
// UNSIGNED transaction the agent signs and broadcasts with its own wallet. The
// server holds no keys and never signs. The pure tools (draft / propose_link /
// accept_link) hold no state, so a link minted here is self-contained.
//
// Run it as an MCP stdio server:
//
//	fray-mcp
//
// or point an MCP client config at the built binary.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/txbuild"
)

const (
	serverName     = "fray"
	serverVersion  = "0.1.0"
	defaultBaseURL = "https://fray.bet"
)

var errNoToken = errors.New("provide token or link")

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("fray-mcp: %v", err)
	}
}

func run(ctx context.Context) error {
	return newServer().Run(ctx, &mcp.StdioTransport{})
}

// newServer builds the Fray MCP server with all tools and resources
// registered. Split out from run so tests can drive it over an in-memory
// transport.
func newServer() *mcp.Server {
	s := mcp.NewServer(&mcp.Implementation{Name: serverName, Version: serverVersion}, nil)

	// Compose tools (pure, stateless).
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_draft",
		Description: "Build and validate bilateral bet terms; returns the canonical termsHash. No transaction.",
	}, draftTool)
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_propose_link",
		Description: "Generate a shareable one-sided challenge link. Post it (e.g. on X) so a bot or person can take the other side. No funds move until both fund the escrow.",
	}, proposeTool)
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_accept_link",
		Description: "Accept a challenge link as the counterparty. Returns the resolved draft, plus the unsigned create transaction if a factory is given.",
	}, acceptTool)

	// Lifecycle tools (each returns an UNSIGNED transaction).
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_create",
		Description: "Build the UNSIGNED transaction to create a bet via the BetEscrowFactory.",
	}, createTool)
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_approve",
		Description: "Build the UNSIGNED ERC-20 approve transaction (run before funding).",
	}, approveTool)
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_fund",
		Description: "Build the UNSIGNED transaction to fund your side of a bet.",
	}, escrowAction("fund", func(from, escrow core.Address) (txEnvelope, error) {
		tx, err := txbuild.Fund(from, escrow)
		if err != nil {
			return txEnvelope{}, err
		}
		return envelope("fund", "", tx)
	}))
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_claim",
		Description: "Build the UNSIGNED transaction to claim an outcome (YES or NO), opening the challenge window.",
	}, claimTool)
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_challenge",
		Description: "Build the UNSIGNED transaction to challenge a claim within its window (requires an arbiter).",
	}, escrowAction("challenge", func(from, escrow core.Address) (txEnvelope, error) {
		tx, err := txbuild.Challenge(from, escrow)
		if err != nil {
			return txEnvelope{}, err
		}
		return envelope("challenge", "", tx)
	}))
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_finalize",
		Description: "Build the UNSIGNED transaction to finalize an unchallenged claim after its window closes.",
	}, escrowAction("finalize", func(from, escrow core.Address) (txEnvelope, error) {
		tx, err := txbuild.Finalize(from, escrow)
		if err != nil {
			return txEnvelope{}, err
		}
		return envelope("finalize", "", tx)
	}))
	mcp.AddTool(s, &mcp.Tool{
		Name:        "bet_void",
		Description: "Build the UNSIGNED transaction to refund an unclaimed, expired bet.",
	}, escrowAction("void", func(from, escrow core.Address) (txEnvelope, error) {
		tx, err := txbuild.VoidUnclaimed(from, escrow)
		if err != nil {
			return txEnvelope{}, err
		}
		return envelope("void", "", tx)
	}))
	mcp.AddTool(s, &mcp.Tool{
		Name:        "usdc_mint",
		Description: "Dev/anvil only: build the UNSIGNED transaction to mint mock USDC to an address.",
	}, mintTool)

	// A read-only onboarding guide.
	s.AddResource(&mcp.Resource{
		URI:         "fray://guide",
		Name:        "fray-guide",
		Title:       "Fray agent guide",
		Description: "How an agent uses Fray: the bet lifecycle and the non-custodial signing model.",
		MIMEType:    "text/markdown",
	}, func(_ context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return &mcp.ReadResourceResult{Contents: []*mcp.ResourceContents{{
			URI:      "fray://guide",
			MIMEType: "text/markdown",
			Text:     guideText,
		}}}, nil
	})

	return s
}

// mint is registered inline so it can reach txbuild without an escrow shape.
type mintArgs struct {
	From   string `json:"from" jsonschema:"tx sender (the minter)"`
	Token  string `json:"token" jsonschema:"MockUSDC token address"`
	To     string `json:"to" jsonschema:"recipient address"`
	Amount string `json:"amount" jsonschema:"amount in base units (USDC 6dp)"`
}

func mintTool(_ context.Context, _ *mcp.CallToolRequest, in mintArgs) (*mcp.CallToolResult, txEnvelope, error) {
	from, err := mustAddr(in.From, "from")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	token, err := mustAddr(in.Token, "token")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	to, err := mustAddr(in.To, "to")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	amt, err := mustBig(in.Amount, "amount")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	tx, err := txbuild.MintUSDC(from, token, to, amt)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	env, err := envelope("mint usdc", "", tx)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	return nil, env, nil
}

func resolveBaseURL(flagVal string) string {
	if strings.TrimSpace(flagVal) != "" {
		return flagVal
	}
	if env := strings.TrimSpace(os.Getenv("FRAY_BASE_URL")); env != "" {
		return env
	}
	return defaultBaseURL
}

var guideText = fmt.Sprintf(`# Fray — agent guide

Fray is a non-custodial, agent-to-agent binary betting protocol. This MCP
server never holds keys or funds: every transaction tool returns an **unsigned**
transaction that you sign and broadcast with your own wallet.

## Challenge a bot or person
1. `+"`bet_propose_link`"+` — commit to terms and the side you take; you get a
   shareable link + token. Post it on X. The counterparty is still open.
2. The taker opens the link and calls `+"`bet_accept_link`"+` with their address
   (optionally a factory) to get the resolved draft / unsigned create tx.

## Strike a bet directly
1. `+"`bet_draft`"+` → validate terms, get the termsHash.
2. `+"`bet_create`"+` → unsigned tx to deploy the escrow via the factory.
3. `+"`bet_approve`"+` then `+"`bet_fund`"+` → both sides escrow their stake; the bet goes Live.

## Resolve
- Fast path: both sides agree off-tool and co-sign the same outcome on-chain.
- Claim path: `+"`bet_claim`"+` (YES/NO) opens a challenge window; `+"`bet_finalize`"+`
  settles it if unchallenged, or `+"`bet_challenge`"+` routes to the arbiter.
- `+"`bet_void`"+` refunds an unclaimed, expired bet.

Arbitered bets carry a small protocol fee (base %% of the pot + a fixed
execution fee the loser pays on a real dispute). Non-arbitered bets are free.

Dev: on a local anvil chain, `+"`usdc_mint`"+` funds test wallets with mock USDC.`)
