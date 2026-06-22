package core

import (
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"
	"testing"
)

// parityFile mirrors testdata/parity_vectors.json — the shared source of truth
// that both this Go test and the Solidity contracts/test/Parity.t.sol assert
// against. Keep the two readers in sync with the JSON shape.
type parityFile struct {
	Domain struct {
		Name              string `json:"name"`
		Version           string `json:"version"`
		ChainID           int64  `json:"chainId"`
		VerifyingContract string `json:"verifyingContract"`
	} `json:"domain"`
	BetTermsTypeHash string         `json:"betTermsTypeHash"`
	DomainSeparator  string         `json:"domainSeparator"`
	Vectors          []parityVector `json:"vectors"`
}

type parityVector struct {
	Name  string `json:"name"`
	Terms struct {
		YesAgent        string `json:"yesAgent"`
		NoAgent         string `json:"noAgent"`
		CollateralToken string `json:"collateralToken"`
		YesStake        string `json:"yesStake"`
		NoStake         string `json:"noStake"`
		Statement       string `json:"statement"`
		EventTime       uint64 `json:"eventTime"`
		ClaimDeadline   uint64 `json:"claimDeadline"`
		ChallengeWindow uint64 `json:"challengeWindow"`
		PrimarySource   string `json:"primarySource"`
		FallbackSource  string `json:"fallbackSource"`
		Arbiter         string `json:"arbiter"`
		Nonce           string `json:"nonce"`
	} `json:"terms"`
	TermsHash string `json:"termsHash"`
	Digest    string `json:"digest"`
}

func (v parityVector) toTerms(t *testing.T) BetTerms {
	t.Helper()
	return BetTerms{
		YesAgent:        MustHexToAddress(v.Terms.YesAgent),
		NoAgent:         MustHexToAddress(v.Terms.NoAgent),
		CollateralToken: MustHexToAddress(v.Terms.CollateralToken),
		YesStake:        mustBig(t, v.Terms.YesStake),
		NoStake:         mustBig(t, v.Terms.NoStake),
		Statement:       v.Terms.Statement,
		EventTime:       v.Terms.EventTime,
		ClaimDeadline:   v.Terms.ClaimDeadline,
		ChallengeWindow: v.Terms.ChallengeWindow,
		PrimarySource:   v.Terms.PrimarySource,
		FallbackSource:  v.Terms.FallbackSource,
		Arbiter:         MustHexToAddress(v.Terms.Arbiter),
		Nonce:           mustBig(t, v.Terms.Nonce),
	}
}

// TestParity is the Go side of Go<->Solidity hashing parity. It recomputes the
// type hash, domain separator, and per-vector terms hash + digest, and asserts
// they equal the frozen values in the shared JSON. contracts/test/Parity.t.sol
// asserts the same JSON from Solidity. Run via `make parity`.
//
// Set UPDATE_PARITY=1 to print computed values (used once to freeze new vectors).
func TestParity(t *testing.T) {
	pf := loadParity(t)
	update := os.Getenv("UPDATE_PARITY") == "1"

	domain := EIP712Domain{
		Name:              pf.Domain.Name,
		Version:           pf.Domain.Version,
		ChainID:           big.NewInt(pf.Domain.ChainID),
		VerifyingContract: MustHexToAddress(pf.Domain.VerifyingContract),
	}

	check(t, update, "betTermsTypeHash", BetTermsTypeHash.Hex(), pf.BetTermsTypeHash)
	check(t, update, "domainSeparator", domain.Separator().Hex(), pf.DomainSeparator)

	if len(pf.Vectors) == 0 {
		t.Fatal("no parity vectors found")
	}
	for _, v := range pf.Vectors {
		terms := v.toTerms(t)
		if err := terms.Validate(); err != nil && !update {
			t.Errorf("vector %q: terms invalid: %v", v.Name, err)
		}
		gotTerms := terms.TermsHash()
		gotDigest := Digest(domain, gotTerms)
		check(t, update, "vector "+v.Name+" termsHash", gotTerms.Hex(), v.TermsHash)
		check(t, update, "vector "+v.Name+" digest", gotDigest.Hex(), v.Digest)
	}
}

func check(t *testing.T, update bool, label, got, want string) {
	t.Helper()
	if update {
		t.Logf("UPDATE %-40s = %s", label, got)
		return
	}
	if want == "" {
		t.Errorf("%s: no frozen value in JSON (run with UPDATE_PARITY=1 to compute: got %s)", label, got)
		return
	}
	if got != want {
		t.Errorf("%s mismatch:\n got  %s\n want %s", label, got, want)
	}
}

func loadParity(t *testing.T) parityFile {
	t.Helper()
	path := filepath.Join("testdata", "parity_vectors.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	var pf parityFile
	if err := json.Unmarshal(raw, &pf); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	return pf
}

func mustBig(t *testing.T, s string) *big.Int {
	t.Helper()
	x, ok := new(big.Int).SetString(s, 10)
	if !ok {
		t.Fatalf("invalid integer %q", s)
	}
	return x
}
