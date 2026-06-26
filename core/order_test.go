package core

import (
	"encoding/hex"
	"math/big"
	"testing"
)

// fixedOrder is a deterministic Order used as a shared Go<->Solidity parity
// vector. The verifyingContract is an arbitrary fixed address so the digest is
// reproducible independent of any live deployment.
func fixedOrder() (Order, EIP712Domain) {
	mk := func(b byte) Address {
		var a Address
		for i := range a {
			a[i] = b
		}
		return a
	}
	o := Order{
		Salt:          big.NewInt(42),
		Maker:         mk(0x11),
		Signer:        mk(0x11),
		Taker:         Address{}, // public order
		TokenID:       bigDec("97847902842828329890918613721842941496657289062292574132461597279393424722477"),
		MakerAmount:   big.NewInt(50_000_000),
		TakerAmount:   big.NewInt(100_000_000),
		Expiration:    big.NewInt(0),
		Nonce:         big.NewInt(0),
		FeeRateBps:    big.NewInt(0),
		Side:          OrderBuy,
		SignatureType: OrderSigEOA,
	}
	exchange := mustAddr("0x5615dEB798BB3E4dFa0139dFa1b3D433Cc23b72f")
	return o, ExchangeDomain(big.NewInt(8453), exchange)
}

func bigDec(s string) *big.Int {
	x, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("bad bigint: " + s)
	}
	return x
}

func mustAddr(s string) Address {
	b, err := hex.DecodeString(s[2:])
	if err != nil || len(b) != 20 {
		panic("bad addr: " + s)
	}
	var a Address
	copy(a[:], b)
	return a
}

func TestOrderHashGolden(t *testing.T) {
	o, domain := fixedOrder()

	if got := hex.EncodeToString(OrderTypeHash[:]); got != orderTypeHashGolden {
		t.Errorf("OrderTypeHash = %s, want %s", got, orderTypeHashGolden)
	}

	sep := domain.Separator()
	if got := hex.EncodeToString(sep[:]); got != exchangeDomainSepGolden {
		t.Errorf("domainSeparator = %s, want %s", got, exchangeDomainSepGolden)
	}

	sh := o.StructHash()
	if got := hex.EncodeToString(sh[:]); got != orderStructHashGolden {
		t.Errorf("StructHash = %s, want %s", got, orderStructHashGolden)
	}

	dg := o.Digest(domain)
	if got := hex.EncodeToString(dg[:]); got != orderDigestGolden {
		t.Errorf("Digest = %s, want %s", got, orderDigestGolden)
	}

	// Internal consistency: digest = keccak(0x1901 || sep || structHash).
	want := Digest(domain, sh)
	if want != dg {
		t.Errorf("digest not consistent with Digest()")
	}

	t.Logf("OrderTypeHash      = %s", hex.EncodeToString(OrderTypeHash[:]))
	t.Logf("exchangeDomainSep  = %s", hex.EncodeToString(sep[:]))
	t.Logf("orderStructHash    = %s", hex.EncodeToString(sh[:]))
	t.Logf("orderDigest        = %s", hex.EncodeToString(dg[:]))
}

// Goldens — filled from the first run, then asserted thereafter and mirrored in
// the Solidity parity test (smart-contracts/test/OrderParity.t.sol).
const (
	orderTypeHashGolden     = "a852566c4e14d00869b6db0220888a9090a13eccdaea03713ff0a3d27bf9767c"
	exchangeDomainSepGolden = "d4b80587b4c0c3c4c9b785853d5c2db5b05e5fe068b7763ae26cca8292aac2b1"
	orderStructHashGolden   = "af94d62f63e26b58df9631d8938b94ee7d4e1afd0cb887dbcdaf690f2df9a747"
	orderDigestGolden       = "e4ffebc92fe466601cc9f65303301e9f98e200df6026808c29dcdd92c6529a93"
)
