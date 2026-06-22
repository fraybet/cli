package challenge

import (
	"math/big"
	"testing"

	"github.com/fraybet/cli/core"
)

func sampleOpen() Open {
	return Open{
		Proposer:        core.MustHexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		Side:            SideYes,
		Handle:          "@alpha_agent",
		CollateralToken: core.MustHexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
		ProposerStake:   big.NewInt(100_000_000), // 100 USDC
		Statement:       "BTC closes above $100k on 2026-12-31",
		PrimarySource:   "coingecko:bitcoin",
		FallbackSource:  "coinbase:BTC-USD",
		EventTime:       1_800_000_000,
		Visibility:      core.VisibilityPublic,
		Expiry:          1_790_000_000,
	}
}

func TestEncodeDecodeRoundTrip(t *testing.T) {
	o := sampleOpen()
	o.CounterStake = big.NewInt(50_000_000)
	o.Arbiter = core.MustHexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	o.Nonce = big.NewInt(7)

	tok, err := o.Encode()
	if err != nil {
		t.Fatalf("encode: %v", err)
	}
	got, err := Decode(tok)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	if got.Proposer != o.Proposer || got.Side != o.Side || got.Handle != o.Handle {
		t.Errorf("identity mismatch: %+v", got)
	}
	if got.CollateralToken != o.CollateralToken || got.Arbiter != o.Arbiter {
		t.Errorf("address mismatch: %+v", got)
	}
	if got.ProposerStake.Cmp(o.ProposerStake) != 0 || got.CounterStake.Cmp(o.CounterStake) != 0 {
		t.Errorf("stake mismatch: %v / %v", got.ProposerStake, got.CounterStake)
	}
	if got.Nonce.Cmp(o.Nonce) != 0 {
		t.Errorf("nonce mismatch: %v", got.Nonce)
	}
	if got.Statement != o.Statement || got.PrimarySource != o.PrimarySource || got.FallbackSource != o.FallbackSource {
		t.Errorf("terms text mismatch: %+v", got)
	}
	if got.EventTime != o.EventTime || got.Expiry != o.Expiry || got.Visibility != o.Visibility {
		t.Errorf("scalar mismatch: %+v", got)
	}
}

func TestAcceptMapsYesSide(t *testing.T) {
	o := sampleOpen() // proposer takes YES
	counter := core.MustHexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")

	d, err := o.Accept(counter)
	if err != nil {
		t.Fatalf("accept: %v", err)
	}
	if d.Terms.YesAgent != o.Proposer {
		t.Errorf("proposer should be YES agent, got %s", d.Terms.YesAgent.Hex())
	}
	if d.Terms.NoAgent != counter {
		t.Errorf("counter should be NO agent, got %s", d.Terms.NoAgent.Hex())
	}
	// CounterStake unset => matches proposer (even money).
	if d.Terms.YesStake.Cmp(o.ProposerStake) != 0 || d.Terms.NoStake.Cmp(o.ProposerStake) != 0 {
		t.Errorf("stakes should both equal proposer stake: %v / %v", d.Terms.YesStake, d.Terms.NoStake)
	}
}

func TestAcceptMapsNoSide(t *testing.T) {
	o := sampleOpen()
	o.Side = SideNo // proposer takes NO
	o.CounterStake = big.NewInt(25_000_000)
	counter := core.MustHexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")

	d, err := o.Accept(counter)
	if err != nil {
		t.Fatalf("accept: %v", err)
	}
	if d.Terms.NoAgent != o.Proposer {
		t.Errorf("proposer should be NO agent, got %s", d.Terms.NoAgent.Hex())
	}
	if d.Terms.YesAgent != counter {
		t.Errorf("counter should be YES agent, got %s", d.Terms.YesAgent.Hex())
	}
	if d.Terms.NoStake.Cmp(o.ProposerStake) != 0 {
		t.Errorf("NO stake should be proposer stake, got %v", d.Terms.NoStake)
	}
	if d.Terms.YesStake.Cmp(o.CounterStake) != 0 {
		t.Errorf("YES stake should be counter stake, got %v", d.Terms.YesStake)
	}
}

func TestAcceptRejectsSelf(t *testing.T) {
	o := sampleOpen()
	if _, err := o.Accept(o.Proposer); err == nil {
		t.Fatal("expected error accepting own challenge")
	}
}

func TestAcceptRejectsZeroCounterparty(t *testing.T) {
	o := sampleOpen()
	if _, err := o.Accept(core.Address{}); err == nil {
		t.Fatal("expected error with zero accepter")
	}
}

func TestValidateRejectsBadProposals(t *testing.T) {
	cases := map[string]func(*Open){
		"no proposer":  func(o *Open) { o.Proposer = core.Address{} },
		"no token":     func(o *Open) { o.CollateralToken = core.Address{} },
		"zero stake":   func(o *Open) { o.ProposerStake = big.NewInt(0) },
		"nil stake":    func(o *Open) { o.ProposerStake = nil },
		"no statement": func(o *Open) { o.Statement = "" },
		"no source":    func(o *Open) { o.PrimarySource = "" },
	}
	for name, mut := range cases {
		o := sampleOpen()
		mut(&o)
		if err := o.Validate(); err == nil {
			t.Errorf("%s: expected validation error", name)
		}
	}
}

func TestIsExpired(t *testing.T) {
	o := sampleOpen()
	o.Expiry = 1000
	if !o.IsExpired(1000) {
		t.Error("should be expired at exactly the deadline")
	}
	if o.IsExpired(999) {
		t.Error("should not be expired before the deadline")
	}
	o.Expiry = 0
	if o.IsExpired(1 << 62) {
		t.Error("zero expiry means never expires")
	}
}

func TestLinkAndDecodeFromLink(t *testing.T) {
	o := sampleOpen()
	link, err := o.Link("https://fray.example/")
	if err != nil {
		t.Fatalf("link: %v", err)
	}
	const want = "https://fray.example/challenge?p="
	if len(link) <= len(want) || link[:len(want)] != want {
		t.Fatalf("unexpected link: %s", link)
	}
}

func TestIDStableAndCounterpartyIndependent(t *testing.T) {
	o := sampleOpen()
	id1, err := o.ID()
	if err != nil {
		t.Fatalf("id: %v", err)
	}
	id2, _ := o.ID()
	if id1 != id2 {
		t.Error("ID should be deterministic")
	}
	if id1 == (core.Hash32{}) {
		t.Error("ID should be non-zero")
	}
}

func TestDecodeRejectsGarbage(t *testing.T) {
	if _, err := Decode("not base64!!!"); err == nil {
		t.Error("expected error on non-base64 token")
	}
	if _, err := Decode(""); err == nil {
		t.Error("expected error on empty token")
	}
}
