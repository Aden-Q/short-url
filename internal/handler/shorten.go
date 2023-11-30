package handler

import (
	"errors"
	"net/http"

	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jxskiss/base62"
	"gorm.io/gorm"
)

// Used for query parameter binding
type bindingQuery struct {
	LongURL string `form:"longURL" binding:"required"`
}

// ShortenHandler shortens a long URL
func ShortenHandler(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*db.Engine)
	if !ok {
		c.JSON(500, gin.H{"message": "dbConn not found"})
		return
	}

	var binding bindingQuery
	if err := c.ShouldBindQuery(&binding); err != nil {
		// TODO: add a log here
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		// First check if the input long URL already exists in the database
		var url model.URL
		if err := db.First(&url, "long_url = ?", binding.LongURL).Error; err == nil {
			// the url already exists, return the short url and do nothing
			c.JSON(http.StatusOK, gin.H{"shortURL": url.ShortURL})
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			// the url does not exist, create a new short url and write it into the db
			url = model.URL{
				LongURL: binding.LongURL,
			}

			// insert the record into the database
			if err := db.Create(&url).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			} else {
				// encode the auto-incremented ID into a short URL with base62
				shortURL := string(base62.FormatUint(url.ID))
				// update the short URL in the database
				if err := db.Model(&url).Update("short_url", shortURL).Error; err != nil {
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
