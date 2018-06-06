package main

import (
    "fmt"
    "structs"
)

func main() {
	structs.NewBlockChain()

	blockchain := structs.CurrentBlockchain()
    blockchain.AddBlock(structs.NewBlock(blockchain.Blocks[len(blockchain.Blocks) - 1], "first"))
    blockchain.AddBlock(structs.NewBlock(blockchain.Blocks[len(blockchain.Blocks) - 1], "second"))
    fmt.Println(blockchain);
}
