package main

import (
	"fmt"
	"math/big"

	"github.com/core-coin/go-core/v2/core"
	"github.com/core-coin/go-core/v2/params"
)

func main() {
	officialGocoreGenesis := core.DefaultGenesisBlock() // get go-core mainnet genesis block

	genesisHash := officialGocoreGenesis.ToBlock(nil).Hash() // get genesis block hash
	fmt.Println("Mainnet genesis block hash: ", genesisHash.Hex())
	if genesisHash != params.MainnetGenesisHash {
		fmt.Println("Genesis block hash not match")
		return
	}
	fmt.Println("Mainnet genesis allocation: ")
	for addr, account := range officialGocoreGenesis.Alloc {
		fmt.Printf("Address %v - %v XCB\n", addr,  account.Balance.Div(account.Balance, big.NewInt(1000000000000000000)))
	}
}