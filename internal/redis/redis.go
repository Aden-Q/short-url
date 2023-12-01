package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Config is the configuration for redis client
type Config struct {
	Addr string
}

// Client is the interface for redis client
type Client interface {
	// Get runs the redis GET command
	Get(key string) (string, error)
}

// client is the redis client struct that implements the Client interface
type client struct {
	// ctx is required to make redis calls, if ctx is canceled, redis calls will not succeed
	ctx    context.Context
	config Config
	*redis.Client
}

// NewClient creates a new redis client
func NewClient(ctx context.Context, config Config) (Client, error) {
	r := &client{
		ctx:    ctx,
		config: config,
		Client: redis.NewClient(&redis.Options{
			Addr: config.Addr,
		}),
	}

	if err := r.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *client) Get(key string) (string, error) {
	return r.Client.Get(r.ctx, key).Result()
}
