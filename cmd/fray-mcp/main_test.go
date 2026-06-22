package main

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// connect wires a client to newServer() over an in-memory transport.
func connect(t *testing.T) (*mcp.ClientSession, context.Context) {
	t.Helper()
	ctx := context.Background()
	st, ct := mcp.NewInMemoryTransports()
	if _, err := newServer().Connect(ctx, st, nil); err != nil {
		t.Fatalf("server connect: %v", err)
	}
	client := mcp.NewClient(&mcp.Implementation{Name: "test", Version: "0"}, nil)
	cs, err := client.Connect(ctx, ct, nil)
	if err != nil {
		t.Fatalf("client connect: %v", err)
	}
	t.Cleanup(func() { cs.Close() })
	return cs, ctx
}

// structured re-decodes a tool result's StructuredContent into v.
func structured(t *testing.T, res *mcp.CallToolResult, v any) {
	t.Helper()
	if res.IsError {
		var msg strings.Builder
		for _, c := range res.Content {
			if tc, ok := c.(*mcp.TextContent); ok {
				msg.WriteString(tc.Text)
			}
		}
		t.Fatalf("tool returned error: %s", msg.String())
	}
	b, err := json.Marshal(res.StructuredContent)
	if err != nil {
		t.Fatalf("marshal structured: %v", err)
	}
	if err := json.Unmarshal(b, v); err != nil {
		t.Fatalf("decode structured: %v", err)
	}
}

func TestListToolsExposesSuite(t *testing.T) {
	cs, ctx := connect(t)
	lt, err := cs.ListTools(ctx, nil)
	if err != nil {
		t.Fatalf("list tools: %v", err)
	}
	got := map[string]bool{}
	for _, tool := range lt.Tools {
		got[tool.Name] = true
	}
	for _, want := range []string{
		"bet_draft", "bet_propose_link", "bet_accept_link", "bet_create",
		"bet_approve", "bet_fund", "bet_claim", "bet_challenge", "bet_finalize",
		"bet_void", "usdc_mint",
	} {
		if !got[want] {
			t.Errorf("missing tool %q", want)
		}
	}
}

func TestProposeThenAcceptRoundTrip(t *testing.T) {
	cs, ctx := connect(t)

	// 1) propose a one-sided challenge (proposer takes YES).
	res, err := cs.CallTool(ctx, &mcp.CallToolParams{
		Name: "bet_propose_link",
		Arguments: map[string]any{
			"proposer":  "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			"side":      "YES",
			"token":     "0x5FbDB2315678afecb367f032d93F642f64180aa3",
			"stake":     "100000000",
			"statement": "BTC > 100k EOY",
			"source":    "coingecko",
			"eventTime": 1800000000,
			"public":    true,
			"baseUrl":   "https://fray.example",
		},
	})
	if err != nil {
		t.Fatalf("propose call: %v", err)
	}
	var prop proposeView
	structured(t, res, &prop)
	if prop.AccepterTakes != "NO" || prop.Token == "" {
		t.Fatalf("unexpected propose view: %+v", prop)
	}
	if !strings.HasPrefix(prop.Link, "https://fray.example/challenge?p=") {
		t.Errorf("unexpected link: %s", prop.Link)
	}

	// 2) accept it as the counterparty, also emitting the create tx.
	res, err = cs.CallTool(ctx, &mcp.CallToolParams{
		Name: "bet_accept_link",
		Arguments: map[string]any{
			"token":   prop.Token,
			"as":      "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
			"factory": "0x5FbDB2315678afecb367f032d93F642f64180aa3",
		},
	})
	if err != nil {
		t.Fatalf("accept call: %v", err)
	}
	var acc acceptView
	structured(t, res, &acc)

	// Proposer (YES) maps to yesAgent; accepter to noAgent; stakes match.
	if !strings.EqualFold(acc.Draft.Terms.YesAgent, "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266") {
		t.Errorf("yesAgent = %s", acc.Draft.Terms.YesAgent)
	}
	if !strings.EqualFold(acc.Draft.Terms.NoAgent, "0x70997970c51812dc3a010c7d01b50e0d17dc79c8") {
		t.Errorf("noAgent = %s", acc.Draft.Terms.NoAgent)
	}
	if acc.Draft.Terms.YesStake != "100000000" || acc.Draft.Terms.NoStake != "100000000" {
		t.Errorf("stakes = %s / %s", acc.Draft.Terms.YesStake, acc.Draft.Terms.NoStake)
	}
	if acc.CreateTx == nil {
		t.Fatal("expected a create tx when factory is provided")
	}
	if acc.CreateTx.TermsHash != acc.Draft.TermsHash {
		t.Errorf("termsHash mismatch: %s vs %s", acc.CreateTx.TermsHash, acc.Draft.TermsHash)
	}
	if !strings.HasPrefix(acc.CreateTx.UnsignedTx.Data, "0x") || acc.CreateTx.UnsignedTx.Data == "0x" {
		t.Errorf("expected create calldata, got %q", acc.CreateTx.UnsignedTx.Data)
	}
}

func TestFundReturnsUnsignedTx(t *testing.T) {
	cs, ctx := connect(t)
	res, err := cs.CallTool(ctx, &mcp.CallToolParams{
		Name: "bet_fund",
		Arguments: map[string]any{
			"from":   "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
			"escrow": "0x5FbDB2315678afecb367f032d93F642f64180aa3",
		},
	})
	if err != nil {
		t.Fatalf("fund call: %v", err)
	}
	var env txEnvelope
	structured(t, res, &env)
	if env.Action != "fund" || env.UnsignedTx.To == "" || !strings.Contains(env.Note, "UNSIGNED") {
		t.Fatalf("unexpected fund envelope: %+v", env)
	}
}

func TestBadArgsReturnsToolError(t *testing.T) {
	cs, ctx := connect(t)
	res, err := cs.CallTool(ctx, &mcp.CallToolParams{
		Name:      "bet_fund",
		Arguments: map[string]any{"from": "not-an-address", "escrow": ""},
	})
	if err != nil {
		t.Fatalf("call transport error: %v", err)
	}
	if !res.IsError {
		t.Error("expected IsError for bad address args")
	}
}

func TestGuideResource(t *testing.T) {
	cs, ctx := connect(t)
	rr, err := cs.ReadResource(ctx, &mcp.ReadResourceParams{URI: "fray://guide"})
	if err != nil {
		t.Fatalf("read resource: %v", err)
	}
	if len(rr.Contents) == 0 || !strings.Contains(rr.Contents[0].Text, "non-custodial") {
		t.Errorf("guide content unexpected: %+v", rr.Contents)
	}
}
