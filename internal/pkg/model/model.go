package model

// URL is the model for the Url table
type URL struct {
	ID       uint32 `json:"id"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}
