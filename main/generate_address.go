package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"log"
)

func main() {
	privateKey, err := crypto.GenerateKey()

	if err != nil{
		log.Fatal(err)
	}

//	使用FromECDSA将其转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)

//	十六进制编码之后删除0x
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

//	公钥是从私钥派生的
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("不能访问的公钥类型")
	}

	//剥离0x和04
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

//	通过publicKey生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	hash := sha3.
}
