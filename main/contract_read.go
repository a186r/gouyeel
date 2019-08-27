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

	version, err := instance.Version(nil)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(version)
}