package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//使用正则表达式判断是不是一个地址
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Println("is valid:", re.MatchString("0x6219C7c1849ca5098Ea505186106B99D33E1eFEC"))
	fmt.Println("is valid:", re.MatchString("0x2D022555f3A916761141db09622d93F2459B1DaC11"))

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	//部署的智能合约地址
	address := common.HexToAddress("0xBBB765796B0F9B0AE57efA6041A291A28893101d")
	bytescode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytescode) > 0

	fmt.Println("is Contract:", isContract)

	//	用户地址
	address = common.HexToAddress("0x9E9643E5d04262D6386a3B73F8EE568e4D95d172")
	bytescode, err = client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract = len(bytescode) > 0
	fmt.Println("is Contract:", isContract)
}
