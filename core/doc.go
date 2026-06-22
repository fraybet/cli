// Package core holds the shared types, schemas, EIP-712 hashing, and validation
// reused across the CLI, MCP server, web, indexer, and monitors.
//
// EIP-712 hashing is the foundation everything else trusts: a bet's identity is
// the EIP-712 struct hash of its terms, and the same hash must be reproducible
// byte-for-byte in Solidity. The Go implementation here and the Solidity library
// in contracts/src/custom/EIP712Terms.sol are kept in lockstep by the shared
// parity vectors in testdata/parity_vectors.json (see TestParity).
//
// We use EIP-712 typed-data hashing rather than hand-rolled canonical-JSON
// hashing on purpose: canonical JSON across two languages is a nondeterminism
// footgun (key ordering, number formatting, unicode normalization). EIP-712 has
// a precise, well-specified encoding that both go and solc implement natively.
package core
