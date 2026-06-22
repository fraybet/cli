package bets

import (
	"math/big"
	"testing"

	"github.com/fraybet/cli/core"
)

func sampleInput() DraftInput {
	return DraftInput{
		YesAgent:        core.MustHexToAddress("0x1111111111111111111111111111111111111111"),
		NoAgent:         core.MustHexToAddress("0x2222222222222222222222222222222222222222"),
		CollateralToken: core.MustHexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e"),
		Arbiter:         core.MustHexToAddress("0x3333333333333333333333333333333333333333"),
		YesStake:        big.NewInt(500_000_000),
		NoStake:         big.NewInt(500_000_000),
		Statement:       "BTC-USD daily close on Coinbase will be above 120000 on 2026-07-01 UTC.",
		PrimarySource:   "Coinbase BTC-USD daily candle",
		FallbackSource:  "Chainlink BTC/USD reference feed",
		EventTime:       1782518400,
		Visibility:      core.VisibilityPublic,
	}
}

func TestNewDraftAppliesDefaults(t *testing.T) {
	d, err := NewDraft(sampleInput())
	if err != nil {
		t.Fatalf("NewDraft: %v", err)
	}
	if d.Terms.ChallengeWindow != DefaultChallengeWindow {
		t.Errorf("challengeWindow = %d, want default %d", d.Terms.ChallengeWindow, DefaultChallengeWindow)
	}
	if d.Terms.ClaimDeadline != 1782518400+DefaultClaimGrace {
		t.Errorf("claimDeadline = %d, want eventTime+grace", d.Terms.ClaimDeadline)
	}
	if d.Terms.Nonce == nil || d.Terms.Nonce.Sign() != 0 {
		t.Errorf("nonce = %v, want 0", d.Terms.Nonce)
	}
	if d.Visibility != core.VisibilityPublic {
		t.Errorf("visibility = %s, want public", d.Visibility)
	}
}

func TestNewDraftValidates(t *testing.T) {
	in := sampleInput()
	in.Statement = ""
	if _, err := NewDraft(in); err == nil {
		t.Fatal("expected validation error for empty statement")
	}
}

func TestClonePreservesTermsChangesIdentity(t *testing.T) {
	src, err := NewDraft(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	newYes := core.MustHexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	newNo := core.MustHexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")

	clone, err := Clone(src.Terms, newYes, newNo, big.NewInt(7), core.VisibilityPrivate)
	if err != nil {
		t.Fatalf("Clone: %v", err)
	}
	// Question + sources + timing + stakes preserved.
	if clone.Terms.Statement != src.Terms.Statement || clone.Terms.PrimarySource != src.Terms.PrimarySource {
		t.Error("clone did not preserve the question/source")
	}
	if clone.Terms.YesStake.Cmp(src.Terms.YesStake) != 0 {
		t.Error("clone did not preserve stake")
	}
	// Identity changed.
	if clone.Terms.YesAgent != newYes || clone.Terms.NoAgent != newNo || clone.Terms.Nonce.Int64() != 7 {
		t.Error("clone did not take the new agents/nonce")
	}
	if clone.Visibility != core.VisibilityPrivate {
		t.Error("clone visibility not applied")
	}
	// A different pair of agents => a different bet identity.
	if clone.TermsHash() == src.TermsHash() {
		t.Error("clone hash collides with source")
	}
}

func TestCloneRequiresNonce(t *testing.T) {
	src, _ := NewDraft(sampleInput())
	if _, err := Clone(src.Terms, src.Terms.YesAgent, src.Terms.NoAgent, nil, core.VisibilityPublic); err == nil {
		t.Fatal("expected error for nil nonce")
	}
}

func TestStatusOutcomeStrings(t *testing.T) {
	if StatusContested.String() != "contested" || OutcomeYes.String() != "YES" {
		t.Errorf("enum strings wrong: %s / %s", StatusContested, OutcomeYes)
	}
}
