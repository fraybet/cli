package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/fraybet/cli/bets"
	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/txbuild"
)

const betUsage = `fray bet <command> [flags]

commands:
  draft         build + validate bet terms, print terms + termsHash
  clone         clone a public bet's terms for new agents
  propose-link  generate a shareable challenge link (post on X to find a taker)
  accept-link   accept a challenge link as the counterparty -> draft / create tx
  create        build the unsigned tx to create a bet via the factory
  approve     build the unsigned ERC-20 approve tx (before funding)
  fund        build the unsigned tx to fund your side
  claim       build the unsigned tx to claim YES/NO
  challenge   build the unsigned tx to challenge a claim
  finalize    build the unsigned tx to finalize an unchallenged claim
  void        build the unsigned tx to refund an unclaimed, expired bet

All tx commands return UNSIGNED transactions — sign and broadcast with your
own wallet. The platform never signs or holds keys.
`

// runBet dispatches `fray bet <cmd>`.
func runBet(args []string, out io.Writer) error {
	if len(args) == 0 {
		fmt.Fprint(out, betUsage)
		return fmt.Errorf("bet: missing command")
	}
	cmd, rest := args[0], args[1:]
	switch cmd {
	case "draft":
		return betDraft(rest, out)
	case "clone":
		return betClone(rest, out)
	case "propose-link":
		return betProposeLink(rest, out)
	case "accept-link":
		return betAcceptLink(rest, out)
	case "create":
		return betCreate(rest, out)
	case "approve":
		return betApprove(rest, out)
	case "fund", "challenge", "finalize", "void":
		return betSimpleAction(cmd, rest, out)
	case "claim":
		return betClaim(rest, out)
	case "help", "-h", "--help":
		fmt.Fprint(out, betUsage)
		return nil
	default:
		fmt.Fprint(out, betUsage)
		return fmt.Errorf("bet: unknown command %q", cmd)
	}
}

// draftFlags registers the bet-term flags and returns a builder for DraftInput.
func draftFlags(fs *flag.FlagSet) func() (bets.DraftInput, error) {
	yes := fs.String("yes", "", "YES agent address")
	no := fs.String("no", "", "NO agent address")
	token := fs.String("token", "", "collateral token address (USDC)")
	arbiter := fs.String("arbiter", "", "arbiter address (optional)")
	yesStake := fs.String("yes-stake", "", "YES stake in base units (USDC 6dp)")
	noStake := fs.String("no-stake", "", "NO stake in base units")
	statement := fs.String("statement", "", "the bet statement")
	source := fs.String("source", "", "primary resolution source")
	fallback := fs.String("fallback", "", "fallback resolution source")
	eventTime := fs.Uint64("event-time", 0, "event time (unix seconds)")
	claimDeadline := fs.Uint64("claim-deadline", 0, "claim deadline (unix; default event+24h)")
	challengeWindow := fs.Uint64("challenge-window", 0, "challenge window seconds (default 24h)")
	nonce := fs.String("nonce", "0", "disambiguation nonce")
	public := fs.Bool("public", false, "mark the bet public (discoverable, cloneable)")

	return func() (bets.DraftInput, error) {
		yesA, err := parseAddr(*yes, "yes")
		if err != nil {
			return bets.DraftInput{}, err
		}
		noA, err := parseAddr(*no, "no")
		if err != nil {
			return bets.DraftInput{}, err
		}
		tokenA, err := parseAddr(*token, "token")
		if err != nil {
			return bets.DraftInput{}, err
		}
		arbA, err := parseOptAddr(*arbiter)
		if err != nil {
			return bets.DraftInput{}, err
		}
		ys, err := parseBig(*yesStake, "yes-stake")
		if err != nil {
			return bets.DraftInput{}, err
		}
		ns, err := parseBig(*noStake, "no-stake")
		if err != nil {
			return bets.DraftInput{}, err
		}
		nonceB, err := parseBig(*nonce, "nonce")
		if err != nil {
			return bets.DraftInput{}, err
		}
		vis := core.VisibilityPrivate
		if *public {
			vis = core.VisibilityPublic
		}
		return bets.DraftInput{
			YesAgent: yesA, NoAgent: noA, CollateralToken: tokenA, Arbiter: arbA,
			YesStake: ys, NoStake: ns,
			Statement: *statement, PrimarySource: *source, FallbackSource: *fallback,
			EventTime: *eventTime, ClaimDeadline: *claimDeadline, ChallengeWindow: *challengeWindow,
			Nonce: nonceB, Visibility: vis,
		}, nil
	}
}

func betDraft(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("draft", flag.ContinueOnError)
	fs.SetOutput(out)
	build := draftFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	in, err := build()
	if err != nil {
		return err
	}
	d, err := bets.NewDraft(in)
	if err != nil {
		return err
	}
	return writeJSON(out, renderDraft(d))
}

