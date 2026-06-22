package wallet

import (
	"testing"

	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
)

func TestCreateLoadRoundTrip(t *testing.T) {
	s, err := NewStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	w, err := s.Create("alice")
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if (w.Address == core.Address{}) {
		t.Fatal("zero address")
	}
	loaded, err := s.Load("alice")
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if loaded.Address != w.Address {
		t.Errorf("address mismatch: %s != %s", loaded.Address.Hex(), w.Address.Hex())
	}
}

func TestCreateRejectsDuplicate(t *testing.T) {
	s, _ := NewStore(t.TempDir())
	s.Create("alice")
	if _, err := s.Create("alice"); err == nil {
		t.Error("expected error creating duplicate")
	}
}

func TestSignRecoversToAddress(t *testing.T) {
	s, _ := NewStore(t.TempDir())
	w, _ := s.Create("bob")

	var digest core.Hash32
	copy(digest[:], []byte("a 32-byte digest to sign......!!"))
	sig, err := w.Sign(digest)
	if err != nil {
		t.Fatal(err)
	}
	got, err := chain.RecoverSigner(digest, sig)
	if err != nil {
		t.Fatal(err)
	}
	if got != w.Address {
		t.Errorf("recovered %s, want %s", got.Hex(), w.Address.Hex())
	}
}

func TestList(t *testing.T) {
	s, _ := NewStore(t.TempDir())
	s.Create("alice")
	s.Create("bob")
	names, _ := s.List()
	if len(names) != 2 {
		t.Errorf("expected 2 wallets, got %v", names)
	}
}
