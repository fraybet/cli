package wallet

import (
	"os"
	"path/filepath"
	"strings"
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

func TestCreateIsHDAndRecoverable(t *testing.T) {
	s, _ := NewStore(t.TempDir())
	w, err := s.Create("alice")
	if err != nil {
		t.Fatal(err)
	}
	if len(strings.Fields(w.Mnemonic)) != 12 {
		t.Fatalf("expected a 12-word seed phrase, got %q", w.Mnemonic)
	}
	// Re-importing the phrase elsewhere yields the same address.
	s2, _ := NewStore(t.TempDir())
	w2, err := s2.Import("alice2", w.Mnemonic)
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	if w2.Address != w.Address {
		t.Errorf("import mismatch: %s != %s", w2.Address.Hex(), w.Address.Hex())
	}
}

func TestKnownMnemonicDerivesAnvilAccount0(t *testing.T) {
	// The well-known hardhat/anvil dev mnemonic must derive account 0 to this
	// address — proving the BIP-44 path m/44'/60'/0'/0/0 is correct.
	const mnemonic = "test test test test test test test test test test test junk"
	const want = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	s, _ := NewStore(t.TempDir())
	w, err := s.Import("anvil", mnemonic)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.EqualFold(w.Address.Hex(), want) {
		t.Fatalf("derived %s, want %s", w.Address.Hex(), want)
	}
}

func TestLegacyHexKeyfileLoads(t *testing.T) {
	dir := t.TempDir()
	s, _ := NewStore(dir)
	// anvil account 0 raw key (legacy keyfile format)
	const raw = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	if err := os.WriteFile(filepath.Join(dir, "legacy.key"), []byte(raw), 0o600); err != nil {
		t.Fatal(err)
	}
	w, err := s.Load("legacy")
	if err != nil {
		t.Fatalf("load legacy: %v", err)
	}
	if !strings.EqualFold(w.Address.Hex(), "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266") {
		t.Errorf("legacy address %s", w.Address.Hex())
	}
	if w.Mnemonic != "" {
		t.Error("legacy wallet should have no mnemonic")
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
