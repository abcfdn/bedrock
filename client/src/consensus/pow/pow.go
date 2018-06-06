package pow

import (
    "common"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    // "math"
    "math/big"
    "strings"
    "time"
)

const difficulty = 1

func Target() *big.Int {
    target := big.NewInt(1)
    target.Lsh(target, uint(256 - difficulty))
    return target
}

// Blockchain Validation
func CalculateHash(block common.Block, nonce uint64) (string, []byte) {
    record := []byte(string(block.GetHeight()) + block.GetData() + string(nonce))
    record = append(record, block.GetPrevHash()...)
    h := sha256.New() 
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed), hashed
}

func isValidHash(hash string, difficulty int) bool {
    prefix := strings.Repeat("0", difficulty)
    return strings.HasPrefix(hash, prefix)
}

func Run(block common.Block) (uint64, []byte) {
    var nonce uint64 = 0
    var hashBytes []byte

    for nonce = 0; ; nonce++ {
        hashString, hash := CalculateHash(block, nonce)
        if !isValidHash(hashString, difficulty) {
            fmt.Println(hashString, " do more work!")
            time.Sleep(time.Second)
            continue
        } else {
            hashBytes = hash
            fmt.Println(hashString, " work done!")
            break;
        }
    }

    return nonce, hashBytes
}

// func Run(b common.Block) (uint64, []byte) {
//     var hashInt big.Int
//     var hash [32]byte
//     var nonce uint64 = 0
//     target := Target()

//     for nonce < math.MaxUint64 {
//         data := append(b.Content(), common.UintToHex(nonce)...)
//         hash := sha256.Sum256(data)
//         hashInt.SetBytes(hash[:])

//         if (hashInt.Cmp(target) == -1) {
//             break
//         } else {
//             nonce++
//         }
//     }

//     return nonce, hash[:]
// }

func Validate(b common.Block) bool {
    var hashInt big.Int
    var hash [32]byte

    data := append(b.Content(), common.UintToHex(b.GetNonce())...)
    hash = sha256.Sum256(data)
    hashInt.SetBytes(hash[:])
    return hashInt.Cmp(Target()) == -1
}
