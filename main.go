package main

import (
	"fmt"
	"github.com/tensor-programming/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First BLock after G")
	chain.AddBlock("second BLock after G")
	chain.AddBlock("third BLock after G")

	for _, block := range chain.BLocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}

}
