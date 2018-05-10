package main

import (
    "fmt"
    "common"
)

func main() {
    bc := common.NewBlockChain()
    bc.AddBlock("first")
    bc.AddBlock("second")

}
