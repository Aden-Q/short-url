package handler

import (
	"errors"
	"net/http"

	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/model"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/gin-gonic/gin"
	"github.com/jxskiss/base62"
)

// Used for query parameter binding
type bindingQuery struct {
	LongURL string `form:"longURL" binding:"required"`
}

// ShortenHandler shortens a long URL
func ShortenHandler(c *gin.Context) {
	// make sure mysql connection is established is is stored into the context
	dbClient, ok := c.MustGet("dbConn").(db.Engine)
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

	var binding bindingQuery
	if err := c.ShouldBindQuery(&binding); err != nil {
		// TODO: add a log here
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		// check if the input long URL is valid
		if !model.ValidateURL(binding.LongURL) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid URL"})
			return
		}

		// First check if the input long URL already exists in the database
		var url model.URL
		if err := dbClient.First(&url, "long_url = ?", binding.LongURL); err == nil {
			// the url already exists, return the short url and do nothing
			c.JSON(http.StatusOK, gin.H{"shortURL": url.ShortURL})
		} else if errors.Is(err, db.ErrRecordNotFound) {
			// the url does not exist, create a new short url and write it into the db
			url = model.URL{
				LongURL: binding.LongURL,
			}

			// insert the record into the database
			if err := dbClient.Create(&url); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			} else {
				// encode the auto-incremented ID into a short URL with base62
				shortURL := string(base62.FormatUint(url.ID))
				// update the short URL in the database
				if err := dbClient.Update(&url, "short_url", shortURL).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"shortURL": url.ShortURL})
			}
		} else {
			// other errors on the server side
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		return
	}
}
