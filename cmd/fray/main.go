// Command fray is the developer / operator / power-agent CLI for the
// protocol. Stage 0 ships the skeleton + a working `version`; later stages add
// the agent/wallet/bet/arbiter/market/order/position/resolution domains
// (design §17).
package main

import (
	"fmt"
	"os"
)

// version is overridden at build time via -ldflags "-X main.version=...".
var version = "0.0.0-dev"

const usage = `fray — agent-first, non-custodial binary betting CLI

usage: fray <domain> <command> [flags]

domains (added per execution-plan stage):
  agent       register/inspect agents and policies        (Stage 1)
  wallet      create/inspect self-custodied wallets        (Stage 1)
  bet         draft/create/fund/claim/challenge bets      (Stage 1)
  tx          sign + broadcast an unsigned tx (keystore)  (Stage 1)
  arbiter     list/quote/validate/verdict                 (Stage 1)
  market      draft/create/duplicate/resolvability        (Stage 3)
  order       place/cancel/fill (signed orders)           (Stage 4)
  position    list/pnl/redeem                             (Stage 4)
  resolution  propose/challenge/finalize (UMA)            (Stage 5)
  mcp         run the MCP server

meta:
  version     print version and exit
  help        print this message
`

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Print(usage)
		os.Exit(2)
	}
	switch args[0] {
	case "version", "--version", "-v":
		fmt.Printf("fray %s\n", version)
	case "help", "--help", "-h":
		fmt.Print(usage)
	case "bet":
		if err := runBet(args[1:], os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "fray:", err)
			os.Exit(1)
		}
	case "wallet":
		if err := runWallet(args[1:], os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "fray:", err)
			os.Exit(1)
		}
	case "tx":
		if err := runTx(args[1:], os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "fray:", err)
			os.Exit(1)
		}
	case "agent":
		if err := runAgent(args[1:], os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "fray:", err)
			os.Exit(1)
		}
	case "order":
		if err := runOrder(args[1:], os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "fray:", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "fray: unknown command %q (not yet implemented)\n\n", args[0])
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
}
