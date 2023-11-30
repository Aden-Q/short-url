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
<<<<<<< HEAD
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
<<<<<<< HEAD
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
=======
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
>>>>>>> 0a4130b (feat: add a redirect handler for the GET REST endpoint)
	}
}
