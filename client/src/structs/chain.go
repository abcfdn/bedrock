package structs

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

func NewBlockchain(genesisBlock *Block) {
    blockchain = &BlockChain{[]*Block{genesisBlock}}
}

func ReplaceChain(newBlockchain *BlockChain) {
    if len(newBlockchain.Blocks) > len(blockchain.Blocks) {
        blockchain = newBlockchain
    }
}
