package core

import (
	"math/big"
	"strconv"
)

// EIP-712 Order hashing — must stay byte-for-byte identical to the vendored
// CTF-Exchange Hashing mixin (OpenZeppelin EIP712 _hashTypedDataV4) and the
// ORDER_TYPEHASH in contracts/src/markets/exchange/libraries/OrderStructs.sol.
//
// Agents sign orders off-chain; the operator submits matched orders via
// exchange.matchOrders, which recovers the signer from exactly this digest.
// The domain separator uses the standard (non-salt) EIP712Domain, which is what
// OpenZeppelin's EIP712 emits, so EIP712Domain.Separator() above matches the
// exchange's _domainSeparatorV4() with name="Fray CTF Exchange", version="1".

// OrderSide is the side of an order.
type OrderSide uint8

const (
	OrderBuy  OrderSide = 0 // pay collateral, receive outcome tokens
	OrderSell OrderSide = 1 // give outcome tokens, receive collateral
)

// OrderSignatureType selects how the signature is validated on-chain. Only EOA
// is supported — Fray agent wallets are externally owned accounts.
type OrderSignatureType uint8

const OrderSigEOA OrderSignatureType = 0

// Order is the EIP-712-hashable order. Field order MUST match orderType and the
// Solidity Order struct exactly. The on-chain `signature` bytes are not part of
// the hash and so are not modeled here.
type Order struct {
	Salt          *big.Int
	Maker         Address // source of funds
	Signer        Address // signer of the order (== Maker for EOA orders)
	Taker         Address // counterparty; zero address = public order
	TokenID       *big.Int
	MakerAmount   *big.Int
	TakerAmount   *big.Int
	Expiration    *big.Int // unix seconds; 0 = no expiry
	Nonce         *big.Int // on-chain cancellation nonce
	FeeRateBps    *big.Int
	Side          OrderSide
	SignatureType OrderSignatureType
}

const orderType = "Order(" +
	"uint256 salt,address maker,address signer,address taker," +
	"uint256 tokenId,uint256 makerAmount,uint256 takerAmount," +
	"uint256 expiration,uint256 nonce,uint256 feeRateBps,uint8 side,uint8 signatureType)"

// OrderTypeHash is the EIP-712 type hash for Order.
var OrderTypeHash = keccak256([]byte(orderType))

// StructHash returns the EIP-712 hashStruct(Order).
func (o Order) StructHash() Hash32 {
	return keccak256(
		OrderTypeHash[:],
		encodeUint(o.Salt),
		encodeAddress(o.Maker),
		encodeAddress(o.Signer),
		encodeAddress(o.Taker),
		encodeUint(o.TokenID),
		encodeUint(o.MakerAmount),
		encodeUint(o.TakerAmount),
		encodeUint(o.Expiration),
		encodeUint(o.Nonce),
		encodeUint(o.FeeRateBps),
		encodeUint64(uint64(o.Side)),
		encodeUint64(uint64(o.SignatureType)),
	)
}

// FrayExchangeDomainName / Version are the EIP-712 domain identifiers baked into
// the deployed CTFExchange (see CTFExchange constructor).
const (
	FrayExchangeDomainName    = "Fray CTF Exchange"
	FrayExchangeDomainVersion = "1"
)

// ExchangeDomain returns the EIP-712 domain for the Fray CTF Exchange deployed at
// `exchange` on the chain identified by `chainID`.
func ExchangeDomain(chainID *big.Int, exchange Address) EIP712Domain {
	return EIP712Domain{
		Name:              FrayExchangeDomainName,
		Version:           FrayExchangeDomainVersion,
		ChainID:           chainID,
		VerifyingContract: exchange,
	}
}

// Digest returns the EIP-712 signing digest for this order under `domain` — the
// 32 bytes an agent wallet signs and the exchange recovers the signer from.
func (o Order) Digest(domain EIP712Domain) Hash32 {
	return Digest(domain, o.StructHash())
}

// CancelOrderMessage is the EIP-191 personal_sign payload to cancel a resting
// order off-chain. The maker signs it; the operator recovers the signer and
// removes the order only if it is the order's maker. Off-chain cancellation is
// instant and gas-free, so this proves intent without an on-chain tx. `expires`
// (unix seconds) bounds replay.
func CancelOrderMessage(digest Hash32, expires int64) string {
	return "Cancel order " + digest.Hex() + ".\nexpires: " + strconv.FormatInt(expires, 10)
}
