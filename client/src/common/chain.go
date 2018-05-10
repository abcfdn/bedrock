package common

type BlockChain struct {
    Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
    prevBlock := bc.Blocks[len(bc.Blocks) - 1]
    block := NewBlock(prevBlock, data)
    bc.Blocks = append(bc.Blocks, block)
}

func (bc *BlockChain) String() string {
    bs := make([]byte, 0)
    for _, block := range bc.Blocks {
        bs = append(bs, []byte(block.String())...)
        bs = append(bs, '\n')
    }
    return string(bs)
}

func NewBlockChain() *BlockChain {
    return &BlockChain{[]*Block{NewGenesisBlock()}}
}
