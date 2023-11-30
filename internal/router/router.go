package router

import (
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/handler"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	DB *db.DBEngine
}

type Router struct {
	*gin.Engine
}

func NewRouter(config RouterConfig) *Router {
	// default gin engine with Logger and Recovery middleware
	r := Router{
		Engine: gin.Default(),
	}

	// the health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	apiv1 := r.Group("/api/v1")
	{
		// The POST endpoint for shortening a long URL
		// We bind the request string because we don't want to cache the request
		apiv1.POST("/data/shorten", handler.ShortenHandler(config.DB))

		// The GET endpoint for redirecting a short URL to a long URL, returns the long URL
		// We use URI binding because we want to cache the request
		apiv1.GET("/:shortURL", handler.RedirectHandler(config.DB))
	}

	return &r
}
