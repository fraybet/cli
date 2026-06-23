package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fraybet/cli/chain"
	"github.com/fraybet/cli/core"
	"github.com/fraybet/cli/txbuild"
	"github.com/fraybet/cli/wallet"
)

const walletUsage = `fray wallet <command> [flags]

commands:
  new       generate a new HD wallet (BIP-39 seed phrase; prints address + phrase)
  import    recreate a wallet from a seed phrase (--mnemonic)
  list      list wallet names ( * marks the default )
  use       set the default wallet — tx commands act as it unless --wallet is given
  send      send ERC-20 tokens (default USDC) from a wallet to an address
  address   print a wallet's address
  balance   print a wallet's native (ETH) balance  (--rpc required)
  export    print a wallet's PRIVATE KEY (for e.g. forge --private-key)

Keys live under --keystore (default ~/.agentbet/wallets), one file per wallet,
unencrypted (dev-grade). The platform is non-custodial: these keys are yours.
Commands act as the default wallet (set with `+"`fray wallet use <name>`"+`); pick a
different HD account within a wallet with --account <n>.
`

// defaultKeystore is ~/.agentbet/wallets, overridable with --keystore.
func defaultKeystore() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return ".agentbet/wallets"
	}
	return filepath.Join(home, ".agentbet", "wallets")
}

func runWallet(args []string, out io.Writer) error {
	if len(args) == 0 {
		fmt.Fprint(out, walletUsage)
		return fmt.Errorf("wallet: missing command")
	}
	cmd, rest := args[0], args[1:]
	switch cmd {
	case "new":
		return walletNew(rest, out)
	case "import":
		return walletImport(rest, out)
	case "list":
		return walletList(rest, out)
	case "use":
		return walletUse(rest, out)
	case "send":
		return walletSend(rest, out)
	case "address":
		return walletAddress(rest, out)
	case "balance":
		return walletBalance(rest, out)
	case "export":
		return walletExport(rest, out)
	case "help", "-h", "--help":
		fmt.Fprint(out, walletUsage)
		return nil
	default:
		fmt.Fprint(out, walletUsage)
		return fmt.Errorf("wallet: unknown command %q", cmd)
	}
}

func openStore(keystore string) (*wallet.Store, error) {
	if strings.TrimSpace(keystore) == "" {
		keystore = defaultKeystore()
	}
	return wallet.NewStore(keystore)
}

func walletNew(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("new", flag.ContinueOnError)
	fs.SetOutput(out)
	name := fs.String("name", "", "wallet name (e.g. deployer)")
	keystore := fs.String("keystore", "", "keystore dir (default ~/.agentbet/wallets)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*name) == "" {
		return fmt.Errorf("--name is required")
	}
	st, err := openStore(*keystore)
	if err != nil {
		return err
	}
	w, err := st.Create(*name)
	if err != nil {
		return err
	}
	ks := *keystore
	if ks == "" {
		ks = defaultKeystore()
	}
	fmt.Fprintf(out, "Created HD wallet %q\n", w.Name)
	fmt.Fprintf(out, "  address:  %s\n", w.Address.Hex())
	fmt.Fprintf(out, "  keystore: %s\n", filepath.Join(ks, w.Name+".key"))
	fmt.Fprintf(out, "\n  ⚠️  SEED PHRASE — write this down; it is the ONLY way to recover this wallet:\n")
	fmt.Fprintf(out, "  %s\n", w.Mnemonic)
	fmt.Fprintf(out, "\nFund the address above. Recover later with: fray wallet import --name %s --mnemonic \"…\"\n", w.Name)
	return nil
}

func walletImport(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("import", flag.ContinueOnError)
	fs.SetOutput(out)
	name := fs.String("name", "", "wallet name")
	mnemonic := fs.String("mnemonic", "", "BIP-39 seed phrase (quote it)")
	keystore := fs.String("keystore", "", "keystore dir (default ~/.agentbet/wallets)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*name) == "" {
		return fmt.Errorf("--name is required")
	}
	if strings.TrimSpace(*mnemonic) == "" {
		return fmt.Errorf("--mnemonic is required")
	}
	st, err := openStore(*keystore)
	if err != nil {
		return err
	}
	w, err := st.Import(*name, *mnemonic)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Imported wallet %q\n  address: %s\n", w.Name, w.Address.Hex())
	return nil
}

func walletList(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("list", flag.ContinueOnError)
	fs.SetOutput(out)
	keystore := fs.String("keystore", "", "keystore dir")
	if err := fs.Parse(args); err != nil {
		return err
	}
	st, err := openStore(*keystore)
	if err != nil {
		return err
	}
	names, err := st.List()
	if err != nil {
		return err
	}
	if len(names) == 0 {
		fmt.Fprintln(out, "(no wallets)")
		return nil
	}
	def, _ := st.Default()
	for _, n := range names {
		marker := "  "
		if n == def {
			marker = "* " // the default wallet
		}
		w, err := st.Load(n)
		if err != nil {
			fmt.Fprintf(out, "%s%s\t(error: %v)\n", marker, n, err)
			continue
		}
		fmt.Fprintf(out, "%s%s\t%s\n", marker, n, w.Address.Hex())
	}
	return nil
}

