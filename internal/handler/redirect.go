package handler

import (
	"errors"
	"net/http"

	"github.com/Aden-Q/short-url/internal/cache"
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/model"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/gin-gonic/gin"
)

// Used for URI binding
type bindingURI struct {
	ShortURL string `uri:"shortURL" binding:"required"`
}

// @Summary RedirectHandler redirects a short URL to a long URL
// @Produce json
// @Success 301 {object} model.URL "long URL"
// @Failure 400 {string} string "Invalid URL"
// @Failure 404 {string} string "URL not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/{shortURL} [get]
func RedirectHandler(c *gin.Context) {
	// make sure mysql connection is established is is stored into the context
	d, ok := c.MustGet("dbConn").(db.Engine)
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
	redisCache, ok := c.MustGet("cache").(cache.Cache)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no cache"})
		return
	}

	var binding bindingURI
	if err := c.ShouldBindUri(&binding); err != nil {
		// TODO: add a log here
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		// check if the input short URL is valid
		if !model.ValidateShortURL(binding.ShortURL) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid URL"})
			return
		}

		// check whether the request short URL is cached
		var cacheObj cacheObject
		if err := redisCache.Get(binding.ShortURL, &cacheObj); err == nil {
			// the url already exists in cache, return the short url and do nothing
			c.Redirect(http.StatusMovedPermanently, cacheObj.URL)
			return
		} else if !errors.Is(err, cache.ErrCacheMiss) {
			// other errors on the server side
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		// cache miss, continue to query the database
		// if there's a record in the database, redirect to the long URL
		// otherwise return 404
		var url model.URL
		if err := d.First(&url, "short_url = ?", binding.ShortURL); err == nil {
			// the url already exists, return the long url and cache the result
			if err := redisCache.Set(binding.ShortURL, &cacheObject{URL: url.LongURL}); err == nil {
				c.Redirect(http.StatusMovedPermanently, url.LongURL)
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		} else if errors.Is(err, db.ErrRecordNotFound) {
			// TODO: for empty lookup, maybe we can also cache the result
			// in that case we need to checked the cached result first
			c.JSON(http.StatusNotFound, gin.H{"message": "URL not found"})
			return
		} else {
			// other errors on the server side
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		return
	}
}
