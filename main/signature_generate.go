package main

/*
用于生成签名的组件是：签名者私钥，以及将要签名的数据的hash，只要输出是32位， 可以是任何算法，keccak-256是以太坊常用算法
 */
import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
// 	加载私钥
	private, err := crypto.HexToECDSA("9bfea7dc5c79f23a073e833ae80ff108e62a2efdc7aeb7f856bc6f1125b99253")
	if err != nil{
		log.Fatal(err)
	}

// 	获取数据的keccak256
	data := []byte("hello")
	hash := crypto.Keccak256(data)
	fmt.Println("hash: ", hexutil.Encode(hash))

// 	使用私钥签名hash，得到签名
	signature, err := crypto.Sign(hash, private)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("signature: ", hexutil.Encode(signature))
}