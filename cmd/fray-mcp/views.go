package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/fraybet/cli/bets"
	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
)

// unsignedTxNote is attached to every transaction-producing tool result. The
// MCP server is non-custodial: it returns calldata for the agent to sign and
// broadcast with its own wallet. The server never signs or holds keys.
const unsignedTxNote = "UNSIGNED — review the calldata, then sign and broadcast with your own wallet. Fray never signs or holds keys."

// --- structured output views (plain string fields so the inferred JSON schema
// matches the emitted JSON; chain.UnsignedTx has a custom marshaler that would
// otherwise diverge from its struct-derived schema). ---

type termsView struct {
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

type draftView struct {
	TermsHash  string    `json:"termsHash"`
	Visibility string    `json:"visibility"`
	Terms      termsView `json:"terms"`
}

func renderDraft(d bets.Draft) draftView {
	t := d.Terms
	return draftView{
		TermsHash:  d.TermsHash().Hex(),
		Visibility: d.Visibility.String(),
		Terms: termsView{
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

type unsignedTxView struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Value   string `json:"value"`
	Data    string `json:"data"`
	Gas     string `json:"gas"`
	Nonce   string `json:"nonce"`
	ChainID string `json:"chainId"`
}

// toTxView reuses chain.UnsignedTx's own hex wire format (its MarshalJSON)
// rather than re-deriving the encoding here.
func toTxView(t chain.UnsignedTx) (unsignedTxView, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return unsignedTxView{}, err
	}
	var v unsignedTxView
	if err := json.Unmarshal(b, &v); err != nil {
		return unsignedTxView{}, err
	}
	return v, nil
}

type txEnvelope struct {
	Action     string         `json:"action"`
	TermsHash  string         `json:"termsHash,omitempty"`
	UnsignedTx unsignedTxView `json:"unsignedTx"`
	Note       string         `json:"note"`
}

func envelope(action, termsHash string, t chain.UnsignedTx) (txEnvelope, error) {
	v, err := toTxView(t)
	if err != nil {
		return txEnvelope{}, err
	}
	return txEnvelope{Action: action, TermsHash: termsHash, UnsignedTx: v, Note: unsignedTxNote}, nil
}

// --- parse helpers (string args -> core types) ---

func mustAddr(s, name string) (core.Address, error) {
	if strings.TrimSpace(s) == "" {
		return core.Address{}, fmt.Errorf("%s is required", name)
	}
	return core.HexToAddress(s)
}

func optAddr(s string) (core.Address, error) {
	if strings.TrimSpace(s) == "" {
		return core.Address{}, nil
	}
	return core.HexToAddress(s)
}

func mustBig(s, name string) (*big.Int, error) {
	if strings.TrimSpace(s) == "" {
		return nil, fmt.Errorf("%s is required", name)
	}
	n, ok := new(big.Int).SetString(strings.TrimSpace(s), 10)
	if !ok {
		return nil, fmt.Errorf("%s %q is not a base-10 integer", name, s)
	}
	return n, nil
}

func bigOrZero(s string) (*big.Int, error) {
	if strings.TrimSpace(s) == "" {
		return big.NewInt(0), nil
	}
	return mustBig(s, "value")
}

func parseOutcome(s string) (bets.Outcome, error) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "YES":
		return bets.OutcomeYes, nil
	case "NO":
		return bets.OutcomeNo, nil
	default:
		return bets.OutcomeUnresolved, fmt.Errorf("outcome must be YES or NO, got %q", s)
	}
}
