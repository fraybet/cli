package chain

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/fraybet/cli/core"
)

func TestLoadProfile(t *testing.T) {
	data := []byte(`{
	  "chainId": 84532,
	  "rpcUrl": "https://sepolia.base.org",
	  "contracts": {"BetEscrowFactory": "0x036CbD53842c5426634e7929541eC2318f3dCF7e"}
	}`)
	p, err := LoadProfile("base-sepolia", data)
	if err != nil {
		t.Fatalf("LoadProfile: %v", err)
	}
	if p.ChainID.Int64() != 84532 {
		t.Errorf("chainId = %s, want 84532", p.ChainID)
	}
	got, err := p.Contract("BetEscrowFactory")
	if err != nil {
		t.Fatalf("Contract: %v", err)
	}
	if got != core.MustHexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e") {
		t.Errorf("resolved address = %s", got.Hex())
	}
	if _, err := p.Contract("Nope"); err == nil {
		t.Error("expected error for unknown contract")
	}
}

func TestUnsignedTxJSON(t *testing.T) {
	tx := BuildCall(
		core.MustHexToAddress("0x1111111111111111111111111111111111111111"),
		core.MustHexToAddress("0x2222222222222222222222222222222222222222"),
		big.NewInt(0),
		[]byte{0xde, 0xad, 0xbe, 0xef},
	)
	tx.Gas = 21000
	tx.Nonce = 7
	tx.ChainID = big.NewInt(84532)

	b, err := json.Marshal(tx)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]string
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m["data"] != "0xdeadbeef" {
		t.Errorf("data = %q", m["data"])
	}
	if m["gas"] != "0x5208" { // 21000
		t.Errorf("gas = %q, want 0x5208", m["gas"])
	}
	if m["chainId"] != "0x14a34" { // 84532
		t.Errorf("chainId = %q, want 0x14a34", m["chainId"])
	}
}

// Anvil/Hardhat well-known account 0 — fixed key for a deterministic round-trip.
const (
	anvilKey0  = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	anvilAddr0 = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266" // lowercase
)

func TestSignRecoverRoundTrip(t *testing.T) {
	from, err := AddressFromPrivKey(anvilKey0)
	if err != nil {
		t.Fatalf("AddressFromPrivKey: %v", err)
	}
	if from.Hex() != anvilAddr0 {
		t.Fatalf("derived address = %s, want %s", from.Hex(), anvilAddr0)
	}

	digest := core.BetTerms{
		YesAgent: from, NoAgent: core.MustHexToAddress("0x2222222222222222222222222222222222222222"),
		CollateralToken: core.MustHexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e"),
		YesStake:        big.NewInt(1), NoStake: big.NewInt(1),
		Statement: "x", EventTime: 1, ClaimDeadline: 2, ChallengeWindow: 1,
		PrimarySource: "s", Arbiter: core.Address{}, Nonce: big.NewInt(1),
	}.TermsHash()

	sig, err := SignDigest(anvilKey0, digest)
	if err != nil {
		t.Fatalf("SignDigest: %v", err)
	}
	if len(sig) != 65 {
		t.Fatalf("signature length = %d, want 65", len(sig))
	}
	recovered, err := RecoverSigner(digest, sig)
	if err != nil {
		t.Fatalf("RecoverSigner: %v", err)
	}
	if recovered != from {
		t.Errorf("recovered %s, want %s", recovered.Hex(), from.Hex())
	}

	// V in {27,28} convention must also recover.
	sig27 := make([]byte, 65)
	copy(sig27, sig)
	sig27[64] += 27
	if r2, err := RecoverSigner(digest, sig27); err != nil || r2 != from {
		t.Errorf("V=27 recovery failed: r=%s err=%v", r2.Hex(), err)
	}
}
