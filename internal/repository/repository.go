package repository

import (
	"context"
	"errors"

	"github.com/artacone/go-url-short/internal/models"
)

var ErrNotFound = errors.New("not found")

type Repository interface {
	Create(ctx context.Context, url models.URL, code models.ShortURL) (models.ShortURL, error)
	GetURL(ctx context.Context, code models.ShortURL) (models.URL, error)
	GetCode(ctx context.Context, url models.URL) (models.ShortURL, error)
}

func New(dsn string) (Repository, error) {
	switch dsn {
	case "":
		return newCache()
	default:
		return newDB(dsn)
	}
}
