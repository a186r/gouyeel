/*
读取合约的字节码，有时需要读取已经部署在链上的合约的字节码
 */
package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil{
		log.Fatal(err)
	}

	contratAddress := common.HexToAddress("0x3678095a4e8007e87d46594af6b63dd77e3f2ddf")
	byteCode, err := client.CodeAt(context.Background(), contratAddress, nil)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(byteCode))
}
