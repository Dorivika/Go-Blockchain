package main

import (
	"fmt"
	"strconv"

	"github.com/Dorivika/Go-Blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First BLock")
	chain.AddBlock("second BLock")
	chain.AddBlock("third BLock")

	for _, block := range chain.BLocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}

}
