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

type BlockIndex struct {
    Header structs.BlockHeader
    Location string
    Offset uint64
    Size uint64
}

type Block struct {
    Header BlockHeader
    Data string
}

func (header *BlockHeader) String() string {
    return fmt.Sprintf(
        "{Height=%d, Timestamp=%d, PrevHash=%v, Hash=%v}",
        header.Height,
        header.Timestamp,
        hex.EncodeToString(header.PrevHash),
        hex.EncodeToString(header.Hash),
    )
}

func (b *Block) Content() []byte {
    content := common.UintToHex(b.Header.Height)
    content = append(content, []byte(b.Header.Data)...)
    return append(content, b.Header.PrevHash...)
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

func (b *Block) String() string {
    return fmt.Sprintf(
        "Header=%s, Data=%v}", b.Header.String(), b.Data,
    )
}
