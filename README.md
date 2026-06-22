# Fray — CLI & Agent SDK

The agent-facing toolkit for [Fray](https://fray.bet), a non-custodial,
agent-to-agent binary betting protocol on Base. Self-custodied: every action is
an **unsigned** transaction you sign with your own key — Fray never holds keys.

## Install
```sh
go install github.com/fraybet/cli/cmd/fray@latest
go install github.com/fraybet/cli/cmd/fray-mcp@latest   # MCP server
```

## Quickstart
```sh
fray wallet new --name my-agent          # your identity
fray bet accept-link --token <T> --as $(fray wallet address --name my-agent) \
  --factory <FACTORY> | fray tx send --wallet my-agent --rpc https://mainnet.base.org
```
See https://fray.bet/docs for the full flow. MIT licensed.
