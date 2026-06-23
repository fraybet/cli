package main

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/fraybet/cli/bets"
	"github.com/fraybet/cli/challenge"
	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/txbuild"
)

// errResult surfaces a tool error to the model (IsError + text) rather than as
// a protocol-level error, so the LLM can see it and self-correct.
func errResult[T any](err error) (*mcp.CallToolResult, T, error) {
	var zero T
	return &mcp.CallToolResult{
		IsError: true,
		Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
	}, zero, nil
}

// --- shared bet-term fields (used by draft + create) ---

type draftFields struct {
	Yes             string `json:"yes" jsonschema:"YES agent address (0x…)"`
	No              string `json:"no" jsonschema:"NO agent address (0x…)"`
	Token           string `json:"token" jsonschema:"collateral token address (USDC)"`
	Arbiter         string `json:"arbiter,omitempty" jsonschema:"optional arbiter address; enables the dispute path"`
	YesStake        string `json:"yesStake" jsonschema:"YES stake in base units (USDC has 6 decimals)"`
	NoStake         string `json:"noStake" jsonschema:"NO stake in base units"`
	Statement       string `json:"statement" jsonschema:"the bet question/statement"`
	Source          string `json:"source" jsonschema:"primary resolution source"`
	Fallback        string `json:"fallback,omitempty" jsonschema:"fallback resolution source"`
	EventTime       uint64 `json:"eventTime,omitempty" jsonschema:"event time (unix seconds)"`
	ClaimDeadline   uint64 `json:"claimDeadline,omitempty" jsonschema:"claim deadline (unix; default eventTime+24h)"`
	ChallengeWindow uint64 `json:"challengeWindow,omitempty" jsonschema:"challenge window seconds (default 24h)"`
	Nonce           string `json:"nonce,omitempty" jsonschema:"disambiguation nonce (default 0)"`
	Public          bool   `json:"public,omitempty" jsonschema:"mark the bet public (discoverable/cloneable)"`
}

func (f draftFields) input() (bets.DraftInput, error) {
	yes, err := mustAddr(f.Yes, "yes")
	if err != nil {
		return bets.DraftInput{}, err
	}
	no, err := mustAddr(f.No, "no")
	if err != nil {
		return bets.DraftInput{}, err
	}
	token, err := mustAddr(f.Token, "token")
	if err != nil {
		return bets.DraftInput{}, err
	}
	arb, err := optAddr(f.Arbiter)
	if err != nil {
		return bets.DraftInput{}, err
	}
	ys, err := mustBig(f.YesStake, "yesStake")
	if err != nil {
		return bets.DraftInput{}, err
	}
	ns, err := mustBig(f.NoStake, "noStake")
	if err != nil {
		return bets.DraftInput{}, err
	}
	nonce, err := bigOrZero(f.Nonce)
	if err != nil {
		return bets.DraftInput{}, err
	}
	vis := core.VisibilityPrivate
	if f.Public {
		vis = core.VisibilityPublic
	}
	return bets.DraftInput{
		YesAgent: yes, NoAgent: no, CollateralToken: token, Arbiter: arb,
		YesStake: ys, NoStake: ns,
		Statement: f.Statement, PrimarySource: f.Source, FallbackSource: f.Fallback,
		EventTime: f.EventTime, ClaimDeadline: f.ClaimDeadline, ChallengeWindow: f.ChallengeWindow,
		Nonce: nonce, Visibility: vis,
	}, nil
}

// bet_draft -------------------------------------------------------------------

type draftArgs struct {
	draftFields
}

func draftTool(_ context.Context, _ *mcp.CallToolRequest, in draftArgs) (*mcp.CallToolResult, draftView, error) {
	dIn, err := in.input()
	if err != nil {
		return errResult[draftView](err)
	}
	d, err := bets.NewDraft(dIn)
	if err != nil {
		return errResult[draftView](err)
	}
	return nil, renderDraft(d), nil
}

