package repository

import "github.com/artacone/go-url-short/internal/models"

type Repository interface {
	Create(url models.URL, code models.ShortURL) (models.ShortURL, error)
	GetURL(code models.ShortURL) (models.URL, error)
	GetCode(url models.URL) (models.ShortURL, error)
}

func New() Repository {
	return newCache()
}
