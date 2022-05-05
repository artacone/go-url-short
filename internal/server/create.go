package server

import (
	"context"

	pb "github.com/artacone/go-url-short/internal/api"
	"github.com/artacone/go-url-short/internal/models"
)

func (s URLShortenerServer) Create(ctx context.Context, url *pb.URL) (*pb.ShortURL, error) {
	code, err := s.shortener.Create(ctx, models.URL{URL: url.GetURL()})
	short := &pb.ShortURL{URL: code.Code}
	return short, err
}
