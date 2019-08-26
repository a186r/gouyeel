package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/karalabe/go-ethereum/crypto/sha3"
	"log"
	"math/big"
)

func main() {
	//创建客户端
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil{
		log.Fatal(err)
	}
//	准备好需要的私钥
	privateKey, err := crypto.HexToECDSA("9bfea7dc5c79f23a073e833ae80ff108e62a2efdc7aeb7f856bc6f1125b99253")
	if err != nil{
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok{
		log.Fatal("密钥格式错误")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil{
		log.Fatal(err)
	}

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil{
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x063B1c8Dd824e050f59294D13616F57aE11Ae411")
	tokenAddress := common.HexToAddress("0x8C5d6983373AaCa3825DB38d5dD5808Af55f30B5")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	//	只取前4个字节,8位
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))
	//发送代币的地址左填充到32字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))
	//确定一下发送代币的数量
	amount := new(big.Int)
	amount.SetString("2000000000000000000", 10) // 2个UBT
	// 发送代币的数量也左填充32字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
// 	在客户端查询GasLimit
// 事实证明这里的GasLimit很可能不够用, 这里的Gas是针对计算量给出的建议值，字段全填充之后算出来的值刚刚好
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		To:       &tokenAddress,
		GasPrice: gasPrice,
		Value:    value,
		Data:     data,
	})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(gasLimit)
// 	组装交易数据
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil{
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil{
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

// 0xa9059cbb
// 000000000000000000000000063b1c8dd824e050f59294d13616f57ae11ae411
// 0000000000000000000000000000000000000000000000001bc16d674ec80000