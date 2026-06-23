// Package wallet is agent self-custody: each agent generates and holds its own
// key (non-custodial — the platform never holds it). New wallets are HD wallets:
// a BIP-39 mnemonic (seed phrase) from which keys are derived at the standard
// Ethereum path m/44'/60'/0'/0/N, so a wallet is recoverable from its phrase.
// The mnemonic is stored as the keyfile; legacy raw-hex keyfiles still load.
// This is dev-grade (unencrypted) for local/CLI use — production would use an
// encrypted keystore or a remote signer / smart account.
package wallet

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"

	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
)

// Wallet is an agent's self-custodied key (derived from an HD seed phrase, or a
// legacy raw key). Mnemonic is set only for HD wallets. Account is the HD
// derivation index (m/44'/60'/0'/0/Account); 0 is the default account.
type Wallet struct {
	Name     string
	Address  core.Address
	Mnemonic string // empty for legacy raw-hex wallets
	Account  uint32
	privHex  string
}

// PrivHex returns the 0x-prefixed private key (handle with care; never log).
func (w *Wallet) PrivHex() string { return "0x" + w.privHex }

// Sign signs a 32-byte digest, returning the 65-byte [R||S||V] signature.
func (w *Wallet) Sign(digest core.Hash32) ([]byte, error) {
	return chain.SignDigest(w.privHex, digest)
}

// deriveKey derives the private key (hex) at m/44'/60'/0'/0/index from a mnemonic.
func deriveKey(mnemonic string, index uint32) (string, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return "", fmt.Errorf("wallet: invalid mnemonic")
	}
	seed := bip39.NewSeed(mnemonic, "") // no passphrase
	master, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", fmt.Errorf("wallet: master key: %w", err)
	}
	const h = bip32.FirstHardenedChild // 0x80000000
	// m / 44' / 60' / 0' / 0 / index
	path := []uint32{h + 44, h + 60, h + 0, 0, index}
	k := master
	for _, p := range path {
		if k, err = k.NewChildKey(p); err != nil {
			return "", fmt.Errorf("wallet: derive: %w", err)
		}
	}
	return fmt.Sprintf("%x", k.Key), nil
}

// fromMnemonic builds an HD wallet at the given account index from a seed phrase.
func fromMnemonic(name, mnemonic string, account uint32) (*Wallet, error) {
	priv, err := deriveKey(mnemonic, account)
	if err != nil {
		return nil, err
	}
	addr, err := chain.AddressFromPrivKey(priv)
	if err != nil {
		return nil, err
	}
	return &Wallet{Name: name, Address: addr, Mnemonic: mnemonic, Account: account, privHex: priv}, nil
}

// fromHex builds a wallet from a raw private key (legacy keyfiles).
func fromHex(name, privHex string) (*Wallet, error) {
	priv := strings.TrimSpace(strings.TrimPrefix(privHex, "0x"))
	addr, err := chain.AddressFromPrivKey(priv)
	if err != nil {
		return nil, err
	}
	return &Wallet{Name: name, Address: addr, privHex: priv}, nil
}

// Store is a directory-backed keystore.
type Store struct {
	dir string
}

// NewStore opens (creating if needed) a keystore directory.
func NewStore(dir string) (*Store, error) {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, fmt.Errorf("wallet: mkdir %s: %w", dir, err)
	}
	return &Store{dir: dir}, nil
}

func (s *Store) path(name string) string { return filepath.Join(s.dir, name+".key") }

func validName(name string) error {
	if name == "" || strings.ContainsAny(name, "/\\.") {
		return fmt.Errorf("wallet: invalid name %q", name)
	}
	return nil
}

// Create generates a new HD wallet (fresh 12-word mnemonic) and persists the
// mnemonic. It returns the wallet; w.Mnemonic is the seed phrase to back up.
func (s *Store) Create(name string) (*Wallet, error) {
	if err := validName(name); err != nil {
		return nil, err
	}
	if _, err := os.Stat(s.path(name)); err == nil {
		return nil, fmt.Errorf("wallet: %q already exists", name)
	}
	entropy, err := bip39.NewEntropy(128) // 12 words
	if err != nil {
		return nil, fmt.Errorf("wallet: entropy: %w", err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, fmt.Errorf("wallet: mnemonic: %w", err)
	}
	return s.save(name, mnemonic)
}

// Import recreates a wallet from an existing seed phrase.
func (s *Store) Import(name, mnemonic string) (*Wallet, error) {
	if err := validName(name); err != nil {
		return nil, err
	}
	mnemonic = strings.Join(strings.Fields(mnemonic), " ")
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, fmt.Errorf("wallet: invalid seed phrase")
	}
	if _, err := os.Stat(s.path(name)); err == nil {
		return nil, fmt.Errorf("wallet: %q already exists", name)
	}
	return s.save(name, mnemonic)
}

func (s *Store) save(name, mnemonic string) (*Wallet, error) {
	w, err := fromMnemonic(name, mnemonic, 0)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(s.path(name), []byte(mnemonic+"\n"), 0o600); err != nil {
		return nil, fmt.Errorf("wallet: save %q: %w", name, err)
	}
	// First wallet created becomes the default.
	if def, _ := s.Default(); def == "" {
		_ = s.SetDefault(name)
	}
	return w, nil
}

// Load reads a wallet by name at HD account 0. Legacy raw-hex wallets load too.
func (s *Store) Load(name string) (*Wallet, error) { return s.LoadAccount(name, 0) }

// LoadAccount reads a wallet by name, deriving HD account `account`
// (m/44'/60'/0'/0/account). Legacy raw-hex wallets ignore account.
func (s *Store) LoadAccount(name string, account uint32) (*Wallet, error) {
	data, err := os.ReadFile(s.path(name))
	if err != nil {
		return nil, fmt.Errorf("wallet: load %q: %w", name, err)
	}
	content := strings.TrimSpace(string(data))
	if bip39.IsMnemonicValid(content) {
		return fromMnemonic(name, content, account)
	}
	w, err := fromHex(name, content)
	if err != nil {
		return nil, fmt.Errorf("wallet: %q corrupt: %w", name, err)
	}
	return w, nil
}

// --- default wallet ---

func (s *Store) defaultPath() string { return filepath.Join(s.dir, ".default") }

// SetDefault records `name` as the default wallet for commands that don't pass
// --wallet explicitly.
func (s *Store) SetDefault(name string) error {
	if _, err := os.Stat(s.path(name)); err != nil {
		return fmt.Errorf("wallet: %q not found", name)
	}
	return os.WriteFile(s.defaultPath(), []byte(name+"\n"), 0o600)
}

// Default returns the default wallet name ("" if none is set).
func (s *Store) Default() (string, error) {
	data, err := os.ReadFile(s.defaultPath())
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// List returns the names of stored wallets.
func (s *Store) List() ([]string, error) {
	entries, err := os.ReadDir(s.dir)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, e := range entries {
		if n := strings.TrimSuffix(e.Name(), ".key"); n != e.Name() {
			names = append(names, n)
		}
	}
	return names, nil
}
