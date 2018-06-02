package common

import (
    "strconv"
)

func UintToHex(n uint64) []byte {
    return []byte(strconv.FormatInt(n, 16))
}
