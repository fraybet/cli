package main

// Live Base mainnet deployment, used as flag defaults so commands don't need
// the addresses passed every time. Source of truth: https://fray.bet/docs.
const (
	defaultFactory  = "0xBbb99f873f483cC30f2FC96d7F1AD85cDE92c50b" // BetEscrowFactory (open bets)
	defaultUSDC     = "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913" // Base USDC
	defaultRegistry = "0xFC5AEBb030f49A3cEA2eba04B7e734898bf78574" // AgentRegistry (long-lived)
	defaultRPCURL   = "https://mainnet.base.org"
)
