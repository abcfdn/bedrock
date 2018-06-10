package pow

import (
    "common"
    "structs"
    "crypto/sha256"
    "encoding/hex"
    "strings"
    "time"
    "fmt"
)

const difficulty = 1

func SetHash(b *structs.Block) {
    nonce, hash := run(b)
    b.Header.Hash = hash[:]
    b.Header.Nonce = nonce

    addBlock(b)
}

func MineGenesisBlock() {
    header := &structs.BlockHeader{1, common.MakeTimestamp(), 0, []byte{}, []byte{}}
    genesisBlock := &structs.Block{header, "Genesis Block"}
    SetHash(genesisBlock)
}

// calculateHash does proof of work for the given block.
func calculateHash(block *structs.Block, nonce uint64) (string, []byte) {
    record := []byte(string(block.Header.Height) + block.Data + string(nonce))
    record = append(record, block.Header.PrevHash...)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed), hashed
}

func isValidHash(hash string, difficulty int) bool {
    prefix := strings.Repeat("0", difficulty)
    return strings.HasPrefix(hash, prefix)
}

func run(block *structs.Block) (uint64, []byte) {
    var nonce uint64 = 0
    var hashBytes []byte

    for nonce = 0; ; nonce++ {
        hashString, hash := calculateHash(block, nonce)
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

func addBlock(block *structs.Block) {
    if structs.CurrentBlockchain() == nil {
        // Attach the valid block into the blockchain as the genesis block.
        structs.NewBlockchain(block)
    }

    blockchain := structs.CurrentBlockchain()
    if !isBlockValid(block, blockchain.Blocks[len(blockchain.Blocks) - 1]) {
        return
    }

    blockchain.AddBlock(block)
}

func isBlockValid(newBlock *structs.Block,  prevBlock *structs.Block) bool {
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

    _, hash := calculateHash(newBlock, newBlock.Header.Nonce)
    for i := range hash {
        if hash[i] != newBlock.Header.Hash[i] {
            return false
        }
    }

    return true
}
