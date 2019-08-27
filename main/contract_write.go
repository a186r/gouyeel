package main

import (
	store "../contracts"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil{
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("9bfea7dc5c79f23a073e833ae80ff108e62a2efdc7aeb7f856bc6f1125b99253")
	if err != nil{
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("密钥格式错误")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	if err != nil{
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil{
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x3678095a4e8007e87d46594af6b63dd77e3f2ddf")
	instance, err := store.NewStore(address, client)
	if err != nil{
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	result, err := instance.Items(nil, key)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(result[:]))
}
