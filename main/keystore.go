package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func createKs(){
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

func importKs(){
	file := "./tmp/UTC--2019-08-22T02-21-14.056039586Z--69ba338190b6b26eeeb2ab74400440512747b4b5"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
//	读取file
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil{
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
	if err := os.Remove(file); err != nil{
		log.Fatal(err)
	}
}

func main() {
	//createKs()
	importKs()
}
