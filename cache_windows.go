package cache

import (
	"github.com/dgraph-io/badger"
	"github.com/gocacher/cacher"
)

func NewBadgerCache(path string) cacher.Cacher {
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
