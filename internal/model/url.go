package model

import (
	"regexp"
)

// URL is the model for the Url table
type URL struct {
	Model
	ShortURL string `gorm:"index"`
	LongURL  string `gorm:"index"`
}

var (
	// short URL should be a legal base62 string
	shortURLMatchRegex *regexp.Regexp
	// long URL should be a legal http/https url
	longURLMatchRegex *regexp.Regexp
)

func init() {
	shortURLMatchRegex = regexp.MustCompile(`^[A-Za-z0-9]+$`)
	longURLMatchRegex = regexp.MustCompile(`(https:\/\/www\.|http:\/\/www\.|https:\/\/|http:\/\/)?[a-zA-Z0-9]{2,}(\.[a-zA-Z0-9]{2,})(\.[a-zA-Z0-9]{2,})?`)
}

// CompileURLRegex compiles the regex for validating url
func CompileURLRegex(expr string) *regexp.Regexp {
	regex, err := regexp.Compile(expr)
	if err != nil {
		panic(err)
	}

	return regex
}

// ValidateShortURL validates a short url, return true if valid
// shortURL should be a legal base62 string
func ValidateShortURL(shortURL string) bool {
	return shortURLMatchRegex.MatchString(shortURL)
}

// ValidateLongURL validates a long url, return true if valid
// longURL should be a legal http/https url
func ValidateLongURL(longURL string) bool {
	return longURLMatchRegex.MatchString(longURL)
}
