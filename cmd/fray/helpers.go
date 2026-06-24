package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"

	"github.com/fraybet/cli/bets"
	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
)

func parseAddr(s, name string) (core.Address, error) {
	if strings.TrimSpace(s) == "" {
		return core.Address{}, fmt.Errorf("--%s is required", name)
	}
	return core.HexToAddress(s)
}

func parseOptAddr(s string) (core.Address, error) {
	if strings.TrimSpace(s) == "" {
		return core.Address{}, nil // zero address = no arbiter
	}
	return core.HexToAddress(s)
}

func parseBig(s, name string) (*big.Int, error) {
	if strings.TrimSpace(s) == "" {
		return nil, fmt.Errorf("--%s is required", name)
	}
	n, ok := new(big.Int).SetString(strings.TrimSpace(s), 10)
	if !ok {
		return nil, fmt.Errorf("--%s %q is not a base-10 integer", name, s)
	}
	return n, nil
}

func parseOutcome(s string) (bets.Outcome, error) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "YES":
		return bets.OutcomeYes, nil
	case "NO":
		return bets.OutcomeNo, nil
	default:
		return bets.OutcomeUnresolved, fmt.Errorf("--outcome must be YES or NO, got %q", s)
	}
}

// parseArbiterOutcome also accepts VOID — the arbiter may void a contested bet
// (refund both sides) when neither YES nor NO can be established.
func parseArbiterOutcome(s string) (bets.Outcome, error) {
	if strings.EqualFold(strings.TrimSpace(s), "VOID") {
		return bets.OutcomeVoid, nil
	}
	return parseOutcome(s)
}

// parseSideYN returns true for YES, false for NO.
func parseSideYN(s string) (bool, error) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "YES":
		return true, nil
	case "NO":
		return false, nil
	default:
		return false, fmt.Errorf("--side must be YES or NO, got %q", s)
	}
}

func writeJSON(out io.Writer, v any) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

type termsJSON struct {
	YesAgent        string `json:"yesAgent"`
	NoAgent         string `json:"noAgent"`
	CollateralToken string `json:"collateralToken"`
	Arbiter         string `json:"arbiter"`
	YesStake        string `json:"yesStake"`
	NoStake         string `json:"noStake"`
	Statement       string `json:"statement"`
	EventTime       uint64 `json:"eventTime"`
	ClaimDeadline   uint64 `json:"claimDeadline"`
	ChallengeWindow uint64 `json:"challengeWindow"`
	PrimarySource   string `json:"primarySource"`
	FallbackSource  string `json:"fallbackSource"`
	Nonce           string `json:"nonce"`
}

type draftJSON struct {
	TermsHash  string    `json:"termsHash"`
	Visibility string    `json:"visibility"`
	Terms      termsJSON `json:"terms"`
}

func renderDraft(d bets.Draft) draftJSON {
	t := d.Terms
	return draftJSON{
		TermsHash:  d.TermsHash().Hex(),
		Visibility: d.Visibility.String(),
		Terms: termsJSON{
			YesAgent:        t.YesAgent.Hex(),
			NoAgent:         t.NoAgent.Hex(),
			CollateralToken: t.CollateralToken.Hex(),
			Arbiter:         t.Arbiter.Hex(),
			YesStake:        t.YesStake.String(),
			NoStake:         t.NoStake.String(),
			Statement:       t.Statement,
			EventTime:       t.EventTime,
			ClaimDeadline:   t.ClaimDeadline,
			ChallengeWindow: t.ChallengeWindow,
			PrimarySource:   t.PrimarySource,
			FallbackSource:  t.FallbackSource,
			Nonce:           t.Nonce.String(),
		},
	}
}

// txEnvelope is the safe-by-default CLI output for an economic action: an
// UNSIGNED transaction the caller must sign + broadcast themselves (design
// §17.15). No broadcast happens here.
type txEnvelope struct {
	Action     string           `json:"action"`
	DryRun     bool             `json:"dryRun"`
	TermsHash  string           `json:"termsHash,omitempty"`
	UnsignedTx chain.UnsignedTx `json:"unsignedTx"`
	Note       string           `json:"note"`
}

// loadTerms reads a terms JSON file (as produced by `bet draft`) into BetTerms.
func loadTerms(path string) (core.BetTerms, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return core.BetTerms{}, err
	}
	var dj draftJSON
	if err := json.Unmarshal(data, &dj); err != nil {
		return core.BetTerms{}, fmt.Errorf("parse terms file: %w", err)
	}
	return termsFromJSON(dj.Terms)
}

func termsFromJSON(tj termsJSON) (core.BetTerms, error) {
	yes, err := core.HexToAddress(tj.YesAgent)
	if err != nil {
		return core.BetTerms{}, fmt.Errorf("yesAgent: %w", err)
	}
	no, err := core.HexToAddress(tj.NoAgent)
	if err != nil {
		return core.BetTerms{}, fmt.Errorf("noAgent: %w", err)
	}
	token, err := core.HexToAddress(tj.CollateralToken)
	if err != nil {
		return core.BetTerms{}, fmt.Errorf("collateralToken: %w", err)
	}
	arb, err := core.HexToAddress(tj.Arbiter)
	if err != nil {
		return core.BetTerms{}, fmt.Errorf("arbiter: %w", err)
	}
	ys, err := parseBig(tj.YesStake, "yesStake")
	if err != nil {
		return core.BetTerms{}, err
	}
	ns, err := parseBig(tj.NoStake, "noStake")
	if err != nil {
		return core.BetTerms{}, err
	}
	nonce, err := parseBig(tj.Nonce, "nonce")
	if err != nil {
		return core.BetTerms{}, err
	}
	return core.BetTerms{
		YesAgent: yes, NoAgent: no, CollateralToken: token, Arbiter: arb,
		YesStake: ys, NoStake: ns,
		Statement: tj.Statement, EventTime: tj.EventTime, ClaimDeadline: tj.ClaimDeadline,
		ChallengeWindow: tj.ChallengeWindow, PrimarySource: tj.PrimarySource,
		FallbackSource: tj.FallbackSource, Nonce: nonce,
	}, nil
}

func writeUnsignedTx(out io.Writer, action string, tx chain.UnsignedTx, termsHash string) error {
	return writeJSON(out, txEnvelope{
		Action:     action,
		DryRun:     true,
		TermsHash:  termsHash,
		UnsignedTx: tx,
		Note:       "UNSIGNED — review the calldata, then sign and broadcast with your own wallet. The platform never signs or holds keys.",
	})
}
