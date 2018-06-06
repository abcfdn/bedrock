package structs

import (
    "encoding/hex"
    "consensus/pow"
    "common"
    "fmt"
)

type Block struct {
    Height uint64
    Data string
    PrevHash, Hash []byte
    Nonce uint64
}

func (b *Block) String() string {
  return fmt.Sprintf(
      "Block %d: Data: %v, PrevHash: %v, Hash: %v}",
      b.Height,
      b.Data,
      hex.EncodeToString(b.PrevHash),
      hex.EncodeToString(b.Hash),
  )
}

func (b *Block) SetHash() {
    nonce, hash := pow.Run(b)
    b.Hash = hash[:]
    b.Nonce = nonce
}

func NewGenesisBlock() *Block {
    block := &Block{1, "Genesis Block", []byte{}, []byte{}, 0}
    block.SetHash()
    return block
}

func NewBlock(prev *Block, data string) *Block {
    block := &Block{prev.Height + 1, data, prev.Hash, []byte{}, 0}
    block.SetHash()
    return block
}

func IsBlockValid(newBlock *Block,  prevBlock *Block) bool {
    if prevBlock.Height + 1 != newBlock.Height {
      return false
    }

    if len(newBlock.PrevHash) != len(prevBlock.Hash) {
      return false
    }

    for i := range newBlock.Hash {
        if prevBlock.Hash[i] != newBlock.PrevHash[i] {
          return false
        }
    }

    _, hash := pow.CalculateHash(newBlock, newBlock.Nonce)
    for i := range hash {
        if hash[i] != newBlock.Hash[i] {
            return false
        }
    }

    return true
}

// Block Interface

func (b *Block) Content() []byte {
    content := common.UintToHex(b.Height)
    content = append(content, []byte(b.Data)...)
    return append(content, b.PrevHash...)
}

func (b *Block) GetHeight() uint64 {
    return b.Height
}

func (b *Block) GetData() string {
    return b.Data
}

func (b *Block) GetHash() []byte {
    return b.Hash
}

func (b *Block) GetPrevHash() []byte {
    return b.PrevHash
}

func (b *Block) GetNonce() uint64 {
    return b.Nonce
}
