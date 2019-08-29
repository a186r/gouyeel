package main

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {

	// 私钥
	privateKey, err := crypto.HexToECDSA("9bfea7dc5c79f23a073e833ae80ff108e62a2efdc7aeb7f856bc6f1125b99253")
	if err != nil{
		log.Fatal(err)
	}

	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok{
		log.Fatal("密钥格式错误")
	}

// 	需要字节格式的公钥
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
// 	需要原始数据的hash
	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println("hash: ", hash.Hex())

// 	签名
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("signature: ", hexutil.Encode(signature))

// 	调用椭圆曲线恢复来检索签名者的公钥
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil{
		log.Fatal(err)
	}
// 比较预期的公钥与恢复的公钥是否一致
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println("签名验证：", matches)

// 	还有sigToPub可以达到同样的目的, 区别是它返回的ECDSA类型中的签名公钥
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err !=nil {
		log.Fatal(err)
	}
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
// 	比较公钥是否一致
	matches2 := bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println("签名验证方法2：", matches2)

// 	方法三，为了方便起见，以太坊还提供了一个VerifySignature函数，该函数接收原始数据的签名、哈希值和字节格式的公钥
// 	但是需要先移除最后一个字节，它是ECDSA恢复ID,不能包含它
	signatureNoRecoverID := signature[:len(signature)-1]
	verfied := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println("verfied:", verfied)
}