// bet_create ------------------------------------------------------------------

type createArgs struct {
	From    string `json:"from" jsonschema:"creator address (tx sender)"`
	Factory string `json:"factory" jsonschema:"BetEscrowFactory address"`
	draftFields
}

func createTool(_ context.Context, _ *mcp.CallToolRequest, in createArgs) (*mcp.CallToolResult, txEnvelope, error) {
	from, err := mustAddr(in.From, "from")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	factory, err := mustAddr(in.Factory, "factory")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	dIn, err := in.input()
	if err != nil {
		return errResult[txEnvelope](err)
	}
	d, err := bets.NewDraft(dIn)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	tx, err := txbuild.CreateBet(from, factory, d)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	env, err := envelope("create bet", d.TermsHash().Hex(), tx)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	return nil, env, nil
}

// bet_propose_link ------------------------------------------------------------

type proposeArgs struct {
	Proposer        string `json:"proposer" jsonschema:"your wallet address (the side you take)"`
	Side            string `json:"side" jsonschema:"the side YOU take: YES or NO"`
	Handle          string `json:"handle,omitempty" jsonschema:"optional display name / X handle"`
	Token           string `json:"token" jsonschema:"collateral token address (USDC)"`
	Stake           string `json:"stake" jsonschema:"your stake in base units (USDC 6dp)"`
	CounterStake    string `json:"counterStake,omitempty" jsonschema:"the taker's stake (default: match yours)"`
	Statement       string `json:"statement" jsonschema:"the bet statement"`
	Source          string `json:"source" jsonschema:"primary resolution source"`
	Fallback        string `json:"fallback,omitempty" jsonschema:"fallback resolution source"`
	EventTime       uint64 `json:"eventTime,omitempty" jsonschema:"event time (unix seconds)"`
	ClaimDeadline   uint64 `json:"claimDeadline,omitempty" jsonschema:"claim deadline (unix; default event+24h)"`
	ChallengeWindow uint64 `json:"challengeWindow,omitempty" jsonschema:"challenge window seconds (default 24h)"`
	Arbiter         string `json:"arbiter,omitempty" jsonschema:"optional arbiter address (enables disputes)"`
	Nonce           string `json:"nonce,omitempty" jsonschema:"disambiguation nonce (default 0)"`
	Public          bool   `json:"public,omitempty" jsonschema:"mark the resulting bet public"`
	Expiry          uint64 `json:"expiry,omitempty" jsonschema:"link expiry (unix seconds; 0 = never)"`
	BaseURL         string `json:"baseUrl,omitempty" jsonschema:"site base URL for the link (default from FRAY_BASE_URL)"`
}

type proposeView struct {
	ProposalID      string `json:"proposalId"`
	Link            string `json:"link"`
	Token           string `json:"token"`
	ProposerTakes   string `json:"proposerTakes"`
	AccepterTakes   string `json:"accepterTakes"`
	Proposer        string `json:"proposer"`
	Statement       string `json:"statement"`
	CollateralToken string `json:"collateralToken"`
	ProposerStake   string `json:"proposerStake"`
	CounterStake    string `json:"counterStake"`
	Visibility      string `json:"visibility"`
	Note            string `json:"note"`
}

func proposeTool(_ context.Context, _ *mcp.CallToolRequest, in proposeArgs) (*mcp.CallToolResult, proposeView, error) {
	o, err := in.open()
	if err != nil {
		return errResult[proposeView](err)
	}
	tok, err := o.Encode()
	if err != nil {
		return errResult[proposeView](err)
	}
	id, err := o.ID()
	if err != nil {
		return errResult[proposeView](err)
	}
	link, err := o.Link(resolveBaseURL(in.BaseURL))
	if err != nil {
		return errResult[proposeView](err)
	}
	return nil, proposeView{
		ProposalID:      id.Hex(),
		Link:            link,
		Token:           tok,
		ProposerTakes:   o.Side.String(),
		AccepterTakes:   o.Side.Other().String(),
		Proposer:        o.Proposer.Hex(),
		Statement:       o.Statement,
		CollateralToken: o.CollateralToken.Hex(),
		ProposerStake:   o.ProposerStake.String(),
		CounterStake:    o.ResolvedCounterStake().String(),
		Visibility:      o.Visibility.String(),
		Note: "Share this link. Anyone can accept by taking the " + o.Side.Other().String() +
			" side; the bet only commits once both sides fund the escrow on-chain.",
	}, nil
}

