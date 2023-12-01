package handler

import (
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	// make sure mysql connection is established is is stored into the context
	_, ok := c.MustGet("dbConn").(db.Engine)
	if !ok {
		c.JSON(500, gin.H{"message": "mysql connection not established"})
		return
	}

	// make sure redis connection is established is is stored into the context
	_, ok = c.MustGet("redisConn").(redis.Client)
	if !ok {
		c.JSON(500, gin.H{"message": "redis connection not established"})
		return
	}

	c.JSON(200, gin.H{"message": "healthy"})
}
