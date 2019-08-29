package main

import (
	token "../contracts/erc20"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil{
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0x8c5d6983373aaca3825db38d5dd5808af55f30b5")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil{
		log.Fatal(err)
	}

	address := common.HexToAddress("0x590aD63D60b57082B4eBc246C1Fa4423b0695f74")
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("wei: %s\n", balance)

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("symbol: %s\n", symbol)

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("decimals: %v\n", decimals)

// 	将余额转换为10进制格式
	fbal := new(big.Float)
	fbal.SetString(balance.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f", value)
}
