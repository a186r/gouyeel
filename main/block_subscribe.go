/**
我们将讨论如何订阅以便在区块被开采时获得事件
 */
package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
// 	首先我们需要一个支持websocket的以太坊服务提供者， 使用infura的websocket
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil{
		log.Fatal(err)
	}

	// 创建一个新的通道，用于接收最新的区块头
	headers := make(chan *types.Header)

	// 调用client的SubscribeNewHead方法，接收刚才创建的区块头通道，并返回一个订阅对象
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil{
		log.Fatal(err)
	}

// 	订阅将推送新的区块头事件到我们的通道，因此我们可以选择一个select语句来监听新消息，订阅对象还包括一个error通道，该通道将在失败时发送消息
	for{
		select {
		case err := <-sub.Err() :
			log.Fatal(err)
		case header := <-headers :
			fmt.Println(header.Hash().Hex())
			// 	要获得该区块的完整内容，我们可以将区块头的摘要传递给客户端的BlockByHash函数
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil{
				log.Fatal(err)
			}
			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}


}