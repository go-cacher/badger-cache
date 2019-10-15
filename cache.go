package cache

import (
	"errors"
	"os"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/gocacher/cacher"
)

var DefaultPath = "cache"

type BadgerCache struct {
	path string
	db   *badger.DB
}

func init() {
	cacher.Register(NewBadgerCache(DefaultPath))
}

func (b BadgerCache) Get(key string) ([]byte, error) {
	var bytes []byte
	if err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			bytes = make([]byte, len(val))
			copy(bytes, val)
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return bytes, nil
}

func (b BadgerCache) GetD(key string, v []byte) []byte {
	if ret, err := b.Get(key); err == nil {
		return ret
	}
	return v
}

func (b *BadgerCache) Set(key string, val []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), val)
		return err
	})

}

func (b *BadgerCache) SetWithTTL(key string, val []byte, ttl int64) error {
	return b.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), val).WithTTL(time.Duration(ttl))
		err := txn.SetEntry(e)
		return err
	})
}

func (b BadgerCache) Has(key string) (bool, error) {
	var has bool
	if err := b.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		if err == nil {
			has = true
			return nil
		} else if errors.Is(err, badger.ErrKeyNotFound) {
			has = false
			return nil
		}
		return err
	}); err != nil {
		return false, err
	}
	return has, nil
}

func (b *BadgerCache) Delete(key string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
}

func (b *BadgerCache) Clear() error {
	return os.RemoveAll(b.path)
}

func (b BadgerCache) GetMultiple(keys ...string) (map[string][]byte, error) {
	vals := make(map[string][]byte, len(keys))
	err := b.db.View(func(txn *badger.Txn) error {
		for _, key := range keys {
			item, err := txn.Get([]byte(key))
			if err != nil {
				return err
			}
			var bytes []byte
			err = item.Value(func(val []byte) error {
				bytes = make([]byte, len(val))
				vals[key] = bytes
				return nil
			})
			if err != nil {
				return err
			}
			return nil
		}
		return nil
	})
	return vals, err
}

func (b BadgerCache) SetMultiple(values map[string][]byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		for k, v := range values {
			err := txn.Set([]byte(k), v)
			if err != nil {
				return err
			}
		}
		return nil

	})
}

func (b BadgerCache) DeleteMultiple(keys ...string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		for _, key := range keys {
			err := txn.Delete([]byte(key))
			if err != nil {
				return err
			}
		}
		return nil
	})
}
