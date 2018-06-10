package structs

import (
    "encoding/hex"
    "consensus/pow"
    "common"
    "time"
    "fmt"
)

type BlockHeader struct {
    Height, Timestamp, Nonce uint64
    PrevHash, Hash []byte
}

type Block struct {
    Header BlockHeader
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

func NewGenesisBlock() *Block {
    header := &BlockHeader{1, time.now(), 0, []byte{}, []byte{}}
    block := &Block{header, "Genesis Block"}
    return block
}

func NewBlock(prev *Block, data string) *Block {
    header := &BlockHeader{prev.Header.Height + 1, time.now(), 0, prevHash, []byte{}}
    block := &Block{header, data}
    return block
}

func (b *Block) Content() []byte {
    content := common.UintToHex(b.Header.Height)
    content = append(content, []byte(b.Header.Data)...)
    return append(content, b.Header.PrevHash...)
}
