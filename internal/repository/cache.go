package repository

import (
	"fmt"
	"sync"

	"github.com/artacone/go-url-short/internal/models"
)

type cache struct {
	urls  map[models.ShortURL]models.URL
	codes map[models.URL]models.ShortURL
	mu    sync.RWMutex
}

func newCache() *cache {
	urls := make(map[models.ShortURL]models.URL)
	codes := make(map[models.URL]models.ShortURL)
	return &cache{urls: urls, codes: codes}
}

func (c *cache) Create(url models.URL, code models.ShortURL) (models.ShortURL, error) {
	c.mu.Lock()
	c.urls[code] = url
	c.codes[url] = code
	c.mu.Unlock()
	return code, nil
}

func (c *cache) GetURL(code models.ShortURL) (models.URL, error) {
	c.mu.RLock()
	url, ok := c.urls[code]
	c.mu.RUnlock()
	if ok {
		return url, nil
	}
	return url, fmt.Errorf("%q not found", code)
}

func (c *cache) GetCode(url models.URL) (models.ShortURL, error) {
	c.mu.RLock()
	code, ok := c.codes[url]
	c.mu.RUnlock()
	if ok {
		return code, nil
	}
	return code, fmt.Errorf("%q not found", url)
}
