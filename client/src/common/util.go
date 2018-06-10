package common

import (
    "strconv"
    "time"
)

func UintToHex(n uint64) []byte {
    return []byte(strconv.FormatInt(int64(n), 16))
}

func MakeTimestamp() uint64 {
    return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}
