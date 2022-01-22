package main

import (
	"fmt"
	"log"
	"github.com/xuguangtech/iBuilding"
	"github.com/ethereum/go-ethereum/ethclient"
)

var providerToken = "RRXXG62RXYR52G2FUY37CG3E9ZZPCYMT9M"
var web3HttpProvider = "https://bsc-testnet.web3api.com/v1/RRXXG62RXYR52G2FUY37CG3E9ZZPCYMT9M"
var mainNet = "https://mainnet.infura.io"
var localhost = "http://127.0.0.1:8545"

func main() {
	client, err := ethclient.Dial( mainNet )
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
