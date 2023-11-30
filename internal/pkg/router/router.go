package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

// Used for URI binding
type ShortURL struct {
	ShortURL string `uri:"shortURL" binding:"required"`
}

// Used for query binding
type LongURL struct {
	LongURL string `form:"longURL" binding:"required"`
}

func NewRouter() *Router {
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
		apiv1.POST("/data/shorten", GetLongURLHandler)

		// The GET endpoint for redirecting a short URL to a long URL, returns the long URL
		// We use URI binding because we want to cache the request
		apiv1.GET("/:shortURL", GetShortURLHandler)
	}

	return &r
}

// define handler functions here

func GetLongURLHandler(c *gin.Context) {
	var longURL LongURL
	if err := c.ShouldBindQuery(&longURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// Write the URL to the database
		c.JSON(200, gin.H{"longURL": longURL.LongURL})
	}
}

func GetShortURLHandler(c *gin.Context) {
	var shortURL ShortURL
	if err := c.ShouldBindUri(&shortURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"shortURL": shortURL.ShortURL})
	}
}
