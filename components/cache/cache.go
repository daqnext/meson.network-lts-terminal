package cache

import (
	"sync"

	"github.com/universe-30/UCache"
)

var cache *UCache.Cache
var once sync.Once

func GetSingleInstance() *UCache.Cache {
	once.Do(func() {
		cache = UCache.New()
	})
	return cache
}
