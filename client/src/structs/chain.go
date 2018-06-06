package structs

type BlockChain struct {
    Blocks []*Block
}

var blockchain *BlockChain

func CurrentBlockchain() *BlockChain {
    return blockchain
}

func (bc *BlockChain) AddBlock(block *Block) {
    if !IsBlockValid(block, bc.Blocks[len(bc.Blocks) - 1]) {
        return
    }

    bc.Blocks = append(bc.Blocks, block)
}

func String() string {
    bs := make([]byte, 0)
    for _, block := range blockchain.Blocks {
        bs = append(bs, []byte(block.String())...)
        bs = append(bs, '\n')
    }
    return string(bs)
}

func NewBlockChain() {
    blockchain = &BlockChain{[]*Block{NewGenesisBlock()}}
}

func ReplaceChain(newBlockchain *BlockChain) {
    if len(newBlockchain.Blocks) > len(blockchain.Blocks) {
        blockchain = newBlockchain
    }
}
