package core

import (
	"math/big"

	"golang.org/x/crypto/sha3"
)

// EIP-712 hashing pipeline (must match contracts/src/custom/EIP712Terms.sol):
//
//   BetTerms ──encode fields──▶ encodeData ──keccak256──▶ structHash (TermsHash)
//                                                              │
//   domain ──encode──▶ keccak256 ──▶ domainSeparator           │
//                                          │                   │
//                                          └──┬────────────────┘
//                                             ▼
//        digest = keccak256( 0x19 0x01 ‖ domainSeparator ‖ structHash )
//
// structHash (a.k.a. termsHash) is the canonical bet identifier and needs no
// domain. The digest is what an arbiter actually signs (validation attestation
// §9.8, verdict §9.9), so it is bound to a domain (name/version/chain/contract).

// Keccak256 hashes the concatenation of the given byte slices (Ethereum's
// original Keccak-256). Exported for callers that need raw keccak (e.g. EIP-191
// message hashing in internal/auth).
func Keccak256(parts ...[]byte) Hash32 { return keccak256(parts...) }

// keccak256 hashes the concatenation of the given byte slices (Ethereum's
// original Keccak padding, NOT NIST SHA3 — they differ and are not compatible).
func keccak256(parts ...[]byte) Hash32 {
	h := sha3.NewLegacyKeccak256()
	for _, p := range parts {
		h.Write(p)
	}
	var out Hash32
	copy(out[:], h.Sum(nil))
	return out
}

// encodeUint left-pads a non-negative integer to a 32-byte big-endian word.
// Panics if the value exceeds 256 bits (a programming error, not user input).
func encodeUint(x *big.Int) []byte {
	if x == nil {
		x = big.NewInt(0)
	}
	if x.Sign() < 0 {
		panic("core: cannot EIP-712 encode a negative integer")
	}
	if x.BitLen() > 256 {
		panic("core: integer exceeds 256 bits")
	}
	b := make([]byte, 32)
	x.FillBytes(b)
	return b
}

func encodeUint64(x uint64) []byte { return encodeUint(new(big.Int).SetUint64(x)) }

// encodeAddress left-pads a 20-byte address into a 32-byte word.
func encodeAddress(a Address) []byte {
	b := make([]byte, 32)
	copy(b[12:], a[:])
	return b
}

// encodeString returns keccak256(bytes(s)) as a 32-byte word (EIP-712 hashes
// dynamic types).
func encodeString(s string) []byte {
	h := keccak256([]byte(s))
	return h[:]
}

// EIP712Domain is the standard domain separator input.
type EIP712Domain struct {
	Name              string
	Version           string
	ChainID           *big.Int
	VerifyingContract Address
}

var eip712DomainTypeHash = keccak256([]byte(
	"EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))

// Separator returns the EIP-712 domain separator.
func (d EIP712Domain) Separator() Hash32 {
	return keccak256(
		eip712DomainTypeHash[:],
		encodeString(d.Name),
		encodeString(d.Version),
		encodeUint(d.ChainID),
		encodeAddress(d.VerifyingContract),
	)
}

// BetTerms is the binary bilateral-bet terms object (design §9.3), reduced to
// the EIP-712-hashable fields. Field order here MUST match betTermsType and the
// Solidity struct exactly.
type BetTerms struct {
	YesAgent        Address
	NoAgent         Address
	CollateralToken Address
	YesStake        *big.Int
	NoStake         *big.Int
	Statement       string
	EventTime       uint64
	ClaimDeadline   uint64
	ChallengeWindow uint64
	PrimarySource   string
	FallbackSource  string
	Arbiter         Address
	Nonce           *big.Int
	Visibility      uint8 // bound into the hash so a signature commits to public/private
}

const betTermsType = "BetTerms(" +
	"address yesAgent,address noAgent,address collateralToken," +
	"uint256 yesStake,uint256 noStake,string statement," +
	"uint256 eventTime,uint256 claimDeadline,uint256 challengeWindow," +
	"string primarySource,string fallbackSource,address arbiter,uint256 nonce,uint8 visibility)"

// BetTermsTypeHash is the EIP-712 type hash for BetTerms.
var BetTermsTypeHash = keccak256([]byte(betTermsType))

// StructHash returns the EIP-712 hashStruct(BetTerms) — the canonical TermsHash.
func (t BetTerms) StructHash() Hash32 {
	return keccak256(
		BetTermsTypeHash[:],
		encodeAddress(t.YesAgent),
		encodeAddress(t.NoAgent),
		encodeAddress(t.CollateralToken),
		encodeUint(t.YesStake),
		encodeUint(t.NoStake),
		encodeString(t.Statement),
		encodeUint64(t.EventTime),
		encodeUint64(t.ClaimDeadline),
		encodeUint64(t.ChallengeWindow),
		encodeString(t.PrimarySource),
		encodeString(t.FallbackSource),
		encodeAddress(t.Arbiter),
		encodeUint(t.Nonce),
		encodeUint64(uint64(t.Visibility)),
	)
}

// TermsHash is the canonical bet identifier (alias for StructHash).
func (t BetTerms) TermsHash() Hash32 { return t.StructHash() }

// Digest returns the EIP-712 signing digest for the given domain and struct hash.
func Digest(domain EIP712Domain, structHash Hash32) Hash32 {
	sep := domain.Separator()
	return keccak256([]byte{0x19, 0x01}, sep[:], structHash[:])
}