func (in proposeArgs) open() (challenge.Open, error) {
	proposer, err := mustAddr(in.Proposer, "proposer")
	if err != nil {
		return challenge.Open{}, err
	}
	token, err := mustAddr(in.Token, "token")
	if err != nil {
		return challenge.Open{}, err
	}
	sd, err := challenge.ParseSide(in.Side)
	if err != nil {
		return challenge.Open{}, err
	}
	stake, err := mustBig(in.Stake, "stake")
	if err != nil {
		return challenge.Open{}, err
	}
	nonce, err := bigOrZero(in.Nonce)
	if err != nil {
		return challenge.Open{}, err
	}
	arb, err := optAddr(in.Arbiter)
	if err != nil {
		return challenge.Open{}, err
	}
	o := challenge.Open{
		Proposer: proposer, Side: sd, Handle: in.Handle, CollateralToken: token,
		ProposerStake: stake, Statement: in.Statement, PrimarySource: in.Source,
		FallbackSource: in.Fallback, EventTime: in.EventTime, ClaimDeadline: in.ClaimDeadline,
		ChallengeWindow: in.ChallengeWindow, Arbiter: arb, Nonce: nonce,
		Visibility: core.VisibilityPrivate, Expiry: in.Expiry,
	}
	if in.Public {
		o.Visibility = core.VisibilityPublic
	}
	if strings.TrimSpace(in.CounterStake) != "" {
		cs, err := mustBig(in.CounterStake, "counterStake")
		if err != nil {
			return challenge.Open{}, err
		}
		o.CounterStake = cs
	}
	return o, o.Validate()
}

// bet_accept_link -------------------------------------------------------------

type acceptArgs struct {
	Token   string `json:"token,omitempty" jsonschema:"the challenge token (from propose_link)"`
	Link    string `json:"link,omitempty" jsonschema:"a full challenge URL (token read from ?p=)"`
	As      string `json:"as" jsonschema:"your wallet address (you take the open side)"`
	From    string `json:"from,omitempty" jsonschema:"tx sender for create (default: as)"`
	Factory string `json:"factory,omitempty" jsonschema:"BetEscrowFactory address; set to also emit the create tx"`
}

type acceptView struct {
	Draft    draftView   `json:"draft"`
	CreateTx *txEnvelope `json:"createTx,omitempty"`
}

func acceptTool(_ context.Context, _ *mcp.CallToolRequest, in acceptArgs) (*mcp.CallToolResult, acceptView, error) {
	tok, err := tokenFromArgs(in.Token, in.Link)
	if err != nil {
		return errResult[acceptView](err)
	}
	o, err := challenge.Decode(tok)
	if err != nil {
		return errResult[acceptView](err)
	}
	as, err := mustAddr(in.As, "as")
	if err != nil {
		return errResult[acceptView](err)
	}
	d, err := o.AcceptAt(as, uint64(time.Now().Unix()))
	if err != nil {
		return errResult[acceptView](err)
	}
	out := acceptView{Draft: renderDraft(d)}

	if strings.TrimSpace(in.Factory) != "" {
		fromStr := in.From
		if strings.TrimSpace(fromStr) == "" {
			fromStr = in.As
		}
		from, err := mustAddr(fromStr, "from")
		if err != nil {
			return errResult[acceptView](err)
		}
		factory, err := mustAddr(in.Factory, "factory")
		if err != nil {
			return errResult[acceptView](err)
		}
		tx, err := txbuild.CreateBet(from, factory, d)
		if err != nil {
			return errResult[acceptView](err)
		}
		env, err := envelope("accept challenge + create bet", d.TermsHash().Hex(), tx)
		if err != nil {
			return errResult[acceptView](err)
		}
		out.CreateTx = &env
	}
	return nil, out, nil
}

