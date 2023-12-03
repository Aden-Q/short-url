package middleware

import (
	"github.com/Aden-Q/short-url/internal/cache"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/gin-gonic/gin"
)

// RedisMiddleware is a middleware to establish redis connection, and make sure cache is instantiated
func Redis(redis redis.Client, cache cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redisConn", redis)
		c.Set("cache", cache)
		c.Next()
	}
}
