package handler

type cacheObject struct {
	URL string `json:"url"`
}

// cacheCallback is the function called when cache miss happens. It returns to indicate there's a cache miss
// func cacheCallback(url string) func() (interface{}, error) {
// 	return func() (interface{}, error) {
// 		return &cacheObject{url: url}, cache.ErrCacheMiss
// 	}
// }
