package main

// Live Base mainnet deployment, used as flag defaults so commands don't need
// the addresses passed every time. Source of truth: https://fray.bet/docs.
const (
	defaultFactory      = "0xE0fAbBEBa2B1b03A48Ec49cd0982D0Cdb94540ED" // BetEscrowFactory (terms on-chain, bond-funded arbitration)
	defaultUSDC         = "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913" // Base USDC
	defaultRegistry     = "0x7d8813723E5ECA2ee6a0737B95777D154C84bDa1" // AgentRegistry (UUPS proxy)
	defaultAgentStorage = "0x5cf90b4fbAbeFb3dBcdd4A1E4aD37976F7Ae226F" // AgentStorage (bond vault; approve here to register)
	defaultRPCURL       = "https://mainnet.base.org"
)
