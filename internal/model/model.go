package model

import (
	"regexp"
	"time"
)

// URL is the model for the Url table
type URL struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	ShortURL  string `gorm:"index"`
	LongURL   string `gorm:"index"`
}

var URLMatchRegex *regexp.Regexp

func init() {
	URLMatchRegex = regexp.MustCompile(`(https:\/\/www\.|http:\/\/www\.|https:\/\/|http:\/\/)?[a-zA-Z0-9]{2,}(\.[a-zA-Z0-9]{2,})(\.[a-zA-Z0-9]{2,})?`)
}

// CompileURLRegex compiles the regex for validating url
func CompileURLRegex(expr string) *regexp.Regexp {
	regex, err := regexp.Compile(expr)
	if err != nil {
		panic(err)
	}

	return regex
}

// ValidateURL validates the url, return true if valid
func ValidateURL(url string) bool {
	return URLMatchRegex.MatchString(url)
}
