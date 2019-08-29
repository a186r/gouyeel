/*
模拟客户端，使用client.commit开采区块，适合测试使用
 */
package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func main() {
	// 先生成一个账户私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil{
		log.Fatal(err)
	}

// 	从bind包创建一个NewKeyedTransactor，并为其传递私钥
	auth := bind.NewKeyedTransactor(privateKey)
// 	创建一个创世账户，并为其分配余额
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10)

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address:{
			Balance:balance,
		},
	}

	blockGasLimit := uint64(8000000)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

// 	可以像平时一样使用这个客户端，我们构造一笔交易并广播
	fromAddress := auth.From
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil{
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil{
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x590aD63D60b57082B4eBc246C1Fa4423b0695f74")
	var data []byte
// 	构造交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

// 	签名交易
	chainId := big.NewInt(1337)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil{
		log.Fatal(err)
	}
// 	广播已签名的交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil{
		log.Fatal(err)
	}
	// 交易hash，广播前就有的，其实就是签名后的交易的hash值
	fmt.Println("tx sent: ", signedTx.Hash().Hex())

// 	手动挖矿
	client.Commit()

// 	读取交易回执
	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil{
		log.Fatal(err)
	}

	if receipt == nil{
		log.Fatal("交易回执为空，未广播成功？")
	}

	fmt.Println("status: ", receipt.Status)
}