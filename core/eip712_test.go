package core

import (
	"math/big"
	"testing"
)

// TestEIP712KnownVector anchors the hashing primitives against the canonical
// EIP-712 "Mail" example from the spec, whose final signing digest is a
// well-known constant. If keccak256, the field encoders, or Digest composition
// were wrong, this would not reproduce the constant — so it independently
// validates the primitives that BetTerms hashing (and thus the parity vectors)
// rely on.
func TestEIP712KnownVector(t *testing.T) {
	const wantDigest = "0xbe609aee343fb3c4b28e1df9e632fca64fcfaede20f02e86244efddf30957bd2"

	domain := EIP712Domain{
		Name:              "Ether Mail",
		Version:           "1",
		ChainID:           big.NewInt(1),
		VerifyingContract: MustHexToAddress("0xCcCcccccCCCCcCCCCCCcCccccCcCCCcCcccccccC"),
	}

	personType := keccak256([]byte("Person(string name,address wallet)"))
	mailType := keccak256([]byte("Mail(Person from,Person to,string contents)Person(string name,address wallet)"))

	hashPerson := func(name string, wallet Address) Hash32 {
		h := keccak256(personType[:], encodeString(name), encodeAddress(wallet))
		return h
	}
	from := hashPerson("Cow", MustHexToAddress("0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826"))
	to := hashPerson("Bob", MustHexToAddress("0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB"))
	mailStruct := keccak256(mailType[:], from[:], to[:], encodeString("Hello, Bob!"))

	got := Digest(domain, mailStruct).Hex()
	if got != wantDigest {
		t.Fatalf("EIP-712 Mail digest mismatch:\n got  %s\n want %s\n(primitives are wrong — fix before trusting parity vectors)", got, wantDigest)
	}
}

func TestBetTermsValidate(t *testing.T) {
	good := sampleTerms()
	if err := good.Validate(); err != nil {
		t.Fatalf("valid terms rejected: %v", err)
	}

	tests := map[string]func(*BetTerms){
		"zero yesAgent":      func(b *BetTerms) { b.YesAgent = Address{} },
		"same agents":        func(b *BetTerms) { b.NoAgent = b.YesAgent },
		"zero stake":         func(b *BetTerms) { b.YesStake = big.NewInt(0) },
		"empty statement":    func(b *BetTerms) { b.Statement = "" },
		"deadline pre-event": func(b *BetTerms) { b.ClaimDeadline = b.EventTime - 1 },
		"zero challenge":     func(b *BetTerms) { b.ChallengeWindow = 0 },
		"nil nonce":          func(b *BetTerms) { b.Nonce = nil },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			terms := sampleTerms()
			mutate(&terms)
			if err := terms.Validate(); err == nil {
				t.Errorf("expected validation error for %q, got nil", name)
			}
		})
	}
}

// TestTermsHashDeterministic guards against accidental nondeterminism: the same
// terms must always hash to the same value.
func TestTermsHashDeterministic(t *testing.T) {
	a := sampleTerms().TermsHash()
	b := sampleTerms().TermsHash()
	if a != b {
		t.Fatalf("TermsHash is not deterministic: %s != %s", a, b)
	}
}

func sampleTerms() BetTerms {
	return BetTerms{
		YesAgent:        MustHexToAddress("0x1111111111111111111111111111111111111111"),
		NoAgent:         MustHexToAddress("0x2222222222222222222222222222222222222222"),
		CollateralToken: MustHexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e"),
		YesStake:        big.NewInt(500_000_000),
		NoStake:         big.NewInt(500_000_000),
		Statement:       "BTC-USD daily close on Coinbase will be above 120000 on 2026-07-01 UTC.",
		EventTime:       1782518400,
		ClaimDeadline:   1782604800,
		ChallengeWindow: 86400,
		PrimarySource:   "Coinbase BTC-USD daily candle",
		FallbackSource:  "Chainlink BTC/USD reference feed",
		Arbiter:         MustHexToAddress("0x3333333333333333333333333333333333333333"),
		Nonce:           big.NewInt(1),
	}
}
