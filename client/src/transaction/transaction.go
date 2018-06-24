package transaction

import (
	"fmt"
	"structs"
)

func send(from, to string, amount int) {
	genesis := structs.NewGenesisBlock(to)
	bc := structs.NewBlockchain(genesis)
	// defer bc.db.Close()

	tx := structs.NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*structs.Transaction{tx})
	fmt.Println("Success!")
}