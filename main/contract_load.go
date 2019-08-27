/*
加载智能合约
接收一个智能合约地址，并返回可以与之交互的实例
 */
package main

import (
	store "../contracts"
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

	address := common.HexToAddress("0x3678095a4e8007e87d46594af6b63dd77e3f2ddf")
	instance, err := store.NewStore(address, client)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("合约载入成功")
	_ = instance
}