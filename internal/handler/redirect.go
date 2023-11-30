package handler

import (
	"net/http"

	"github.com/Aden-Q/short-url/internal/db"
	"github.com/gin-gonic/gin"
)

// Used for URI binding
<<<<<<< HEAD
type bindingURI struct {
	ShortURL string `uri:"shortURL" binding:"required"`
}

// ShortenHandlerFunc shortens a long URL
func RedirectHandler(c *gin.Context) {
	_, ok := c.MustGet("dbConn").(*db.Engine)
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
=======
type shortURL struct {
	shortURL string `uri:"shortURL" binding:"required"`
}

// ShortenHandlerFunc shortens a long URL
func RedirectHandler(db *db.DBEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var shortURL shortURL
		if err := c.ShouldBindQuery(&shortURL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			// Write the URL to the database
			c.JSON(200, gin.H{"longURL": shortURL.shortURL})
		}
>>>>>>> f6ad432 (feat: add a handler pkg)
	}
}
