package structs

import (
    "crypto/sha256"
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

func (b *Block) Content() []byte {
    content := common.UintToHex(b.Height)
    content = append(content, []byte(b.Data)...)
    return append(content, b.PrevHash...)
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
    nonce, hash := pow.Run()
    block.Hash = hash[:]
    block.Nonce = nonce
    b.Hash = hash[:]
}

func NewGenesisBlock() *Block {
    block := &Block{1, "Genesis Block", []byte{}, []byte{}}
    block.SetHash()
    return block
}

func NewBlock(prev *Block, data string) *Block {
    block := &Block{prev.Height + 1, data, prev.Hash, []byte{}}
    block.SetHash()
    return block
}
