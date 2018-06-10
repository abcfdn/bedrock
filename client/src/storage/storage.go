package storage

import (
    "structs"
    "common/logger"
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/Rican7/retry"
)

func getDB(dbpath string) leveldb.DB {
    db, err := leveldb.OpenFile(dbpath, nil)
    if err != nil {
        logger.Fatal()
    }
    return db
}

// key = block.Header.Hash
func GetBlockIndexRecord(db DB, key byte[]) (structs.BlockIndexRecord, error) {
    err := retry.Retry(func(attempt uint) error {
        record, err := db.Get(key, nil)
        if nil != err {
            logger.Error("Failed to get block %s with error %v", key, err)
        }
    }, strategy.Limit(3))

    return record, err
}

func CommitBlock(db DB, index structs.BlockIndex) error {
    return retry.Retry(func(attempt uint) error {
        err := db.Put([]byte("key"), index, nil)
        if nil != err {
            logger.Error("Failed to save index %+v with error %v", index, err)
        }
    }, strategy.Limit(3))
}

func FlushBlocksToLocal(filename string, blocks structs.Block[]) error {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        logger.Panic("Failed to open file %s with error %v", filename, err)
    }

    defer f.Close()
    if _, err = f.WriteString(text); err != nil {
        logger.Panic("Failed to commit  ", logFilePath, err)
    }
    return err
}
