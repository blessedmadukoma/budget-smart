package cache

import (
	"context"
	"time"

	rcache "github.com/blessedmadukoma/budgetsmart/engine/internal/common/cache/redis"
)

type Cache interface {
	Set(ctx context.Context, key string, data interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, data interface{}) error
	Delete(ctx context.Context, key string) error
}

func NewCache(urls []string) (Cache, error) {
	ca, err := rcache.NewRedisCache(urls)
	if err != nil {
		return nil, err
	}

	return ca, nil
}
