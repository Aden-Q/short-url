package handler

import (
	"net/http"

	"github.com/Aden-Q/short-url/internal/cache"
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	// make sure mysql connection is established is is stored into the context
	_, ok := c.MustGet("dbConn").(db.Engine)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "mysql connection not established"})
		return
	}

	// make sure redis connection is established is is stored into the context
	_, ok = c.MustGet("redisConn").(redis.Client)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "redis connection not established"})
		return
	}

	// make sure cache is instantiated
	_, ok = c.MustGet("cache").(cache.Cache)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "healthy"})
}
