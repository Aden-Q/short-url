package handler

import (
	"net/http"

	"github.com/Aden-Q/short-url/internal/db"
	"github.com/gin-gonic/gin"
)

// Used for URI binding
type bindingURI struct {
	ShortURL string `uri:"shortURL" binding:"required"`
}

// ShortenHandlerFunc shortens a long URL
func RedirectHandler(c *gin.Context) {
	_, ok := c.MustGet("dbConn").(*db.DBEngine)
	if !ok {
		c.JSON(500, gin.H{"message": "dbConn not found"})
		return
	}

	var binding bindingURI
	if err := c.ShouldBindUri(&binding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// Write the URL to the database
		c.JSON(http.StatusOK, gin.H{"shortURL": binding.ShortURL})
	}
}
