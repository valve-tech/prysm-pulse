package params

// UsePulseChainTestnetV4NetworkConfig uses the PulseChain beacon chain testnet network config.
func UsePulseChainTestnetV4NetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 16492700
	cfg.BootstrapNodes = []string{
		"enr:-Ly4QNnHfCcrfYaxr3kOmX5ZYZ63pwD1ryQ4quebAszaf_31EHKKrzJr3vKPOh2oh-Ot1yzIW9BXF3y1h7YT4aDkJOgDh2F0dG5ldHOIYAAAAAAAAACEZXRoMpCDALR_AAAJRv__________gmlkgnY0gmlwhEEVzL6Jc2VjcDI1NmsxoQIsEtaHkpFowrOeB-H15ELNL_gSqT3K3-WDZW1cMyh_74hzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA", // bootnode-001-hetzner-hel
		"enr:-Ly4QNMw1fgOmq4Njdjl_BaDCR2StP7t9AnY5FHSegGlYNW1FbxTMMa046goN5sf8Q6T6u64V3vTq1joTDpsEBKEOrIDh2F0dG5ldHOIAAAAAAAAYACEZXRoMpCDALR_AAAJRv__________gmlkgnY0gmlwhJ1agXaJc2VjcDI1NmsxoQKqm4KUrcvkTtOxxVjus8BsN9xhOwtGe0fpSj-OkPhdDIhzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA", // bootnode-002-hetzner-fsn
		"enr:-Ly4QBxXH32M0-TVeL1ZFHryMpzk1rG_4RzbWQRFSq8huYdMOnqZ33LpP6xm2xXGu_2ttntw-Si5kvoG8gZdj4PcPPIDh2F0dG5ldHOIAAAAAAMAAACEZXRoMpCDALR_AAAJRv__________gmlkgnY0gmlwhCUbOZ6Jc2VjcDI1NmsxoQKpUaMLmAmXDJa8tNxa7LHMSjVbGz5FJBeMZ3QTxRtriohzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA", // bootnode-001 maintained by www.g4mm4.io
		"enr:-Ly4QGDqnJR6yUu_qjoRGZ5Ret_reJn8pHhB-7UZEdhbQOMxM7LkQgwh4MXFglwnP6JOGaOMW1CaM6VXzEvkjmm1opsDh2F0dG5ldHOIADAAAAAAAACEZXRoMpCDALR_AAAJRv__________gmlkgnY0gmlwhNWFZISJc2VjcDI1NmsxoQN3M1XIu7gyViq2pu3xk_PLJQ5yor4Y3JvZrqOOcYOm-YhzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA", // bootnode-002 maintained by www.g4mm4.io
		"enr:-Ly4QNm-xwMQcMo8DZdVaGW_x4WFTHRaF7JfGDyuJ04eKfszdsvWZhCQD0H5E9JvaxWuABpldltnknqktKT0rB9J7qQDh2F0dG5ldHOIAAMAAAAAAACEZXRoMpCDALR_AAAJRv__________gmlkgnY0gmlwhFUKwbSJc2VjcDI1NmsxoQPEDXm3PE60XsUyjIHP7kCGlO_S7JX-pYcvPzCbyipL_4hzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA", // bootnode-003 maintained by www.g4mm4.io
	}
	OverrideBeaconNetworkConfig(cfg)
}

// PulseChainTestnetV4Config defines the config for the PulseChain beacon chain testnet.
func PulseChainTestnetV4Config() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.ConfigName = PulseChainTestnetV4Name
	cfg.PresetBase = "pulsechain"

	// preset overrides
	cfg.BaseRewardFactor = 64000
	cfg.EffectiveBalanceIncrement = 1 * 1e15
	cfg.MaxEffectiveBalance = 32 * 1e15

	// config overrides
	cfg.TerminalTotalDifficulty = "58750003716598352947541"
	cfg.MinGenesisActiveValidatorCount = 4096
	cfg.MinGenesisTime = 1674864000
	cfg.GenesisForkVersion = []byte{0x00, 0x00, 0x09, 0x43}
	cfg.GenesisDelay = 300
	cfg.AltairForkVersion = []byte{0x00, 0x00, 0x09, 0x44}
	cfg.AltairForkEpoch = 1
	cfg.BellatrixForkVersion = []byte{0x00, 0x00, 0x09, 0x45}
	cfg.BellatrixForkEpoch = 2
	cfg.CapellaForkVersion = []byte{0x00, 0x00, 0x09, 0x46}
	cfg.CapellaForkEpoch = 4200
	cfg.SecondsPerSlot = 10
	cfg.EjectionBalance = 16 * 1e15
	cfg.DepositChainID = 943
	cfg.DepositNetworkID = 943
	cfg.DepositContractAddress = "0x3693693693693693693693693693693693693693"

	cfg.InitializeForkSchedule()
	return cfg
}
