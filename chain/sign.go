package chain

import (
	"fmt"
	"strings"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/fraybet/cli/core"
)

// EIP-712 / secp256k1 signing helpers (go-ethereum crypto). Output matches
// Ethereum's 65-byte [R || S || V] layout with V in {0,1}; add 27 at the call
// site for contracts using the {27,28} convention. These exist for
// arbiters/resolvers signing attestations/verdicts (§9.8/§9.9); the platform
// never signs on a user's behalf.

// SignDigest signs a 32-byte digest with a hex secp256k1 private key.
func SignDigest(privHex string, digest core.Hash32) ([]byte, error) {
	key, err := ethcrypto.HexToECDSA(strings.TrimPrefix(strings.TrimSpace(privHex), "0x"))
	if err != nil {
		return nil, fmt.Errorf("chain: bad private key: %w", err)
	}
	sig, err := ethcrypto.Sign(digest[:], key)
	if err != nil {
		return nil, fmt.Errorf("chain: sign: %w", err)
	}
	return sig, nil
}

// RecoverSigner recovers the signing address from a digest and 65-byte signature.
// V may be {0,1} or {27,28}.
func RecoverSigner(digest core.Hash32, sig []byte) (core.Address, error) {
	if len(sig) != 65 {
		return core.Address{}, fmt.Errorf("chain: signature must be 65 bytes, got %d", len(sig))
	}
	norm := make([]byte, 65)
	copy(norm, sig)
	if norm[64] >= 27 {
		norm[64] -= 27
	}
	pub, err := ethcrypto.SigToPub(digest[:], norm)
	if err != nil {
		return core.Address{}, fmt.Errorf("chain: recover: %w", err)
	}
	return fromETH(ethcrypto.PubkeyToAddress(*pub)), nil
}

// AddressFromPrivKey derives the Ethereum address for a hex private key.
func AddressFromPrivKey(privHex string) (core.Address, error) {
	key, err := ethcrypto.HexToECDSA(strings.TrimPrefix(strings.TrimSpace(privHex), "0x"))
	if err != nil {
		return core.Address{}, fmt.Errorf("chain: bad private key: %w", err)
	}
	return fromETH(ethcrypto.PubkeyToAddress(key.PublicKey)), nil
}
