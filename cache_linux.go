package cache

import (
	"github.com/dgraph-io/badger"
	"github.com/gocacher/cacher"
)

func NewBadgerCache(path string) cacher.Cacher {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		panic(err)
	}
	return &BadgerCache{
		path: path,
		db:   db,
	}
}
