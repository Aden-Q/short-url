package router

import (
	"time"

	_ "github.com/Aden-Q/short-url/docs"
	"github.com/Aden-Q/short-url/internal/cache"
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/handler"
	"github.com/Aden-Q/short-url/internal/middleware"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	RequestTimeout time.Duration
	DB             db.Engine
	Redis          redis.Client
	Cache          cache.Cache
}

type Router struct {
	*gin.Engine
	config Config
}

func New(config Config) *Router {
	// default gin engine with Logger and Recovery middleware
	r := Router{
		Engine: gin.Default(),
		config: config,
	}

	// context timeout for the handler chain
	r.Use(middleware.RequestTimeout(config.RequestTimeout))

	// attach a global middleware in the handler chain to enforce database connection
	r.Use(middleware.DB(config.DB))

	// attach a global middleware in the handler chain to enforce redis connection
	r.Use(middleware.Redis(config.Redis, config.Cache))

	// swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
