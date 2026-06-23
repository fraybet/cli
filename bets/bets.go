// Package bets builds and reasons about bilateral bet terms (design §9). It is
// the domain layer over core.BetTerms: drafting with sensible defaults,
// validation, the public-bet clone flow (design §0.8), and Go mirrors of the
// on-chain lifecycle enums for read models. It deliberately depends only on
// internal/core — calldata/tx building lives in the wiring layer (bindings +
// internal/chain).
package bets

import (
	"fmt"
	"math/big"

	"github.com/fraybet/cli/core"
)

// Defaults applied by NewDraft when a field is left zero.
const (
	DefaultChallengeWindow uint64 = 86_400 // 24h (design §9.3 example)
	DefaultClaimGrace      uint64 = 86_400 // claimDeadline = eventTime + 24h
)

// Status mirrors BetEscrow.Status on-chain (for indexed read models).
type Status uint8

const (
	StatusFunding Status = iota
	StatusLive
	StatusClaimed
	StatusContested
	StatusResolved
	StatusVoided
)

func (s Status) String() string {
	switch s {
	case StatusFunding:
		return "funding"
	case StatusLive:
		return "live"
	case StatusClaimed:
		return "claimed"
	case StatusContested:
		return "contested"
	case StatusResolved:
		return "resolved"
	case StatusVoided:
		return "voided"
	default:
		return fmt.Sprintf("Status(%d)", uint8(s))
	}
}

// Outcome mirrors BetEscrow.Outcome on-chain.
type Outcome uint8

const (
	OutcomeUnresolved Outcome = iota
	OutcomeYes
	OutcomeNo
	OutcomeVoid
)

func (o Outcome) String() string {
	switch o {
	case OutcomeYes:
		return "YES"
	case OutcomeNo:
		return "NO"
	case OutcomeVoid:
		return "VOID"
	case OutcomeUnresolved:
		return "unresolved"
	default:
		return fmt.Sprintf("Outcome(%d)", uint8(o))
	}
}

// DraftInput is the caller's intent for a new bet.
type DraftInput struct {
	YesAgent        core.Address
	NoAgent         core.Address
	CollateralToken core.Address
	Arbiter         core.Address // zero = no arbiter
	YesStake        *big.Int
	NoStake         *big.Int
	Statement       string
	PrimarySource   string
	FallbackSource  string
	EventTime       uint64
	ClaimDeadline   uint64 // 0 => EventTime + DefaultClaimGrace
	ChallengeWindow uint64 // 0 => DefaultChallengeWindow
	Nonce           *big.Int
	Visibility      core.Visibility
}

// Draft is a validated, hashable bet ready to be created on-chain.
type Draft struct {
	Terms      core.BetTerms
	Visibility core.Visibility
}

// NewDraft builds and validates a Draft, applying defaults for the challenge
// window, claim deadline, and nonce.
func NewDraft(in DraftInput) (Draft, error) {
	cw := in.ChallengeWindow
	if cw == 0 {
		cw = DefaultChallengeWindow
	}
	cd := in.ClaimDeadline
	if cd == 0 && in.EventTime != 0 {
		cd = in.EventTime + DefaultClaimGrace
	}
	nonce := in.Nonce
	if nonce == nil {
		nonce = big.NewInt(0)
	}

	terms := core.BetTerms{
		YesAgent:        in.YesAgent,
		NoAgent:         in.NoAgent,
		CollateralToken: in.CollateralToken,
		YesStake:        in.YesStake,
		NoStake:         in.NoStake,
		Statement:       in.Statement,
		EventTime:       in.EventTime,
		ClaimDeadline:   cd,
		ChallengeWindow: cw,
		PrimarySource:   in.PrimarySource,
		FallbackSource:  in.FallbackSource,
		Arbiter:         in.Arbiter,
		Nonce:           nonce,
		Visibility:      uint8(in.Visibility),
	}
	if err := terms.Validate(); err != nil {
		return Draft{}, fmt.Errorf("bets: invalid draft: %w", err)
	}
	return Draft{Terms: terms, Visibility: in.Visibility}, nil
}

// NewOpenDraft builds a validated OPEN bet: the proposer takes one side and the
// other is left open (address(0)), filled later by accept(). proposerYes places
// the proposer on the YES side (NO open); otherwise the YES side is open. The
// proposer's address comes from the matching agent field of the input.
func NewOpenDraft(in DraftInput, proposerYes bool) (Draft, error) {
	cw := in.ChallengeWindow
	if cw == 0 {
		cw = DefaultChallengeWindow
	}
	cd := in.ClaimDeadline
	if cd == 0 && in.EventTime != 0 {
		cd = in.EventTime + DefaultClaimGrace
	}
	nonce := in.Nonce
	if nonce == nil {
		nonce = big.NewInt(0)
	}

	yes, no := in.YesAgent, in.NoAgent
	if proposerYes {
		no = core.Address{} // open NO side
	} else {
		yes = core.Address{} // open YES side
	}

	terms := core.BetTerms{
		YesAgent:        yes,
		NoAgent:         no,
		CollateralToken: in.CollateralToken,
		YesStake:        in.YesStake,
		NoStake:         in.NoStake,
		Statement:       in.Statement,
		EventTime:       in.EventTime,
		ClaimDeadline:   cd,
		ChallengeWindow: cw,
		PrimarySource:   in.PrimarySource,
		FallbackSource:  in.FallbackSource,
		Arbiter:         in.Arbiter,
		Nonce:           nonce,
		Visibility:      uint8(in.Visibility),
	}
	if err := terms.ValidateOpen(); err != nil {
		return Draft{}, fmt.Errorf("bets: invalid open draft: %w", err)
	}
	return Draft{Terms: terms, Visibility: in.Visibility}, nil
}

// TermsHash is the canonical bet identifier (EIP-712 struct hash).
func (d Draft) TermsHash() core.Hash32 { return d.Terms.TermsHash() }

// Clone produces a fresh draft from an existing (typically public) bet's terms
// for a new pair of agents, preserving the question, sources, timing, stakes,
// token, and arbiter (design §0.8 — popular public bets get cloned). The new
// agents and nonce must be supplied so the clone has its own identity.
func Clone(src core.BetTerms, yesAgent, noAgent core.Address, nonce *big.Int, visibility core.Visibility) (Draft, error) {
	if nonce == nil {
		return Draft{}, fmt.Errorf("bets: clone requires a nonce")
	}
	t := src // value copy
	t.YesAgent = yesAgent
	t.NoAgent = noAgent
	t.Nonce = nonce
	t.Visibility = uint8(visibility) // the clone's visibility is bound into its hash
	if err := t.Validate(); err != nil {
		return Draft{}, fmt.Errorf("bets: invalid clone: %w", err)
	}
	return Draft{Terms: t, Visibility: visibility}, nil
}
