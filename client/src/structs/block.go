package structs

import (
    "common"
    "encoding/hex"
    "fmt"
)

type BlockHeader struct {
    Height, Timestamp, Nonce uint64
    PrevHash, Hash []byte
}

type Block struct {
    Header *BlockHeader
    Data string
}

func (b *Block) String() string {
    return fmt.Sprintf(
        "Block %d: Timestamp: %d, PrevHash: %v, Hash: %v, Data: %v}",
        b.Header.Height,
        b.Header.Timestamp,
        hex.EncodeToString(b.Header.PrevHash),
        hex.EncodeToString(b.Header.Hash),
        b.Data,
    )
}

func NewBlock(prev *Block, data string) *Block {
    header := &BlockHeader{prev.Header.Height + 1, common.MakeTimestamp(), 0, prev.Header.Hash, []byte{}}
    block := &Block{header, data}
    return block
}
