package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config is the configuration for redis client
type Config struct {
	Addr string
}

// Client is the interface for redis client
type Client interface {
	// GetContext returns the context attached to the redis client instance
	Context() context.Context
	// GetClient returns the redis client instance
	GetClient() *redis.Client
	// Get runs the redis GET command
	Get(key string) (string, error)
	// Set runs the redis SET command, to set an expiration, use the SetEX command
	Set(key string, value interface{}) error
	// SetEX runs the redis SETEX command
	SetEX(key string, value interface{}, expiration time.Duration) error
	// Close closes the client, releasing any open resources
	Close() error
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

// Context returns the context attached to the redis client instance
func (r *client) Context() context.Context {
	return r.ctx
}

func (r *client) GetClient() *redis.Client {
	return r.Client
}

// Get runs the redis GET command
func (r *client) Get(key string) (string, error) {
	return r.Client.Get(r.ctx, key).Result()
}

// Set runs the redis SET command, to set an expiration, use the SetEX command
func (r *client) Set(key string, value interface{}) error {
	return r.Client.Set(r.ctx, key, value, 0).Err()
}

// SetEX runs the redis SETEX command
func (r *client) SetEX(key string, value interface{}, expiration time.Duration) error {
	return r.Client.SetEx(r.ctx, key, value, expiration).Err()
}

// Close closes the client, releasing any open resources
func (r *client) Close() error {
	return r.Client.Close()
}
