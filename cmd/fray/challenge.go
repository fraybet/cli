package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/fraybet/cli/challenge"
	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/txbuild"
)

// defaultBaseURL is where challenge links point unless --base-url or the
// FRAY_BASE_URL env var overrides it. Defaults to the live site so a
// generated link is shareable out of the box.
const defaultBaseURL = "https://fray.bet"

type proposeLinkOutput struct {
	ProposalID      string `json:"proposalId"`
	Link            string `json:"link"`
	Token           string `json:"token"`
	ProposerTakes   string `json:"proposerTakes"`
	AccepterTakes   string `json:"accepterTakes"`
	Proposer        string `json:"proposer"`
	Handle          string `json:"handle,omitempty"`
	Statement       string `json:"statement"`
	CollateralToken string `json:"collateralToken"`
	ProposerStake   string `json:"proposerStake"`
	CounterStake    string `json:"counterStake"`
	Arbiter         string `json:"arbiter,omitempty"`
	Visibility      string `json:"visibility"`
	Expiry          uint64 `json:"expiry,omitempty"`
	Note            string `json:"note"`
}

// betProposeLink builds a one-sided challenge and prints a shareable link.
func betProposeLink(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("propose-link", flag.ContinueOnError)
	fs.SetOutput(out)
	proposer := fs.String("proposer", "", "your wallet address (the side you take)")
	side := fs.String("side", "", "the side YOU take: YES or NO")
	handle := fs.String("handle", "", "optional display name / X handle")
	token := fs.String("token", "", "collateral token address (USDC)")
	stake := fs.String("stake", "", "your stake in base units (USDC 6dp)")
	counterStake := fs.String("counter-stake", "", "the taker's stake (default: match yours)")
	statement := fs.String("statement", "", "the bet statement")
	source := fs.String("source", "", "primary resolution source")
	fallback := fs.String("fallback", "", "fallback resolution source")
	eventTime := fs.Uint64("event-time", 0, "event time (unix seconds)")
	claimDeadline := fs.Uint64("claim-deadline", 0, "claim deadline (unix; default event+24h)")
	challengeWindow := fs.Uint64("challenge-window", 0, "challenge window seconds (default 24h)")
	arbiter := fs.String("arbiter", "", "arbiter address (optional, enables disputes)")
	nonce := fs.String("nonce", "0", "disambiguation nonce")
	public := fs.Bool("public", false, "mark the resulting bet public (discoverable)")
	expiry := fs.Uint64("expiry", 0, "link expiry (unix seconds; 0 = never)")
	baseURL := fs.String("base-url", "", "site base URL for the link (env FRAY_BASE_URL)")
	if err := fs.Parse(args); err != nil {
		return err
	}

	proposerA, err := parseAddr(*proposer, "proposer")
	if err != nil {
		return err
	}
	tokenA, err := parseAddr(*token, "token")
	if err != nil {
		return err
	}
	sd, err := challenge.ParseSide(*side)
	if err != nil {
		return err
	}
	stakeB, err := parseBig(*stake, "stake")
	if err != nil {
		return err
	}
	nonceB, err := parseBig(*nonce, "nonce")
	if err != nil {
		return err
	}
	arbA, err := parseOptAddr(*arbiter)
	if err != nil {
		return err
	}

	o := challenge.Open{
		Proposer:        proposerA,
		Side:            sd,
		Handle:          *handle,
		CollateralToken: tokenA,
		ProposerStake:   stakeB,
		Statement:       *statement,
		PrimarySource:   *source,
		FallbackSource:  *fallback,
		EventTime:       *eventTime,
		ClaimDeadline:   *claimDeadline,
		ChallengeWindow: *challengeWindow,
		Arbiter:         arbA,
		Nonce:           nonceB,
		Visibility:      core.VisibilityPrivate,
		Expiry:          *expiry,
	}
	if *public {
		o.Visibility = core.VisibilityPublic
	}
	if strings.TrimSpace(*counterStake) != "" {
		cs, err := parseBig(*counterStake, "counter-stake")
		if err != nil {
			return err
		}
		o.CounterStake = cs
	}
	if err := o.Validate(); err != nil {
		return err
	}

	tok, err := o.Encode()
	if err != nil {
		return err
	}
	id, err := o.ID()
	if err != nil {
		return err
	}
	link, err := o.Link(resolveBaseURL(*baseURL))
	if err != nil {
		return err
	}

	res := proposeLinkOutput{
		ProposalID:      id.Hex(),
		Link:            link,
		Token:           tok,
		ProposerTakes:   sd.String(),
		AccepterTakes:   sd.Other().String(),
		Proposer:        proposerA.Hex(),
		Handle:          *handle,
		Statement:       *statement,
		CollateralToken: tokenA.Hex(),
		ProposerStake:   o.ProposerStake.String(),
		CounterStake:    o.ResolvedCounterStake().String(),
		Visibility:      o.Visibility.String(),
		Expiry:          *expiry,
		Note: "Share this link. Anyone can accept by taking the " + sd.Other().String() +
			" side; the bet only commits once both sides fund the escrow on-chain.",
	}
	if arbA != (core.Address{}) {
		res.Arbiter = arbA.Hex()
	}
	return writeJSON(out, res)
}

