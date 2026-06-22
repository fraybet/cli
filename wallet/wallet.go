// Package wallet is agent self-custody: each agent generates and holds its own
// secp256k1 key (non-custodial — the platform never holds it). It can derive its
// address and sign EIP-712 digests. Keys are stored as hex keyfiles under a
// directory; this is dev-grade (unencrypted) for local anvil runs — production
// would use an encrypted keystore or a remote signer / smart account.
package wallet

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
)

// Wallet is an agent's self-custodied key.
type Wallet struct {
	Name    string
	Address core.Address
	privHex string
}

// PrivHex returns the 0x-prefixed private key (handle with care; never log).
func (w *Wallet) PrivHex() string { return "0x" + w.privHex }

// Sign signs a 32-byte digest, returning the 65-byte [R||S||V] signature.
func (w *Wallet) Sign(digest core.Hash32) ([]byte, error) {
	return chain.SignDigest(w.privHex, digest)
}

// generate makes a fresh wallet with a name.
func generate(name string) (*Wallet, error) {
	key, err := ethcrypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("wallet: generate key: %w", err)
	}
	priv := hex.EncodeToString(ethcrypto.FromECDSA(key))
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

// Create generates a new wallet and persists it. Errors if the name exists.
func (s *Store) Create(name string) (*Wallet, error) {
	if name == "" || strings.ContainsAny(name, "/\\.") {
		return nil, fmt.Errorf("wallet: invalid name %q", name)
	}
	if _, err := os.Stat(s.path(name)); err == nil {
		return nil, fmt.Errorf("wallet: %q already exists", name)
	}
	w, err := generate(name)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(s.path(name), []byte(w.privHex), 0o600); err != nil {
		return nil, fmt.Errorf("wallet: save %q: %w", name, err)
	}
	return w, nil
}

// Load reads a wallet by name.
func (s *Store) Load(name string) (*Wallet, error) {
	data, err := os.ReadFile(s.path(name))
	if err != nil {
		return nil, fmt.Errorf("wallet: load %q: %w", name, err)
	}
	priv := strings.TrimSpace(strings.TrimPrefix(string(data), "0x"))
	addr, err := chain.AddressFromPrivKey(priv)
	if err != nil {
		return nil, fmt.Errorf("wallet: %q corrupt: %w", name, err)
	}
	return &Wallet{Name: name, Address: addr, privHex: priv}, nil
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
