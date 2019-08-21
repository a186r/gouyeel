package main

import(
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil{
		log.Fatal(err)
	}

	account := common.HexToAddress("0x0ac26115fEacba375dA2eF39648A87A30519dCB9")

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(balance)

	blockNumber := big.NewInt(0)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	fmt.Println(balanceAt)

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	ethValue := fBalance.Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}
