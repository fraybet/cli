package core

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// Address is a 20-byte EVM address. We define our own rather than depend on
// go-ethereum here so the core package stays dependency-light (only x/crypto
// for keccak). The chain layer (internal/chain) brings in go-ethereum proper.
type Address [20]byte

// HexToAddress parses a hex address, with or without the "0x" prefix.
func HexToAddress(s string) (Address, error) {
	var a Address
	s = strings.TrimPrefix(strings.TrimSpace(s), "0x")
	if len(s) != 40 {
		return a, fmt.Errorf("core: invalid address length %d (want 40 hex chars)", len(s))
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return a, fmt.Errorf("core: invalid address hex: %w", err)
	}
	copy(a[:], b)
	return a, nil
}

// MustHexToAddress is HexToAddress that panics on error. For tests and constants.
func MustHexToAddress(s string) Address {
	a, err := HexToAddress(s)
	if err != nil {
		panic(err)
	}
	return a
}

// Hex returns the lowercase 0x-prefixed hex form. (Checksum casing is a display
// concern handled at the edges, not in the hashing core.)
func (a Address) Hex() string { return "0x" + hex.EncodeToString(a[:]) }

func (a Address) String() string { return a.Hex() }

// Hash32 is a 32-byte keccak output (struct hash, digest, etc.).
type Hash32 [32]byte

// Hex returns the 0x-prefixed hex form.
func (h Hash32) Hex() string { return "0x" + hex.EncodeToString(h[:]) }

func (h Hash32) String() string { return h.Hex() }

// HexToHash32 parses a 32-byte hex hash, with or without the "0x" prefix.
func HexToHash32(s string) (Hash32, error) {
	var h Hash32
	s = strings.TrimPrefix(strings.TrimSpace(s), "0x")
	if len(s) != 64 {
		return h, fmt.Errorf("core: invalid hash length %d (want 64 hex chars)", len(s))
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return h, fmt.Errorf("core: invalid hash hex: %w", err)
	}
	copy(h[:], b)
	return h, nil
}

// Visibility controls whether a bilateral bet is discoverable publicly. Public
// bets drive the website and can be cloned / spawn markets (see mvp.md §0.8).
// Visibility is an indexing/UI concern, not a fund-safety one, so it lives in
// events/metadata rather than gating any economic path.
type Visibility uint8

const (
	VisibilityPrivate Visibility = iota
	VisibilityPublic
)

func (v Visibility) String() string {
	switch v {
	case VisibilityPublic:
		return "public"
	case VisibilityPrivate:
		return "private"
	default:
		return fmt.Sprintf("Visibility(%d)", uint8(v))
	}
}
