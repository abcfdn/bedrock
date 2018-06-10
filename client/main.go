package main

import (
    "fmt"
    "consensus/pow"
    "structs"
)

func main() {
	pow.MineGenesisBlock()

	blockchain := structs.CurrentBlockchain()
	block1 := structs.NewBlock(structs.CurrentBlockchain().Blocks[len(blockchain.Blocks) - 1], "first")
	pow.SetHash(block1)

	blockchain = structs.CurrentBlockchain()
	block2 := structs.NewBlock(blockchain.Blocks[len(blockchain.Blocks) - 1], "second")
	pow.SetHash(block2)

	blockchain = structs.CurrentBlockchain()
    fmt.Println(blockchain);
}
