package repository

import (
	"context"
	"sync"

	"github.com/artacone/go-url-short/internal/models"
)

// cache is a simple in-memory repository.
type cache struct {
	urls  map[models.ShortURL]models.URL
	codes map[models.URL]models.ShortURL
	mu    sync.RWMutex
}

func newCache() (*cache, error) {
	urls := make(map[models.ShortURL]models.URL)
	codes := make(map[models.URL]models.ShortURL)
	return &cache{urls: urls, codes: codes}, nil
}

// Create updates url and code maps with newly received values.
func (c *cache) Create(_ context.Context, url models.URL, code models.ShortURL) (models.ShortURL, error) {
	c.mu.Lock()
	c.urls[code] = url
	c.codes[url] = code
	c.mu.Unlock()
	return code, nil
}

// GetURL tries to get an url corresponding to the code stored in memory.
// It returns an error if such code is not present.
func (c *cache) GetURL(_ context.Context, code models.ShortURL) (models.URL, error) {
	c.mu.RLock()
	url, ok := c.urls[code]
	c.mu.RUnlock()
	if ok {
		return url, nil
	}
	return url, ErrNotFound
}

// GetCode tries to get a code corresponding to the url stored in memory.
// It returns an error if such url is not present.
func (c *cache) GetCode(_ context.Context, url models.URL) (models.ShortURL, error) {
	c.mu.RLock()
	code, ok := c.codes[url]
	c.mu.RUnlock()
	if ok {
		return code, nil
	}
	return code, ErrNotFound
}
