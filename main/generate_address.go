package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/karalabe/go-ethereum/crypto/sha3"
)

func main() {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
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

	// 取最后40个字符(20个字节)并且加0x作为前缀
	// 以下是使用go-ethereum的crypto/sha3 Keccak256函数手动完成的方法。
	hash := sha3.NewKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))

}
