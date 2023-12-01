package router

import (
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/handler"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/gin-gonic/gin"
)

type Config struct {
	DB    db.Engine
	Redis redis.Client
}

type Router struct {
	*gin.Engine
	config Config
}

// DBMiddleware is a middleware to establish mysql database connection
func DBMiddleware(db db.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConn", db)
		c.Next()
	}
}

// RedisMiddleware is a middleware to establish redis connection
func RedisMiddleware(redis redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redisConn", redis)
		c.Next()
	}
}

func NewRouter(config Config) *Router {
	// default gin engine with Logger and Recovery middleware
	r := Router{
		Engine: gin.Default(),
		config: config,
	}

	// attach a global middleware in the handler chain to enforce database connection
	r.Use(DBMiddleware(config.DB))

	// attach a global middleware in the handler chain to enforce redis connection
	r.Use(RedisMiddleware(config.Redis))

	// the health check endpoint
	r.GET("/health", handler.Health)

	apiv1 := r.Group("/api/v1")
	{
		// The POST endpoint for shortening a long URL
		// We bind the request string because we don't want to cache the request
		apiv1.POST("/data/shorten", handler.ShortenHandler)

		// The GET endpoint for redirecting a short URL to a long URL, returns the long URL
		// We use URI binding because we want to cache the request
		apiv1.GET("/:shortURL", handler.RedirectHandler)
	}

	return &r
}
