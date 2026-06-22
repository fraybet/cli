// Package challenge builds shareable one-sided bet proposals — "challenge links".
//
// A challenger commits to the bet terms and the side they will take, but leaves
// the counterparty open. The proposal is encoded into a compact, URL-safe token
// that can be posted anywhere (e.g. on X): whoever opens the link accepts it by
// supplying their own address, which yields a complete, validated bets.Draft
// ready to create on-chain via the factory.
//
// Like the rest of the protocol this is non-custodial. The link is just an
// offer — nothing is escrowed by sharing one, and the platform never holds keys
// or funds. The real commitment is funding the BetEscrow after the bet is
// created. An Open carries an optional Expiry so a stale challenge stops being
// actionable on its own terms.
package challenge

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"strings"

	"github.com/fraybet/cli/bets"
	"github.com/fraybet/cli/core"
)

// Side is the position the proposer takes; the accepter takes the other one.
type Side uint8

const (
	SideYes Side = iota // proposer bets YES, accepter bets NO
	SideNo              // proposer bets NO, accepter bets YES
)

func (s Side) String() string {
	switch s {
	case SideYes:
		return "YES"
	case SideNo:
		return "NO"
	default:
		return fmt.Sprintf("Side(%d)", uint8(s))
	}
}

// ParseSide reads "YES"/"NO" (case-insensitive).
func ParseSide(s string) (Side, error) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "YES":
		return SideYes, nil
	case "NO":
		return SideNo, nil
	default:
		return SideYes, fmt.Errorf("challenge: side must be YES or NO, got %q", s)
	}
}

// Other returns the side the accepter takes.
func (s Side) Other() Side {
	if s == SideYes {
		return SideNo
	}
	return SideYes
}

// Open is a one-sided, counterparty-less bet proposal: everything except who
// takes the other side. Accepting fills in the counterparty and yields a
// complete bets.Draft.
type Open struct {
	Proposer        core.Address // the challenger's wallet (their side)
	Side            Side         // the side the proposer takes
	Handle          string       // optional display name / X handle
	CollateralToken core.Address
	ProposerStake   *big.Int        // the proposer's stake (base units)
	CounterStake    *big.Int        // the accepter's stake; nil/zero => match the proposer's
	Statement       string          // the bet question
	PrimarySource   string          // primary resolution source
	FallbackSource  string          // fallback resolution source
	EventTime       uint64          // event time (unix seconds)
	ClaimDeadline   uint64          // 0 => derived on accept (eventTime + grace)
	ChallengeWindow uint64          // 0 => default on accept
	Arbiter         core.Address    // zero = no arbiter (no dispute path)
	Nonce           *big.Int        // disambiguation nonce; nil => 0
	Visibility      core.Visibility // public bets are discoverable/cloneable
	Expiry          uint64          // unix seconds; 0 = no expiry
}

// ResolvedCounterStake is the accepter's stake, defaulting to the proposer's
// (even-money "I bet N, who'll match?").
func (o Open) ResolvedCounterStake() *big.Int {
	if o.CounterStake != nil && o.CounterStake.Sign() > 0 {
		return o.CounterStake
	}
	return o.ProposerStake
}

// Validate checks the proposal is well-formed enough to encode and share. The
// full bet is validated when accepted (Accept calls bets.NewDraft).
func (o Open) Validate() error {
	if o.Proposer == (core.Address{}) {
		return fmt.Errorf("challenge: proposer address is required")
	}
	if o.Side != SideYes && o.Side != SideNo {
		return fmt.Errorf("challenge: invalid side %d", o.Side)
	}
	if o.CollateralToken == (core.Address{}) {
		return fmt.Errorf("challenge: collateral token is required")
	}
	if o.ProposerStake == nil || o.ProposerStake.Sign() <= 0 {
		return fmt.Errorf("challenge: proposer stake must be > 0")
	}
	if o.CounterStake != nil && o.CounterStake.Sign() < 0 {
		return fmt.Errorf("challenge: counter stake must not be negative")
	}
	if strings.TrimSpace(o.Statement) == "" {
		return fmt.Errorf("challenge: statement is required")
	}
	if strings.TrimSpace(o.PrimarySource) == "" {
		return fmt.Errorf("challenge: a primary resolution source is required")
	}
	// Guarantee the link is actually acceptable: the terms must form a valid bet
	// once a counterparty is supplied (catches e.g. a missing eventTime up front,
	// rather than producing a dead link that fails only when someone accepts it).
	if _, err := o.draft(placeholderCounterparty(o.Proposer)); err != nil {
		return fmt.Errorf("challenge: terms do not form a valid bet: %w", err)
	}
	return nil
}

