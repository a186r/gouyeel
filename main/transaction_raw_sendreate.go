/*
上一步是创建原始交易，现在要将这一笔交易广播出去
 */
package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
)

func main() {
// 	首先将原始交易十六进制码解码为字节格式
	rawTx := "f86a0a843b9aca0082520894063b1c8dd824e050f59294d13616f57ae11ae411872386f26fc100008029a0c79305f4f72ced78e3abd7008b3371ec32c717b8d41dfffbeb7d1a0bd5393f80a049990bb1225afe2fe3fd9f5f90b02d22fc3a866b36da9872f2fbb613bb8561db"
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil{
		log.Fatal(err)
	}

// 	创建一个新的transaction指针
// 调用DecodeBytes将原始事物和指针传递给以太坊事务类型
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

// 	使用以太坊客户端广播交易
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil{
		log.Fatal(err)
	}
	client.SendTransaction(context.Background(), tx)

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
