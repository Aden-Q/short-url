package handler

import (
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	_, ok := c.MustGet("dbConn").(*db.DBEngine)
	if !ok {
		c.JSON(500, gin.H{"message": "dbConn not found"})
		return
	}

	c.JSON(200, gin.H{"message": "pong"})
}
