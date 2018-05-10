package common

import (
    "crypto/sha256"
    "strconv"
    "encoding/hex"
    "fmt"
)

type Block struct {
    Height uint64
    Data string
    PrevHash, Hash []byte
}

func (b *Block) Content() []byte {
    content := []byte(strconv.FormatUint(b.Height, 10))
    content = append(content, []byte(b.Data)...)
    content = append(content, b.PrevHash...)
    return content
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
    hash := sha256.Sum256(b.Content())
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
