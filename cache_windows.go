package cache

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/gocacher/cacher"
)

func newBadgerCache(path string) cacher.Cacher {
	options := badger.DefaultOptions(path)
	options.Truncate = true
	db, err := badger.Open(options)
	if err != nil {
		panic(err)
	}
	return &BadgerCache{
		path: path,
		db:   db,
	}
}
