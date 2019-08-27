package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	//	连接到客户端
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	//	导入私钥
	privateKey, err := crypto.HexToECDSA("9bfea7dc5c79f23a073e833ae80ff108e62a2efdc7aeb7f856bc6f1125b99253")
	if err != nil {
		log.Fatal(err)
	}
	//	想要获取nonce, 需要调用一个方法, 该方法需要我们的address, 所以从私钥派生address出来
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("密钥格式不对")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//	然后获取我们需要的随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1 * 1E18)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress("0x063B1c8Dd824e050f59294D13616F57aE11Ae411")
	//	数据data
	var data []byte
	//	组装交易数据
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	//  下一步是使用发件人的私钥对事务进行签名。
	//  为此，我们调用SignTx方法，该方法接受一个未签名的事务和我们之前构造的私钥。
	//  SignTx方法需要EIP155签名者，这个也需要我们先从客户端拿到链ID。
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}