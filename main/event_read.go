package main

import (
	token "../contracts/erc20"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil{
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x8c5d6983373aaca3825db38d5dd5808af55f30b5")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6280227),
		ToBlock:   big.NewInt(6280257),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil{
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil{
		log.Fatal(err)
	}

	for _, vLog := range logs{
		fmt.Println("blockHash", vLog.BlockHash.Hex())
		fmt.Println("blockNumber", vLog.BlockNumber)
		fmt.Println("txHash", vLog.TxHash.Hex())

		event := struct {
			From common.Address
			To common.Address
			Tokens *big.Int
		}{}
		err := contractAbi.Unpack(&event, "Transfer", vLog.Data)
		if err != nil{
			log.Fatal(err)
		}

		fmt.Println("From: ", event.From)
		fmt.Println("To: ", event.To)
		fmt.Println("Tokens: ", event.Tokens)

		var topics [4]string
		for i := range vLog.Topics{
			topics[i] = vLog.Topics[i].Hex()
			fmt.Println("topics: ", topics[i])
		}
	}

	eventSignature := []byte("Transfer(address,address,uint)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("eventSignature: ", hash.Hex())
}