// betAcceptLink resolves a challenge link as the counterparty, producing the
// full draft and (optionally) the unsigned create tx.
func betAcceptLink(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("accept-link", flag.ContinueOnError)
	fs.SetOutput(out)
	tokenFlag := fs.String("token", "", "the challenge token (from `propose-link`)")
	linkFlag := fs.String("link", "", "a full challenge URL (token is read from ?p=)")
	as := fs.String("as", "", "your wallet address (you take the open side)")
	from := fs.String("from", "", "tx sender for create (default: --as)")
	factory := fs.String("factory", "", "BetEscrowFactory address (emit the create tx)")
	if err := fs.Parse(args); err != nil {
		return err
	}

	tok, err := tokenFromFlags(*tokenFlag, *linkFlag)
	if err != nil {
		return err
	}
	o, err := challenge.Decode(tok)
	if err != nil {
		return err
	}
	asA, err := parseAddr(*as, "as")
	if err != nil {
		return err
	}

	d, err := o.AcceptAt(asA, uint64(time.Now().Unix()))
	if err != nil {
		return err
	}

	// No factory => just hand back the resolved, validated draft.
	if strings.TrimSpace(*factory) == "" {
		return writeJSON(out, renderDraft(d))
	}

	fromStr := *from
	if strings.TrimSpace(fromStr) == "" {
		fromStr = *as
	}
	fromA, err := parseAddr(fromStr, "from")
	if err != nil {
		return err
	}
	factoryA, err := parseAddr(*factory, "factory")
	if err != nil {
		return err
	}
	tx, err := txbuild.CreateBet(fromA, factoryA, d)
	if err != nil {
		return err
	}
	return writeUnsignedTx(out, "accept challenge + create bet", tx, d.TermsHash().Hex())
}

// tokenFromFlags returns the token from --token, or extracts it from --link's
// ?p= query parameter.
func tokenFromFlags(tokenFlag, linkFlag string) (string, error) {
	if strings.TrimSpace(tokenFlag) != "" {
		return strings.TrimSpace(tokenFlag), nil
	}
	if strings.TrimSpace(linkFlag) == "" {
		return "", fmt.Errorf("provide --token or --link")
	}
	u, err := url.Parse(strings.TrimSpace(linkFlag))
	if err != nil {
		return "", fmt.Errorf("invalid --link: %w", err)
	}
	p := u.Query().Get("p")
	if p == "" {
		return "", fmt.Errorf("--link has no ?p= challenge token")
	}
	return p, nil
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
