package cache

import (
	"github.com/universe-30/UCache"
)

var cache *UCache.Cache

func Init() {
	cache = UCache.New()
}

func GetSingleInstance() *UCache.Cache {
	return cache
}
