package exec

import (
	"context"
	"encoding/hex"
	"math/big"
	"os"
	"testing"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/fraybet/cli/bets"
	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
)

// TestAnvilCreateAndFund exercises the executor against a LIVE chain. It skips
// unless ANVIL_RPC is set. Run via scripts/exec-anvil.sh, which boots anvil,
// deploys, and passes ANVIL_RPC / FACTORY / USDC / FUNDER_KEY.
func TestAnvilCreateAndFund(t *testing.T) {
	rpc := os.Getenv("ANVIL_RPC")
	if rpc == "" {
		t.Skip("set ANVIL_RPC (+ FACTORY, USDC, FUNDER_KEY) to run the anvil integration test")
	}
	ctx := context.Background()
	factory := core.MustHexToAddress(os.Getenv("FACTORY"))
	usdc := core.MustHexToAddress(os.Getenv("USDC"))
	funder := os.Getenv("FUNDER_KEY")

	e, err := Dial(ctx, rpc)
	if err != nil {
		t.Fatal(err)
	}
	defer e.Close()

	aliceHex, alice := genAgent(t)
	bobHex, bob := genAgent(t)
	stake := big.NewInt(500_000_000)
	oneETH := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// Funder seeds both agents with gas + USDC.
	must(t, e.SendETH(ctx, funder, alice, oneETH))
	must(t, e.SendETH(ctx, funder, bob, oneETH))
	must(t, e.MintUSDC(ctx, funder, usdc, alice, stake))
	must(t, e.MintUSDC(ctx, funder, usdc, bob, stake))

	d, err := bets.NewDraft(bets.DraftInput{
		YesAgent: alice, NoAgent: bob, CollateralToken: usdc,
		YesStake: stake, NoStake: stake,
		Statement: "ETH above 5000 by 2027", PrimarySource: "Coinbase",
		EventTime: 1782518400, Nonce: big.NewInt(1), Visibility: core.VisibilityPublic,
	})
	if err != nil {
		t.Fatal(err)
	}

	escrow, err := e.CreateBet(ctx, aliceHex, factory, d)
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	t.Logf("escrow deployed at %s", escrow.Hex())

	must(t, e.Approve(ctx, aliceHex, usdc, escrow, stake))
	must(t, e.Fund(ctx, aliceHex, escrow))
	must(t, e.Approve(ctx, bobHex, usdc, escrow, stake))
	must(t, e.Fund(ctx, bobHex, escrow))

	bal, err := e.BalanceOfUSDC(ctx, usdc, escrow)
	if err != nil {
		t.Fatal(err)
	}
	if bal.Cmp(big.NewInt(1_000_000_000)) != 0 {
		t.Fatalf("escrow USDC = %s, want 1000000000 (both stakes funded)", bal)
	}
	t.Logf("escrow holds %s base-unit USDC — both sides funded on-chain", bal)
}

func genAgent(t *testing.T) (hexKey string, addr core.Address) {
	t.Helper()
	key, err := ethcrypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	h := hex.EncodeToString(ethcrypto.FromECDSA(key))
	a, err := chain.AddressFromPrivKey(h)
	if err != nil {
		t.Fatal(err)
	}
	return "0x" + h, a
}

func must(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
