package easy

import (
	"github.com/gocacher/badger-cache/v2"
	"github.com/gocacher/cacher"
)

func init() {
	cacher.Register(cache.New())
}
