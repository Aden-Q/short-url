package handler

import (
	"errors"
	"net/http"

	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Used for URI binding
type bindingURI struct {
	ShortURL string `uri:"shortURL" binding:"required"`
}

// ShortenHandlerFunc shortens a long URL
func RedirectHandler(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*db.Engine)
	if !ok {
		c.JSON(500, gin.H{"message": "dbConn not found"})
		return
	}

	var binding bindingURI
	if err := c.ShouldBindUri(&binding); err != nil {
		// TODO: add a log here
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		// if there's a record in the database, redirect to the long URL
		// otherwise return 404
		var url model.URL
		if err := db.First(&url, "short_url = ?", binding.ShortURL).Error; err == nil {
			c.Redirect(http.StatusMovedPermanently, url.LongURL)
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "URL not found"})
			return
		} else {
			// other errors on the server side
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		return
	}
}
