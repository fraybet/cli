package core

import (
	"fmt"
	"strings"
)

// NormalizeHandle trims surrounding whitespace and a single leading '@' (display
// sugar), yielding the canonical handle that gets signed and stored. It does NOT
// lowercase — uniqueness is enforced case-insensitively elsewhere, but the cased
// handle is preserved for display.
func NormalizeHandle(h string) string {
	return strings.TrimPrefix(strings.TrimSpace(h), "@")
}

// HandleClaimMessage is the EIP-191 personal_sign payload for a handle→address
// claim: it registers a unique Twitter-style handle and points it at the signer's
// address. Signing it IS the mapping — the recovered signer authorizes exactly
// this handle resolving to this address (not merely a proof of identity). The
// handle leads because it is the unique key; the address is what it resolves to.
// A handle is just a username — not chain-scoped and not domain-scoped — so the
// claim carries only the handle, the address, and an expiry (which bounds replay).
// This string MUST be byte-identical on the CLI signer and the backend verifier
// (auth.VerifyHandleClaim).
func HandleClaimMessage(handle string, addr Address, expires int64) string {
	h := NormalizeHandle(handle)
	return strings.Join([]string{
		"Claim the @" + h + " handle.",
		"handle: " + h,
		"address: " + addr.Hex(),
		fmt.Sprintf("expires: %d", expires),
	}, "\n")
}
