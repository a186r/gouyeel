/*
构建原始交易
 */
package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil{
		log.Fatal(nil)
	}

	privateKey, err := crypto.HexToECDSA("9bfea7dc5c79f23a073e833ae80ff108e62a2efdc7aeb7f856bc6f1125b99253")
	if err != nil{
		log.Fatal(err)
	}

	public := privateKey.Public()
	publicKey, ok := public.(*ecdsa.PublicKey)
	if !ok{
		log.Fatal("密钥格式不对")
	}

	// 获取nonce
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	if err != nil{
		log.Fatal(err)
	}

	value := big.NewInt(10000000000000000) // 0.01ETH
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil{
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x063B1c8Dd824e050f59294D13616F57aE11Ae411")
	var data []byte
	// 构造交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil{
		log.Fatal(err)
	}

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil{
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
// 	RLP是以太坊用于序列化对象的特殊编码方法
	rawTxBytes := ts.GetRlp(0)
// 	将原始字节转换为十六进制串，这就是交易广播前的原始数据信息
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Println(rawTxHex)
}