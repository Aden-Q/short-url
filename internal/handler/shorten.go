package handler

import (
	"net/http"

	"github.com/Aden-Q/short-url/internal/db"
	"github.com/gin-gonic/gin"
)

// Used for query binding
type longURL struct {
	longURL string `form:"longURL" binding:"required"`
}

// ShortenHandler shortens a long URL
func ShortenHandler(db *db.DBEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var longURL longURL
		if err := c.ShouldBindQuery(&longURL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			// Write the URL to the database
			c.JSON(200, gin.H{"longURL": longURL.longURL})
		}
	}
}
