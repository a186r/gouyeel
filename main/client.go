package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("已经连接到以太坊节点")

	_ = client
}