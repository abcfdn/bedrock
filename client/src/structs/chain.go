package structs

import (
    "encoding/hex"
)

type BlockChain struct {
    Blocks []*Block
}

var blockchain *BlockChain

func CurrentBlockchain() *BlockChain {
    return blockchain
}

func (blockchain *BlockChain) AddBlock(block *Block) {
    blockchain.Blocks = append(blockchain.Blocks, block)
}

func String() string {
    bs := make([]byte, 0)
    for _, block := range blockchain.Blocks {
        bs = append(bs, []byte(block.String())...)
        bs = append(bs, '\n')
    }
    return string(bs)
}

func NewBlockchain(genesisBlock *Block) *BlockChain {
    return &BlockChain{[]*Block{genesisBlock}}
}

func ReplaceChain(newBlockchain *BlockChain) {
    if len(newBlockchain.Blocks) > len(blockchain.Blocks) {
        blockchain = newBlockchain
    }
}

func (bc *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
    unspentOutputs := make(map[string][]int)
    unspentTXs := bc.FindUnspentTransactions(address)
    accumulated := 0

Work:
    for _, tx := range unspentTXs {
        txID := hex.EncodeToString(tx.ID)

        for outIdx, out := range tx.Vout {
            if out.CanBeUnlockedWith(address) && accumulated < amount {
                accumulated += out.Value
                unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

                if accumulated >= amount {
                    break Work
                }
            }
        }
    }

    return accumulated, unspentOutputs
}

func (bc *BlockChain) FindUnspentTransactions(address string) []Transaction {
  // TODO(jiazhouxiao)
  var unspentTXs []Transaction
  return unspentTXs
}

// MineBlock mines a new block with the provided transactions
func (bc *BlockChain) MineBlock(transactions []*Transaction) *Block {
    // TODO(jiazhouxiao)
    return NewGenesisBlock("")
}