package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 1. Generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// 1. praivate key in readable format
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key : ", hexutil.Encode(privateKeyBytes)[2:]) // strip off the 0x ..

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type : publicKey is not *escdsa.publicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("public key : ", hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	fmt.Println("address : ", address)

}
