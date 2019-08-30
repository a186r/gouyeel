package main

import (
	"context"
	"fmt"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/karalabe/go-ethereum/common"
	"log"
	"math/big"
)

func main() {
	//先连接到一个客户端
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil{
		log.Fatal(err)
	}

//	查询区块高度
	block, err := client.BlockByNumber(context.Background(), big.NewInt(32))
	if err != nil{
		log.Fatal(err)
	}

//	遍历区块中的交易
	for _, tx := range block.Transactions(){
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())

		chainID, err := client.NetworkID(context.Background())
		if err != nil{
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err != nil{
			fmt.Println(msg.From().Hex())
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil{
			log.Fatal(err)
		}

		fmt.Println(receipt.Status)

	}

	blockHash := common.HexToHash("0x7ee4c73322f078c6fa599b7dc5ede740a434ffd648d56d202b338aecc399394a")
	count, err := client.TransactionCount(context.Background(), common2.Hash(blockHash))
	if err != nil{
		log.Fatal(err)
	}

	//遍历区块中的交易hash
	for idx := uint(0); idx < count; idx++{
		tx, err := client.TransactionInBlock(context.Background(), common2.Hash(blockHash), idx)
		if err != nil{
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex())
	}

	txHash := common.HexToHash("0xc58efd280403aa8b8486eb28d8203dbe6f6e26b2514051829ea04e41634222fa")
	tx, isPending, err := client.TransactionByHash(context.Background(), common2.Hash(txHash))
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending)

}
