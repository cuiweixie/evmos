package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evmos/evmos/v15/precompiles/staking"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 确保你有正确的合约ABI
	parsedABI, err := staking.LoadABI()
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// 你需要调用的合约方法和参数
	from := common.HexToAddress("0x7cb61d4117ae31a12e393a1cfa3bac666481d02e")
	input, err := parsedABI.Pack("delegate", from,
		"evmosvaloper10jmp6sgh4cc6zt3e8gw05wavvejgr5pwlawghe", big.NewInt(100))
	if err != nil {
		log.Fatalf("Failed to pack data for method: %v", err)
	}

	contract := common.HexToAddress(staking.PrecompileAddress)
	gasPrice := big.NewInt(1000)
	//gasLimit := uint64(250000)
	msg := ethereum.CallMsg{
		From: from,
		To:   &contract,
		Data: input,
		//Gas:  gasLimit,
		GasPrice: gasPrice,
	}

	// 估计gas
	gas, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Fatalf("Failed to estimate gas: %v", err)
	}

	log.Printf("Gas: %v", gas)
}
