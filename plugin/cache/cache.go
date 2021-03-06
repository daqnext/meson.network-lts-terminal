package cache

import (
	"github.com/universe-30/UCache"
)

var cache *UCache.Cache

func Init() {
	if cache != nil {
		return
	}
	cache = UCache.New()
}

func GetSingleInstance() *UCache.Cache {
	return cache
}
