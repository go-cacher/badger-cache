package cache

import (
	"testing"

	"github.com/gocacher/cacher"
)

func init() {
	cacher.Register(New())
}

func TestNewBadgerCache(t *testing.T) {
	e := cacher.Set("name", []byte("test"))
	if e != nil {
		t.Fatal(e)
	}
}
