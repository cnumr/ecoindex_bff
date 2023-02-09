package config

import (
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var CACHE *cache.Cache

func GetCache() *cache.Cache {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ENV.CacheDsn,
		},
	})

	return cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(10000, time.Duration(ENV.CacheTtl)*time.Second),
	})
}
