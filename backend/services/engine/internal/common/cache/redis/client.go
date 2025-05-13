package rcache

import (
	"context"
	"time"

	"github.com/blessedmadukoma/budgetsmart/engine/pkg/rdb"
	"github.com/go-redis/cache/v9"
)

const cacheSize = 128000

type RedisCache struct {
	cache *cache.Cache
}

func NewRedisCache(addresses []string) (*RedisCache, error) {
	client, err := rdb.NewClient(addresses)
	if err != nil {
		return nil, err
	}

	c := cache.New(&cache.Options{
		Redis: client.Client(),
		// LocalCache: cache.NewTinyLFU(cacheSize, -1),
	})

	r := &RedisCache{cache: c}

	return r, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, data interface{}, ttl time.Duration) error {
	return r.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: data,
		TTL:   ttl,
	})
}

func (r *RedisCache) Get(ctx context.Context, key string, data interface{}) error {
	return r.cache.Get(ctx, key, &data)
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.cache.Delete(ctx, key)
}
