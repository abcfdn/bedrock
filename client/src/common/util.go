package common

import (
    "strconv"
    "log"
    "os"
)

func UintToHex(n uint64) []byte {
    return []byte(strconv.FormatInt(int64(n), 16))
}