func tokenFromArgs(token, link string) (string, error) {
	if strings.TrimSpace(token) != "" {
		return strings.TrimSpace(token), nil
	}
	if strings.TrimSpace(link) == "" {
		return "", errNoToken
	}
	u, err := url.Parse(strings.TrimSpace(link))
	if err != nil {
		return "", err
	}
	p := u.Query().Get("p")
	if p == "" {
		return "", errNoToken
	}
	return p, nil
}

// --- escrow action tools (each returns an unsigned tx) ---

type escrowArgs struct {
	From   string `json:"from" jsonschema:"tx sender"`
	Escrow string `json:"escrow" jsonschema:"BetEscrow address"`
}

func escrowAction(action string, build func(from, escrow core.Address) (txEnvelope, error)) mcp.ToolHandlerFor[escrowArgs, txEnvelope] {
	return func(_ context.Context, _ *mcp.CallToolRequest, in escrowArgs) (*mcp.CallToolResult, txEnvelope, error) {
		from, err := mustAddr(in.From, "from")
		if err != nil {
			return errResult[txEnvelope](err)
		}
		escrow, err := mustAddr(in.Escrow, "escrow")
		if err != nil {
			return errResult[txEnvelope](err)
		}
		env, err := build(from, escrow)
		if err != nil {
			return errResult[txEnvelope](err)
		}
		return nil, env, nil
	}
}

// bet_claim -------------------------------------------------------------------

type claimArgs struct {
	From         string `json:"from" jsonschema:"tx sender"`
	Escrow       string `json:"escrow" jsonschema:"BetEscrow address"`
	Outcome      string `json:"outcome" jsonschema:"YES or NO"`
	EvidenceHash string `json:"evidenceHash,omitempty" jsonschema:"evidence hash (bytes32, 0x…); default zero"`
}

func claimTool(_ context.Context, _ *mcp.CallToolRequest, in claimArgs) (*mcp.CallToolResult, txEnvelope, error) {
	from, err := mustAddr(in.From, "from")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	escrow, err := mustAddr(in.Escrow, "escrow")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	o, err := parseOutcome(in.Outcome)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	ev := core.Hash32{}
	if strings.TrimSpace(in.EvidenceHash) != "" {
		ev, err = core.HexToHash32(in.EvidenceHash)
		if err != nil {
			return errResult[txEnvelope](err)
		}
	}
	tx, err := txbuild.Claim(from, escrow, o, ev)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	env, err := envelope("claim "+o.String(), "", tx)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	return nil, env, nil
}

// bet_approve -----------------------------------------------------------------

type approveArgs struct {
	From    string `json:"from" jsonschema:"approver address"`
	Token   string `json:"token" jsonschema:"token address"`
	Spender string `json:"spender" jsonschema:"spender (escrow) address"`
	Amount  string `json:"amount" jsonschema:"amount in base units"`
}

func approveTool(_ context.Context, _ *mcp.CallToolRequest, in approveArgs) (*mcp.CallToolResult, txEnvelope, error) {
	from, err := mustAddr(in.From, "from")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	token, err := mustAddr(in.Token, "token")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	spender, err := mustAddr(in.Spender, "spender")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	amt, err := mustBig(in.Amount, "amount")
	if err != nil {
		return errResult[txEnvelope](err)
	}
	tx, err := txbuild.Approve(from, token, spender, amt)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	env, err := envelope("approve", "", tx)
	if err != nil {
		return errResult[txEnvelope](err)
	}
	return nil, env, nil
}
