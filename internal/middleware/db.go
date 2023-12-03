package middleware

import (
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/gin-gonic/gin"
)

// DB is a middleware to establish mysql database connection
func DB(db db.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConn", db)
		c.Next()
	}
}
