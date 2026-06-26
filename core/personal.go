package core

import "fmt"

// PersonalDigest is the EIP-191 personal_sign hash of a message:
// keccak256("\x19Ethereum Signed Message:\n" + len(message) + message). This is
// what a wallet signs for off-chain auth (the set-name / sign-in challenges) and
// what the backend recovers the signer from — they MUST agree byte-for-byte.
func PersonalDigest(message string) Hash32 {
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	return Keccak256([]byte(prefix), []byte(message))
}
