package main

import (
	log "github.com/gitsang/golog"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"go.uber.org/zap"
	"strconv"
)

func main() {
	log.InitLogger(
		log.WithLogLevel("DEBUG"),
		log.WithEncoderType("console"),
		log.WithEnableHttp(true),
		log.WithHttpPort(8888),
		log.WithLogFile("./leveldb.log"),
		log.WithLogFileCompress(true),
	)

	// open db
	db, err := leveldb.OpenFile("./leveldb.db", nil)
	if err != nil {
		log.Error("open db failed", zap.Error(err))
		return
	}
	defer func() {
		_ = db.Close()
	}()
	log.Debug("open db success")

	// put
	for i := 0; i < 20; i++ {
		err = db.Put([]byte("key" + strconv.Itoa(i)), []byte("value" + strconv.Itoa(i)), nil)
		if err != nil {
			log.Error("db put failed", zap.Error(err))
		}
		log.Debug("db put success", zap.Int("i", i))
	}
	log.Debug("db put success")

	// iter
	iter := db.NewIterator(&util.Range{
		Start: []byte("key2"),
		Limit: []byte("key7"),
	}, nil)
	defer iter.Release()
	for ok := iter.Seek([]byte("key2")); ok; ok = iter.Next() {
		log.Info("iter", zap.ByteString("key", iter.Key()), zap.ByteString("value", iter.Value()))
	}

	// delete
	//_ = db.Delete([]byte("key1"), nil)
}
