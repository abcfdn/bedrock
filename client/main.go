package main

import (
    "fmt"
    "structs"
)

func main() {
    bc := structs.NewBlockChain()
    bc.AddBlock("first")
    bc.AddBlock("second")
    fmt.Println(bc);
}
