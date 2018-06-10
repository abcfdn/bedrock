package pow

import (
    "structs"
    "common"
    "crypto/sha256"
    "encoding/hex"
    "math/big"
    "strings"
    "time"
    "fmt"
)

const difficulty = 1

func Target() *big.Int {
    target := big.NewInt(1)
    target.Lsh(target, uint(256 - difficulty))
    return target
}

// Blockchain Validation
func CalculateHash(block structs.Block, nonce uint64) (string, []byte) {
    record := []byte(string(block.Header.Height()) + block.Header.Data() + string(nonce))
    record = append(record, block.Header.PrevHash()...)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed), hashed
}

func isValidHash(hash string, difficulty int) bool {
    prefix := strings.Repeat("0", difficulty)
    return strings.HasPrefix(hash, prefix)
}

func Run(block structs.Block) (uint64, []byte) {
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

func Validate(b structs.Block) bool {
    var hashInt big.Int
    var hash [32]byte

    data := append(b.Content(), common.UintToHex(b.Nonce())...)
    hash = sha256.Sum256(data)
    hashInt.SetBytes(hash[:])
    return hashInt.Cmp(Target()) == -1
}

func SetHash(b Block) {
    nonce, hash := pow.Run(b)
    b.Header.Hash = hash[:]
    b.Header.Nonce = nonce
}

func IsBlockValid(newBlock *Block,  prevBlock *Block) bool {
    if prevBlock.Header.Height + 1 != newBlock.Header.Height {
      return false
    }

    if len(newBlock.Header.PrevHash) != len(prevBlock.Header.Hash) {
      return false
    }

    for i := range newBlock.Header.Hash {
        if prevBlock.Header.Hash[i] != newBlock.Header.PrevHash[i] {
          return false
        }
    }

    _, hash := pow.CalculateHash(newBlock, newBlock.Header.Nonce)
    for i := range hash {
        if hash[i] != newBlock.Header.Hash[i] {
            return false
        }
    }

    return true
}
