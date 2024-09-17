package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("hello")
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal((err))
	}
	fmt.Println("we have a connection")
	_ = client
	account := common.HexToAddress("0x388C818CA8B9251b393131C08a736A67ccB19297")
	//blocknumer := big.NewInt(5532993)
	fbalance := new(big.Float)
	fmt.Println(client.BlockNumber(context.Background())) // most recent block
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	fbalance.SetString(balance.String())
	ethvalue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18))) // wei/10^18

	fmt.Println("e value: ", ethvalue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pendingBalance value: ", pendingBalance)

}
