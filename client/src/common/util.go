package common

import (
    "strconv"
)

// Interface that exposed shared block methods.
type Block interface {
    Content() []byte
    GetHeight() uint64
    GetData() string
    GetHash() []byte
    GetPrevHash() []byte
    GetNonce() uint64
}

func UintToHex(n uint64) []byte {
    return []byte(strconv.FormatInt(int64(n), 16))
}
