package caches

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Caches struct {
	StmCache *cache.Cache
}

var caches *Caches

func InitCaches() {
	caches = &Caches{
		StmCache: cache.New(8*time.Hour, 10*time.Hour),
	}
}

func GetCaches() *Caches {
	return caches
}