// placeholderCounterparty returns a non-zero address distinct from the proposer,
// used to dry-run validation before a real accepter is known.
func placeholderCounterparty(proposer core.Address) core.Address {
	var a core.Address
	a[19] = 0x01
	if a == proposer {
		a[19] = 0x02
	}
	return a
}

// IsExpired reports whether the proposal's own deadline has passed.
func (o Open) IsExpired(now uint64) bool {
	return o.Expiry != 0 && now >= o.Expiry
}

// Accept fills in the counterparty and returns a complete, validated draft.
// Side mapping: the proposer's stake/address land on their chosen side, the
// counterparty on the other. Defaults (challenge window, claim deadline, nonce)
// are applied by bets.NewDraft so the on-chain terms are produced one way.
func (o Open) Accept(counterparty core.Address) (bets.Draft, error) {
	if err := o.Validate(); err != nil {
		return bets.Draft{}, err
	}
	return o.draft(counterparty)
}

// draft maps the proposer/counterparty onto YES/NO and produces the validated
// bet. It does NOT call Validate (Validate uses it for a dry run, so calling
// back would recurse).
func (o Open) draft(counterparty core.Address) (bets.Draft, error) {
	if counterparty == (core.Address{}) {
		return bets.Draft{}, fmt.Errorf("challenge: accepter address is required")
	}
	if counterparty == o.Proposer {
		return bets.Draft{}, fmt.Errorf("challenge: accepter cannot be the proposer")
	}

	pStake := o.ProposerStake
	cStake := o.ResolvedCounterStake()

	in := bets.DraftInput{
		CollateralToken: o.CollateralToken,
		Arbiter:         o.Arbiter,
		Statement:       o.Statement,
		PrimarySource:   o.PrimarySource,
		FallbackSource:  o.FallbackSource,
		EventTime:       o.EventTime,
		ClaimDeadline:   o.ClaimDeadline,
		ChallengeWindow: o.ChallengeWindow,
		Nonce:           o.Nonce,
		Visibility:      o.Visibility,
	}
	if o.Side == SideYes {
		in.YesAgent, in.NoAgent = o.Proposer, counterparty
		in.YesStake, in.NoStake = pStake, cStake
	} else {
		in.NoAgent, in.YesAgent = o.Proposer, counterparty
		in.NoStake, in.YesStake = pStake, cStake
	}
	return bets.NewDraft(in)
}

// ID is a stable identifier for the proposal (keccak of its canonical encoding).
// It is NOT the final termsHash — that depends on the counterparty and is only
// known once the link is accepted.
func (o Open) ID() (core.Hash32, error) {
	b, err := json.Marshal(o.toWire())
	if err != nil {
		return core.Hash32{}, err
	}
	return core.Keccak256(b), nil
}

// --- encoding: compact, URL-safe token ---

const wireVersion = 1

type wire struct {
	V               int    `json:"v"`
	Proposer        string `json:"p"`
	Side            string `json:"side"`
	Handle          string `json:"h,omitempty"`
	Token           string `json:"tok"`
	ProposerStake   string `json:"ps"`
	CounterStake    string `json:"cs,omitempty"`
	Statement       string `json:"q"`
	PrimarySource   string `json:"src,omitempty"`
	FallbackSource  string `json:"fb,omitempty"`
	EventTime       uint64 `json:"et,omitempty"`
	ClaimDeadline   uint64 `json:"cd,omitempty"`
	ChallengeWindow uint64 `json:"cw,omitempty"`
	Arbiter         string `json:"arb,omitempty"`
	Nonce           string `json:"n,omitempty"`
	Visibility      string `json:"vis,omitempty"`
	Expiry          uint64 `json:"exp,omitempty"`
}

