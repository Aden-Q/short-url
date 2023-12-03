package cache

import (
	"context"
	"time"

	"github.com/Aden-Q/short-url/internal/redis"
	redisCache "github.com/go-redis/cache/v9"
)

var ErrCacheMiss = redisCache.ErrCacheMiss

type Config struct {
	Redis redis.Client
}

type Cache interface {
	// Get gets the value from cache for the given key
	Get(key string, value interface{}) error
	// Set sets the value in the cache for the given key
	Set(key string, value interface{}) error
	// Delete deletes the value in the cache for the given key
	Delete(key string) error
	// Once gets the value from cache for the given key, if the key does not exist, call f() to get the value and set it in the cache
	Once(key string, value interface{}, f func() (interface{}, error)) error
}

type cache struct {
	// ctx is required to make redis calls, if ctx is canceled, redis calls will not succeed
	ctx context.Context
	*redisCache.Cache
}

func New(config Config) Cache {
	return &cache{
		ctx: config.Redis.Context(),
		Cache: redisCache.New(&redisCache.Options{
			Redis: config.Redis.GetClient(),
			// Cache 10k keys in local memory for 1 minute
			LocalCache: redisCache.NewTinyLFU(1000, time.Minute),
		}),
	}
}

func (c *cache) Get(key string, value interface{}) error {
	return c.Cache.Get(c.ctx, key, value)
}

func (c *cache) Set(key string, value interface{}) error {
	return c.Cache.Set(&redisCache.Item{
		Ctx:   c.ctx,
		Key:   key,
		Value: value,
		TTL:   time.Hour,
	})
}

func (c *cache) Delete(key string) error {
	return c.Cache.Delete(c.ctx, key)
}

func (c *cache) Once(key string, value interface{}, f func() (interface{}, error)) error {
	return c.Cache.Once(&redisCache.Item{
		Ctx:   c.ctx,
		Key:   key,
		Value: value,
		Do: func(i *redisCache.Item) (interface{}, error) {
			return f()
		},
	})
}