// walletUse sets the default wallet that commands act as when --wallet is omitted.
func walletUse(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("use", flag.ContinueOnError)
	fs.SetOutput(out)
	keystore := fs.String("keystore", "", "keystore dir")
	if err := fs.Parse(args); err != nil {
		return err
	}
	rest := fs.Args()
	if len(rest) != 1 {
		return fmt.Errorf("usage: fray wallet use <name>")
	}
	st, err := openStore(*keystore)
	if err != nil {
		return err
	}
	if err := st.SetDefault(rest[0]); err != nil {
		return err
	}
	w, err := st.Load(rest[0])
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "default wallet: %s (%s)\n", w.Name, w.Address.Hex())
	return nil
}

// walletSend signs + broadcasts an ERC-20 transfer from a wallet (default: the
// default wallet) to an address. Used to move funds between your own wallets.
func walletSend(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("send", flag.ContinueOnError)
	fs.SetOutput(out)
	to := fs.String("to", "", "recipient address")
	token := fs.String("token", defaultUSDC, "ERC-20 token address (default USDC)")
	amount := fs.String("amount", "", "amount in base units (USDC has 6 decimals)")
	sf := addSignerFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	toA, err := parseAddr(*to, "to")
	if err != nil {
		return err
	}
	tokenA, err := parseAddr(*token, "token")
	if err != nil {
		return err
	}
	amt, err := parseBig(*amount, "amount")
	if err != nil {
		return err
	}
	return sf.run(out, "send", "", func(from core.Address) (chain.UnsignedTx, error) {
		return txbuild.Transfer(from, tokenA, toA, amt)
	})
}

func walletAddress(args []string, out io.Writer) error {
	w, err := loadWalletFlags("address", args, out)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, w.Address.Hex())
	return nil
}

func walletExport(args []string, out io.Writer) error {
	w, err := loadWalletFlags("export", args, out)
	if err != nil {
		return err
	}
	// Private key to stdout; warning to stderr so piping `export` stays clean.
	fmt.Fprintln(os.Stderr, "WARNING: printing a PRIVATE KEY. Keep it secret; anyone with it controls the funds.")
	fmt.Fprintln(out, w.PrivHex())
	return nil
}

func walletBalance(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("balance", flag.ContinueOnError)
	fs.SetOutput(out)
	name := fs.String("name", "", "wallet name")
	keystore := fs.String("keystore", "", "keystore dir")
	rpc := fs.String("rpc", "", "RPC URL (required)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*rpc) == "" {
		return fmt.Errorf("--rpc is required")
	}
	st, err := openStore(*keystore)
	if err != nil {
		return err
	}
	w, err := st.Load(*name)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cl, err := chain.Dial(ctx, chain.Profile{Name: "rpc", RPCURL: *rpc})
	if err != nil {
		return err
	}
	defer cl.Close()
	bal, err := cl.BalanceAt(ctx, w.Address)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s\n", w.Address.Hex())
	fmt.Fprintf(out, "  %s wei\n", bal.String())
	fmt.Fprintf(out, "  %s ETH\n", formatEther(bal))
	return nil
}

// loadWalletFlags parses --name/--keystore and loads the wallet.
func loadWalletFlags(cmd string, args []string, out io.Writer) (*wallet.Wallet, error) {
	fs := flag.NewFlagSet(cmd, flag.ContinueOnError)
	fs.SetOutput(out)
	name := fs.String("name", "", "wallet name")
	keystore := fs.String("keystore", "", "keystore dir")
	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	if strings.TrimSpace(*name) == "" {
		return nil, fmt.Errorf("--name is required")
	}
	st, err := openStore(*keystore)
	if err != nil {
		return nil, err
	}
	return st.Load(*name)
}

// formatEther renders wei as a decimal ETH string (18 dp, trailing zeros trimmed).
func formatEther(wei *big.Int) string {
	if wei == nil {
		return "0"
	}
	neg := wei.Sign() < 0
	abs := new(big.Int).Abs(wei)
	wholeUnit := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	whole := new(big.Int).Quo(abs, wholeUnit)
	frac := new(big.Int).Rem(abs, wholeUnit)
	s := whole.String()
	if frac.Sign() != 0 {
		fracStr := frac.String()
		fracStr = strings.Repeat("0", 18-len(fracStr)) + fracStr // left-pad to 18 dp
		fracStr = strings.TrimRight(fracStr, "0")
		s += "." + fracStr
	}
	if neg {
		s = "-" + s
	}
	return s
}
