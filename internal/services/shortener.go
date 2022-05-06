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

// Create checks if the url is already present in repository and if
// it is, then it returns the corresponding code. Otherwise,
// it generates a code and stores a pair (url, code) in repository.
func (s *Shortener) Create(ctx context.Context, url models.URL) (models.ShortURL, error) {
	if code, err := s.repo.GetCode(ctx, url); err == nil {
		return code, nil
	}

	code, err := generateCode()
	short := models.ShortURL{Code: code}
	if err != nil {
		return short, err
	}
	if _, err := s.repo.Create(ctx, url, short); err != nil {
		return models.ShortURL{}, err
	}
	return short, nil
}

// Get returns an url corresponding to the code. If the code is not present
// in the repository, an error is returned.
func (s *Shortener) Get(ctx context.Context, code models.ShortURL) (models.URL, error) {
	url, err := s.repo.GetURL(ctx, code)
	return url, err
}

// generateCode returns a string of length 10 which consists of
// random alphanumeric (and '_') characters.
func generateCode() (string, error) {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return encode(b), nil
}

const (
	alphabet    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
	alphabetLen = byte(len(alphabet))
)

func encode(input []byte) string {
	var sb strings.Builder
	sb.Grow(len(input))
	for _, b := range input {
		sb.WriteByte(alphabet[b%alphabetLen])
	}
	return sb.String()
}