func (o Open) toWire() wire {
	w := wire{
		V:               wireVersion,
		Proposer:        o.Proposer.Hex(),
		Side:            o.Side.String(),
		Handle:          o.Handle,
		Token:           o.CollateralToken.Hex(),
		ProposerStake:   bigOrZero(o.ProposerStake),
		Statement:       o.Statement,
		PrimarySource:   o.PrimarySource,
		FallbackSource:  o.FallbackSource,
		EventTime:       o.EventTime,
		ClaimDeadline:   o.ClaimDeadline,
		ChallengeWindow: o.ChallengeWindow,
		Visibility:      o.Visibility.String(),
		Expiry:          o.Expiry,
	}
	if o.CounterStake != nil && o.CounterStake.Sign() > 0 {
		w.CounterStake = o.CounterStake.String()
	}
	if o.Arbiter != (core.Address{}) {
		w.Arbiter = o.Arbiter.Hex()
	}
	if o.Nonce != nil && o.Nonce.Sign() != 0 {
		w.Nonce = o.Nonce.String()
	}
	return w
}

// Encode serializes the proposal into a compact, URL-safe token (base64url of
// the canonical JSON). The token is self-contained: no server state is needed
// to decode and act on it.
func (o Open) Encode() (string, error) {
	if err := o.Validate(); err != nil {
		return "", err
	}
	b, err := json.Marshal(o.toWire())
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// Decode parses a token produced by Encode back into an Open.
func Decode(token string) (Open, error) {
	token = strings.TrimSpace(token)
	if token == "" {
		return Open{}, fmt.Errorf("challenge: empty token")
	}
	raw, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return Open{}, fmt.Errorf("challenge: token is not valid base64url: %w", err)
	}
	var w wire
	if err := json.Unmarshal(raw, &w); err != nil {
		return Open{}, fmt.Errorf("challenge: token payload is not valid: %w", err)
	}
	if w.V != wireVersion {
		return Open{}, fmt.Errorf("challenge: unsupported token version %d (want %d)", w.V, wireVersion)
	}
	return w.toOpen()
}

func (w wire) toOpen() (Open, error) {
	proposer, err := core.HexToAddress(w.Proposer)
	if err != nil {
		return Open{}, fmt.Errorf("challenge: proposer: %w", err)
	}
	token, err := core.HexToAddress(w.Token)
	if err != nil {
		return Open{}, fmt.Errorf("challenge: token: %w", err)
	}
	side, err := ParseSide(w.Side)
	if err != nil {
		return Open{}, err
	}
	ps, err := parseBig(w.ProposerStake)
	if err != nil {
		return Open{}, fmt.Errorf("challenge: proposerStake: %w", err)
	}
	o := Open{
		Proposer:        proposer,
		Side:            side,
		Handle:          w.Handle,
		CollateralToken: token,
		ProposerStake:   ps,
		Statement:       w.Statement,
		PrimarySource:   w.PrimarySource,
		FallbackSource:  w.FallbackSource,
		EventTime:       w.EventTime,
		ClaimDeadline:   w.ClaimDeadline,
		ChallengeWindow: w.ChallengeWindow,
		Expiry:          w.Expiry,
	}
	if w.CounterStake != "" {
		cs, err := parseBig(w.CounterStake)
		if err != nil {
			return Open{}, fmt.Errorf("challenge: counterStake: %w", err)
		}
		o.CounterStake = cs
	}
	if w.Arbiter != "" {
		arb, err := core.HexToAddress(w.Arbiter)
		if err != nil {
			return Open{}, fmt.Errorf("challenge: arbiter: %w", err)
		}
		o.Arbiter = arb
	}
	if w.Nonce != "" {
		n, err := parseBig(w.Nonce)
		if err != nil {
			return Open{}, fmt.Errorf("challenge: nonce: %w", err)
		}
		o.Nonce = n
	}
	if strings.EqualFold(w.Visibility, "public") {
		o.Visibility = core.VisibilityPublic
	}
	return o, nil
}

// Link builds a shareable URL for the proposal under baseURL, of the form
// <baseURL>/challenge?p=<token>. baseURL may include a trailing slash or not.
func (o Open) Link(baseURL string) (string, error) {
	tok, err := o.Encode()
	if err != nil {
		return "", err
	}
	base := strings.TrimRight(strings.TrimSpace(baseURL), "/")
	return base + "/challenge?p=" + url.QueryEscape(tok), nil
}

func bigOrZero(b *big.Int) string {
	if b == nil {
		return "0"
	}
	return b.String()
}

func parseBig(s string) (*big.Int, error) {
	n, ok := new(big.Int).SetString(strings.TrimSpace(s), 10)
	if !ok {
		return nil, fmt.Errorf("%q is not a base-10 integer", s)
	}
	return n, nil
}
