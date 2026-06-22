package txbuild

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/fraybet/cli/bets"
	"github.com/fraybet/cli/bindings/betescrow"
	"github.com/fraybet/cli/bindings/betescrowfactory"
	"github.com/fraybet/cli/core"
)

var (
	from    = core.MustHexToAddress("0x1111111111111111111111111111111111111111")
	target  = core.MustHexToAddress("0x2222222222222222222222222222222222222222")
	token   = core.MustHexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e")
	spender = core.MustHexToAddress("0x3333333333333333333333333333333333333333")
)

func mustHash(t *testing.T, s string) core.Hash32 {
	t.Helper()
	h, err := core.HexToHash32(s)
	if err != nil {
		t.Fatal(err)
	}
	return h
}

func sampleDraft(t *testing.T) bets.Draft {
	t.Helper()
	d, err := bets.NewDraft(bets.DraftInput{
		YesAgent: from, NoAgent: target, CollateralToken: token,
		YesStake: big.NewInt(500_000_000), NoStake: big.NewInt(500_000_000),
		Statement: "x", PrimarySource: "s", EventTime: 1782518400,
		Nonce: big.NewInt(1), Visibility: core.VisibilityPublic,
	})
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func TestCreateSelectorAndTarget(t *testing.T) {
	tx, err := CreateBet(from, target, sampleDraft(t))
	if err != nil {
		t.Fatal(err)
	}
	if tx.From != from || tx.To != target {
		t.Errorf("from/to wrong: %s -> %s", tx.From.Hex(), tx.To.Hex())
	}
	parsed, _ := betescrowfactory.BetEscrowFactoryMetaData.GetAbi()
	if !bytes.Equal(tx.Data[:4], parsed.Methods["create"].ID) {
		t.Errorf("create selector mismatch: %x", tx.Data[:4])
	}
	args, err := parsed.Methods["create"].Inputs.Unpack(tx.Data[4:])
	if err != nil {
		t.Fatalf("unpack: %v", err)
	}
	if len(args) != 1 {
		t.Fatalf("expected 1 arg (terms tuple), got %d", len(args))
	}
}

func TestClaimArgsRoundTrip(t *testing.T) {
	evidence := mustHash(t, "0x"+strings.Repeat("ab", 32))
	tx, err := Claim(from, target, bets.OutcomeYes, evidence)
	if err != nil {
		t.Fatal(err)
	}
	parsed, _ := betescrow.BetEscrowMetaData.GetAbi()
	if !bytes.Equal(tx.Data[:4], parsed.Methods["claim"].ID) {
		t.Fatalf("claim selector mismatch")
	}
	args, err := parsed.Methods["claim"].Inputs.Unpack(tx.Data[4:])
	if err != nil {
		t.Fatalf("unpack: %v", err)
	}
	if args[0].(uint8) != uint8(bets.OutcomeYes) {
		t.Errorf("outcome = %v, want %d", args[0], bets.OutcomeYes)
	}
	if got := args[1].([32]byte); got != [32]byte(evidence) {
		t.Errorf("evidence not preserved")
	}
}

func TestApproveTargetsToken(t *testing.T) {
	tx, err := Approve(from, token, spender, big.NewInt(1_000_000))
	if err != nil {
		t.Fatal(err)
	}
	if tx.To != token {
		t.Errorf("approve should target the token, got %s", tx.To.Hex())
	}
	if len(tx.Data) < 4 {
		t.Error("no calldata")
	}
}

func TestEscrowSelectors(t *testing.T) {
	parsed, _ := betescrow.BetEscrowMetaData.GetAbi()
	builders := map[string]func() ([]byte, error){
		"fund":          func() ([]byte, error) { tx, e := Fund(from, target); return tx.Data, e },
		"challenge":     func() ([]byte, error) { tx, e := Challenge(from, target); return tx.Data, e },
		"finalize":      func() ([]byte, error) { tx, e := Finalize(from, target); return tx.Data, e },
		"voidUnclaimed": func() ([]byte, error) { tx, e := VoidUnclaimed(from, target); return tx.Data, e },
	}
	for name, build := range builders {
		data, err := build()
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if !bytes.Equal(data[:4], parsed.Methods[name].ID) {
			t.Errorf("%s selector mismatch: %x", name, data[:4])
		}
	}
}
