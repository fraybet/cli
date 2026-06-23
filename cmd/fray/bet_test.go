package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fraybet/cli/wallet"
)

// testKeystore creates a temp keystore with one wallet named "test" and returns
// the --keystore flag pair so tx commands can resolve a signer offline.
func testKeystore(t *testing.T) []string {
	t.Helper()
	dir := t.TempDir()
	st, err := wallet.NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := st.Create("test"); err != nil {
		t.Fatal(err)
	}
	return []string{"--keystore", dir, "--wallet", "test", "--unsigned"}
}

const (
	addrYes     = "0x1111111111111111111111111111111111111111"
	addrNo      = "0x2222222222222222222222222222222222222222"
	addrToken   = "0x036CbD53842c5426634e7929541eC2318f3dCF7e"
	addrFactory = "0x4444444444444444444444444444444444444444"
	addrEscrow  = "0x5555555555555555555555555555555555555555"
)

func draftArgs() []string {
	return []string{
		"--yes", addrYes, "--no", addrNo, "--token", addrToken,
		"--yes-stake", "500000000", "--no-stake", "500000000",
		"--statement", "BTC above 120000 on 2026-07-01", "--source", "Coinbase",
		"--event-time", "1782518400", "--nonce", "1", "--public",
	}
}

func TestBetDraft(t *testing.T) {
	var buf bytes.Buffer
	if err := runBet(append([]string{"draft"}, draftArgs()...), &buf); err != nil {
		t.Fatalf("draft: %v", err)
	}
	var d draftJSON
	if err := json.Unmarshal(buf.Bytes(), &d); err != nil {
		t.Fatalf("output not JSON: %v\n%s", err, buf.String())
	}
	if d.Visibility != "public" {
		t.Errorf("visibility = %s", d.Visibility)
	}
	if len(d.TermsHash) != 66 || !strings.HasPrefix(d.TermsHash, "0x") {
		t.Errorf("bad termsHash %q", d.TermsHash)
	}
	if d.Terms.ChallengeWindow != 86400 {
		t.Errorf("default challenge window not applied: %d", d.Terms.ChallengeWindow)
	}
}

func TestBetDraftRejectsMissingAgent(t *testing.T) {
	var buf bytes.Buffer
	err := runBet([]string{"draft", "--no", addrNo, "--token", addrToken,
		"--yes-stake", "1", "--no-stake", "1", "--statement", "x", "--source", "s",
		"--event-time", "1", "--nonce", "1"}, &buf)
	if err == nil {
		t.Fatal("expected error for missing --yes")
	}
}

func TestBetCreateProducesUnsignedTx(t *testing.T) {
	var buf bytes.Buffer
	args := append([]string{"create", "--factory", addrFactory}, draftArgs()...)
	args = append(args, testKeystore(t)...)
	if err := runBet(args, &buf); err != nil {
		t.Fatalf("create: %v", err)
	}
	var env txEnvelope
	if err := json.Unmarshal(buf.Bytes(), &env); err != nil {
		t.Fatalf("output not JSON: %v\n%s", err, buf.String())
	}
	if !env.DryRun {
		t.Error("expected dryRun=true")
	}
	if env.TermsHash == "" {
		t.Error("expected termsHash on create")
	}
}

func TestBetClaim(t *testing.T) {
	var buf bytes.Buffer
	args := append([]string{"claim", "--escrow", addrEscrow, "--outcome", "YES"}, testKeystore(t)...)
	err := runBet(args, &buf)
	if err != nil {
		t.Fatalf("claim: %v", err)
	}
	var env txEnvelope
	json.Unmarshal(buf.Bytes(), &env)
	if env.Action != "claim YES" {
		t.Errorf("action = %q", env.Action)
	}
}

func TestBetFund(t *testing.T) {
	var buf bytes.Buffer
	args := append([]string{"fund", "--escrow", addrEscrow}, testKeystore(t)...)
	if err := runBet(args, &buf); err != nil {
		t.Fatalf("fund: %v", err)
	}
}

func TestBetClone(t *testing.T) {
	// Produce a source draft and write it to a file.
	var src bytes.Buffer
	if err := runBet(append([]string{"draft"}, draftArgs()...), &src); err != nil {
		t.Fatal(err)
	}
	var srcDraft draftJSON
	json.Unmarshal(src.Bytes(), &srcDraft)

	path := filepath.Join(t.TempDir(), "terms.json")
	if err := os.WriteFile(path, src.Bytes(), 0o600); err != nil {
		t.Fatal(err)
	}

	newYes := "0xcccccccccccccccccccccccccccccccccccccccc"
	newNo := "0xdddddddddddddddddddddddddddddddddddddddd"
	var out bytes.Buffer
	err := runBet([]string{"clone", "--from-terms", path, "--yes", newYes, "--no", newNo, "--nonce", "7", "--public"}, &out)
	if err != nil {
		t.Fatalf("clone: %v", err)
	}
	var d draftJSON
	if err := json.Unmarshal(out.Bytes(), &d); err != nil {
		t.Fatalf("clone output not JSON: %v\n%s", err, out.String())
	}
	if d.Terms.YesAgent != newYes || d.Terms.NoAgent != newNo || d.Terms.Nonce != "7" {
		t.Errorf("clone identity wrong: yes=%s no=%s nonce=%s", d.Terms.YesAgent, d.Terms.NoAgent, d.Terms.Nonce)
	}
	if d.Terms.Statement != srcDraft.Terms.Statement {
		t.Error("clone did not preserve the statement")
	}
	if d.TermsHash == srcDraft.TermsHash {
		t.Error("clone hash collides with source")
	}
}
