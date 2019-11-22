package easy

import (
	"github.com/gocacher/badger-cache"
	"github.com/gocacher/cacher"
)

func init() {
	cacher.Register(cache.New())
}
