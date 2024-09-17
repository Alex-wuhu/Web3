package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	//header

	header, err := client.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String())

	blockNumber := big.NewInt(20738238) // block number from header

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Time())
	fmt.Println(block.GasUsed())
}
