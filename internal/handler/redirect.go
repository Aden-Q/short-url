package handler

import (
	"errors"
	"net/http"

	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/model"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/gin-gonic/gin"
)

// Used for URI binding
type bindingURI struct {
	ShortURL string `uri:"shortURL" binding:"required"`
}

// ShortenHandlerFunc shortens a long URL
func RedirectHandler(c *gin.Context) {
	// make sure mysql connection is established is is stored into the context
	d, ok := c.MustGet("dbConn").(db.Engine)
	if !ok {
		c.JSON(500, gin.H{"message": "dbConn not found"})
		return
	}

	// make sure redis connection is established is is stored into the context
	_, ok = c.MustGet("redisConn").(redis.Client)
	if !ok {
		c.JSON(500, gin.H{"message": "redis connection not established"})
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
		if err := d.First(&url, "short_url = ?", binding.ShortURL); err == nil {
			c.Redirect(http.StatusMovedPermanently, url.LongURL)
			return
		} else if errors.Is(err, db.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "URL not found"})
			return
		} else {
			// other errors on the server side
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		return
	}
}
