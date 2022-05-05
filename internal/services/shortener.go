package services

import (
	"context"
	"crypto/rand"
	"strings"

	"github.com/artacone/go-url-short/internal/models"
	"github.com/artacone/go-url-short/internal/repository"
)

type Shortener struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Shortener {
	return &Shortener{repo: repo}
}

func (s *Shortener) Create(ctx context.Context, url models.URL) (models.ShortURL, error) {
	if code, err := s.repo.GetCode(url); err == nil {
		return code, nil
	}

	code, err := generateCode()
	short := models.ShortURL{Code: code}
	if err != nil {
		return short, err
	}
	if _, err := s.repo.Create(url, short); err != nil {
		return models.ShortURL{}, err
	}
	return short, nil
}

func (s *Shortener) Get(ctx context.Context, code models.ShortURL) (models.URL, error) {
	url, err := s.repo.GetURL(code)
	return url, err
}

const (
	alphabet    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
	alphabetLen = byte(len(alphabet))
)

func generateCode() (string, error) {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return encode(b), nil
}

func encode(input []byte) string {
	var sb strings.Builder
	sb.Grow(len(input))
	for _, b := range input {
		sb.WriteByte(alphabet[b%alphabetLen])
	}
	return sb.String()
}
