package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/evmos/evmos/v15/precompiles/staking"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA("E9B1D63E8ACD7FE676ACB43AFB390D4B0202DAB61ABEC9CF2A561E4BECB147DE")
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)     // in wei (1 eth)
	gasLimit := uint64(200000) // in units
	gasPrice := big.NewInt(1000)

	contract := common.HexToAddress(staking.PrecompileAddress)
	var data []byte

	// 确保你有正确的合约ABI
	parsedABI, err := staking.LoadABI()
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// 你需要调用的合约方法和参数
	from := common.HexToAddress("0x7cb61d4117ae31a12e393a1cfa3bac666481d02e")
	data, err = parsedABI.Pack("delegate", from,
		"evmosvaloper10jmp6sgh4cc6zt3e8gw05wavvejgr5pwlawghe", big.NewInt(100))
	if err != nil {
		log.Fatalf("Failed to pack data for method: %v", err)
	}

	tx := types.NewTransaction(nonce, contract, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0x77043e221eb4a76c9db6d59c59471be5a430aeb0455ef9d7c9e2e03b6e27e2a9

	var receipt *types.Receipt
	for {
		receipt, err = client.TransactionReceipt(context.Background(), signedTx.Hash())
		if err == nil {
			break
		}
	}
	fmt.Printf("receipt.GasUsed %v\n", receipt.GasUsed)
}
