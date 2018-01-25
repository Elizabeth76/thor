package thor

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

// Constants of block chain.
const (
	BlockInterval             uint64 = 10 // time interval between two consecutive blocks.
	ClauseGas                 uint64 = params.TxGas * 2 / 3
	ClauseGasContractCreation uint64 = params.TxGasContractCreation * 2 / 3

	MinGasLimit          uint64 = 1000 * 1000
	InitialGasLimit      uint64 = 10 * 1000 * 1000 // InitialGasLimit gas limit value int genesis block.
	GasLimitBoundDivisor uint64 = 1024             // from ethereum

	MaxTxWorkDelay uint32 = 32 // (unit: block number) if tx delay exeeds this value, no energy can be exchanged.
)

// vars
var (
	WorkEnergyExchangeRate = big.NewInt(1e10) //TODO to be determined
)
