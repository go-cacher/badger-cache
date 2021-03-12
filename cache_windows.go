package cache

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/gocacher/cacher"
)

func newBadgerCache(path string) cacher.Cacher {
	options := badger.LSMOnlyOptions(path)
	options.CompactL0OnClose = false
	db, err := badger.Open(options)
	if err != nil {
		panic(err)
	}
	return &BadgerCache{
		path: path,
		db:   db,
	}
}
