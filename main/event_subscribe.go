package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
// 	websocket客户端
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil{
		log.Fatal(err)
	}

// 	创建筛选查询
	contractAddress := common.HexToAddress("0x8c5d6983373aaca3825db38d5dd5808af55f30b5")
	query := ethereum.FilterQuery{
		Addresses:[]common.Address{contractAddress},
	}

// 	接收事件的方式是go channel，创建一个名为logs的channel
	logs := make(chan types.Log)

// 	现在我们要做的就是通过客户端调用SubscribeFilterLogs来订阅，它接收查询选项和输出通道。这将返回包含unsubscribe和error的订阅结构
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil{
		log.Fatal(err)
	}

// 	最后我们就需要写一个select语句设置一个连续循环来读入新的日志事件或者错误
	for{
		select {
		case err := <- sub.Err():
			log.Fatal(err)
		case vLog := <- logs:
			fmt.Println(vLog)
		}
	}
}

