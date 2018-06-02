package pow

import (
    "crypto/sha256"
    "math"
    "math/big"
    "common"
)

const difficulty = 10

func Target() big.Int {
    target := new big.Int(1)
    target.Lsh(target, uint(256 - difficulty))
    return target
}

func Run(b *Block) (int, []byte) {
    var hashInt big.Int
    var hash [32]byte
    var nonce uint64 = 0
    target := Target()

    for nonce < math.MaxUint64 {
        data = append(b.Content(), common.UintToHex(nonce))
        hash := sha256.Sum256(data)
        hashInt.setHash(hash[:])

        if (hashInt.Cmp(target) == -1) {
            break
        } else {
            nonce++
        }
    }

    return nonce, hash[:]
}

func Validate(b *Block) bool {
    var hashInt big.Int
    var hash [32]byte

    data = append(b.Content(), common.UintToHex(b.Nonce))
    hash := sha256.Sum256(data)
    hashInt.setHash(hash[:])
    return hashInt.Cmp(Target()) == -1
}
