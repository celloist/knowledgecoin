package main

import (
	blockchain "athmare/knowledgecoin/core"
	"fmt"
	"strconv"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("send 1 KNC to Pedro")
	bc.AddBlock("send 2 KNC to Pedro")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
