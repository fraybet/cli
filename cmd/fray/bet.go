package main

import (
	"flag"
	"fmt"
	"io"
	"strings"

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
  propose-open  create an OPEN bet on-chain (your side; the other side open)
  create        build the unsigned tx to create a bilateral bet via the factory
  approve     build the unsigned ERC-20 approve tx (before funding)
  fund        build the unsigned tx to fund your side
  accept      take the open side of an OPEN bet (escrows your stake -> Live)
  revoke      cancel your un-accepted OPEN bet and reclaim your stake
  claim       build the unsigned tx to claim YES/NO
  agree       co-sign an outcome (fast-settle; both sides agree -> instant payout)
  challenge   build the unsigned tx to challenge a claim
  finalize    build the unsigned tx to finalize an unchallenged claim
  void        build the unsigned tx to refund an unclaimed, expired bet

tx commands sign with your default wallet and broadcast to Base mainnet.
Override the signer with --wallet <name> / --account <n>, the chain with --rpc,
or print the UNSIGNED tx (sign it elsewhere) with --unsigned. Keys are yours,
held locally — the platform never signs or holds them.
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
	case "propose-open":
		return betProposeOpen(rest, out)
	case "approve":
		return betApprove(rest, out)
	case "fund", "challenge", "finalize", "void", "accept", "revoke":
		return betSimpleAction(cmd, rest, out)
	case "claim":
		return betClaim(rest, out)
	case "agree":
		return betAgree(rest, out)
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
	factory := fs.String("factory", defaultFactory, "BetEscrowFactory address")
	build := draftFlags(fs)
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
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
	return sf.run(out, "create bet", d.TermsHash().Hex(), func(from core.Address) (chain.UnsignedTx, error) {
		return txbuild.CreateBet(from, factoryA, d)
	})
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
	token := fs.String("token", defaultUSDC, "token address")
	spender := fs.String("spender", "", "spender address (e.g. an escrow or the registry)")
	amount := fs.String("amount", "", "amount in base units")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
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
	return sf.run(out, "approve", "", func(from core.Address) (chain.UnsignedTx, error) {
		return txbuild.Approve(from, tokenA, spenderA, amt)
	})
}

// betProposeOpen builds the unsigned tx to create an OPEN bet: the proposer takes
// one side, the other is left open for any taker to accept().
func betProposeOpen(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("propose-open", flag.ContinueOnError)
	fs.SetOutput(out)
	side := fs.String("side", "", "the side YOU take: YES or NO")
	token := fs.String("token", defaultUSDC, "collateral token address")
	stake := fs.String("stake", "", "your stake in base units (USDC 6dp)")
	counterStake := fs.String("counter-stake", "", "the taker's stake (default: match yours)")
	statement := fs.String("statement", "", "the bet statement")
	source := fs.String("source", "", "primary resolution source")
	fallback := fs.String("fallback", "", "fallback resolution source")
	eventTime := fs.Uint64("event-time", 0, "event time (unix seconds)")
	claimDeadline := fs.Uint64("claim-deadline", 0, "claim deadline (unix; default event+24h)")
	challengeWindow := fs.Uint64("challenge-window", 0, "challenge window seconds (default 24h)")
	arbiter := fs.String("arbiter", "", "arbiter address (optional; enables disputes)")
	nonce := fs.String("nonce", "0", "disambiguation nonce")
	public := fs.Bool("public", false, "mark the bet public (discoverable)")
	factory := fs.String("factory", defaultFactory, "BetEscrowFactory address")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	proposerYes, err := parseSideYN(*side)
	if err != nil {
		return err
	}
	tokenA, err := parseAddr(*token, "token")
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
	csB := stakeB
	if strings.TrimSpace(*counterStake) != "" {
		if csB, err = parseBig(*counterStake, "counter-stake"); err != nil {
			return err
		}
	}
	factoryA, err := parseAddr(*factory, "factory")
	if err != nil {
		return err
	}
	vis := core.VisibilityPrivate
	if *public {
		vis = core.VisibilityPublic
	}
	w, err := sf.resolve() // the proposer is the signing wallet
	if err != nil {
		return err
	}
	in := bets.DraftInput{
		CollateralToken: tokenA, Arbiter: arbA, Statement: *statement,
		PrimarySource: *source, FallbackSource: *fallback,
		EventTime: *eventTime, ClaimDeadline: *claimDeadline, ChallengeWindow: *challengeWindow,
		Nonce: nonceB, Visibility: vis,
	}
	if proposerYes {
		in.YesAgent, in.YesStake, in.NoStake = w.Address, stakeB, csB
	} else {
		in.NoAgent, in.NoStake, in.YesStake = w.Address, stakeB, csB
	}
	d, err := bets.NewOpenDraft(in, proposerYes)
	if err != nil {
		return err
	}
	tx, err := txbuild.CreateBet(w.Address, factoryA, d)
	if err != nil {
		return err
	}
	return sf.emit(out, "create open bet", d.TermsHash().Hex(), tx, w)
}

func betSimpleAction(action string, args []string, out io.Writer) error {
	fs := flag.NewFlagSet(action, flag.ContinueOnError)
	fs.SetOutput(out)
	escrow := fs.String("escrow", "", "BetEscrow address")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	escrowA, err := parseAddr(*escrow, "escrow")
	if err != nil {
		return err
	}
	return sf.run(out, action, "", func(from core.Address) (chain.UnsignedTx, error) {
		switch action {
		case "fund":
			return txbuild.Fund(from, escrowA)
		case "challenge":
			return txbuild.Challenge(from, escrowA)
		case "finalize":
			return txbuild.Finalize(from, escrowA)
		case "void":
			return txbuild.VoidUnclaimed(from, escrowA)
		case "accept":
			return txbuild.Accept(from, escrowA)
		case "revoke":
			return txbuild.Revoke(from, escrowA)
		}
		return chain.UnsignedTx{}, fmt.Errorf("unknown action %q", action)
	})
}

// betAgree co-signs an outcome (the fast-settle path). When both participants
// agree on the same outcome the bet pays out immediately, with no challenge
// window and no arbiter fees. Honest losers use this to settle cleanly.
func betAgree(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("agree", flag.ContinueOnError)
	fs.SetOutput(out)
	escrow := fs.String("escrow", "", "BetEscrow address")
	outcome := fs.String("outcome", "", "YES or NO")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
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
	return sf.run(out, "agree "+o.String(), "", func(from core.Address) (chain.UnsignedTx, error) {
		return txbuild.AgreeOutcome(from, escrowA, o)
	})
}

func betClaim(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("claim", flag.ContinueOnError)
	fs.SetOutput(out)
	escrow := fs.String("escrow", "", "BetEscrow address")
	outcome := fs.String("outcome", "", "YES or NO")
	evidence := fs.String("evidence-hash", "0x0000000000000000000000000000000000000000000000000000000000000000", "evidence hash (bytes32)")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
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
	return sf.run(out, "claim "+o.String(), "", func(from core.Address) (chain.UnsignedTx, error) {
		return txbuild.Claim(from, escrowA, o, ev)
	})
}