func betCreate(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("create", flag.ContinueOnError)
	fs.SetOutput(out)
	from := fs.String("from", "", "creator address (tx sender)")
	factory := fs.String("factory", "", "BetEscrowFactory address")
	build := draftFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	fromA, err := parseAddr(*from, "from")
	if err != nil {
		return err
	}
	factoryA, err := parseAddr(*factory, "factory")
	if err != nil {
		return err
	}
	in, err := build()
	if err != nil {
		return err
	}
	d, err := bets.NewDraft(in)
	if err != nil {
		return err
	}
	tx, err := txbuild.CreateBet(fromA, factoryA, d)
	if err != nil {
		return err
	}
	return writeUnsignedTx(out, "create bet", tx, d.TermsHash().Hex())
}

func betClone(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("clone", flag.ContinueOnError)
	fs.SetOutput(out)
	src := fs.String("from-terms", "", "path to a terms JSON (from `bet draft`) to clone")
	yes := fs.String("yes", "", "new YES agent")
	no := fs.String("no", "", "new NO agent")
	nonce := fs.String("nonce", "", "new nonce (required)")
	public := fs.Bool("public", false, "mark the clone public")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *src == "" {
		return fmt.Errorf("--from-terms is required")
	}
	srcTerms, err := loadTerms(*src)
	if err != nil {
		return err
	}
	yesA, err := parseAddr(*yes, "yes")
	if err != nil {
		return err
	}
	noA, err := parseAddr(*no, "no")
	if err != nil {
		return err
	}
	nonceB, err := parseBig(*nonce, "nonce")
	if err != nil {
		return err
	}
	vis := core.VisibilityPrivate
	if *public {
		vis = core.VisibilityPublic
	}
	d, err := bets.Clone(srcTerms, yesA, noA, nonceB, vis)
	if err != nil {
		return err
	}
	return writeJSON(out, renderDraft(d))
}

func betApprove(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("approve", flag.ContinueOnError)
	fs.SetOutput(out)
	from := fs.String("from", "", "approver address")
	token := fs.String("token", "", "token address")
	spender := fs.String("spender", "", "spender (escrow) address")
	amount := fs.String("amount", "", "amount in base units")
	if err := fs.Parse(args); err != nil {
		return err
	}
	fromA, err := parseAddr(*from, "from")
	if err != nil {
		return err
	}
	tokenA, err := parseAddr(*token, "token")
	if err != nil {
		return err
	}
	spenderA, err := parseAddr(*spender, "spender")
	if err != nil {
		return err
	}
	amt, err := parseBig(*amount, "amount")
	if err != nil {
		return err
	}
	tx, err := txbuild.Approve(fromA, tokenA, spenderA, amt)
	if err != nil {
		return err
	}
	return writeUnsignedTx(out, "approve", tx, "")
}

func betSimpleAction(action string, args []string, out io.Writer) error {
	fs := flag.NewFlagSet(action, flag.ContinueOnError)
	fs.SetOutput(out)
	from := fs.String("from", "", "tx sender")
	escrow := fs.String("escrow", "", "BetEscrow address")
	if err := fs.Parse(args); err != nil {
		return err
	}
	fromA, err := parseAddr(*from, "from")
	if err != nil {
		return err
	}
	escrowA, err := parseAddr(*escrow, "escrow")
	if err != nil {
		return err
	}
	var tx chain.UnsignedTx
	switch action {
	case "fund":
		tx, err = txbuild.Fund(fromA, escrowA)
	case "challenge":
		tx, err = txbuild.Challenge(fromA, escrowA)
	case "finalize":
		tx, err = txbuild.Finalize(fromA, escrowA)
	case "void":
		tx, err = txbuild.VoidUnclaimed(fromA, escrowA)
	}
	if err != nil {
		return err
	}
	return writeUnsignedTx(out, action, tx, "")
}

func betClaim(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("claim", flag.ContinueOnError)
	fs.SetOutput(out)
	from := fs.String("from", "", "tx sender")
	escrow := fs.String("escrow", "", "BetEscrow address")
	outcome := fs.String("outcome", "", "YES or NO")
	evidence := fs.String("evidence-hash", "0x0000000000000000000000000000000000000000000000000000000000000000", "evidence hash (bytes32)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	fromA, err := parseAddr(*from, "from")
	if err != nil {
		return err
	}
	escrowA, err := parseAddr(*escrow, "escrow")
	if err != nil {
		return err
	}
	o, err := parseOutcome(*outcome)
	if err != nil {
		return err
	}
	ev, err := core.HexToHash32(*evidence)
	if err != nil {
		return err
	}
	tx, err := txbuild.Claim(fromA, escrowA, o, ev)
	if err != nil {
		return err
	}
	return writeUnsignedTx(out, "claim "+o.String(), tx, "")
}
