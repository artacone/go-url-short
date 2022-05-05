package server

import (
	"context"

	pb "github.com/artacone/go-url-short/internal/api"
	"github.com/artacone/go-url-short/internal/models"
)

func (s URLShortenerServer) Get(ctx context.Context, short *pb.ShortURL) (*pb.URL, error) {
	url, err := s.shortener.Get(ctx, models.ShortURL{Code: short.GetURL()})
	long := &pb.URL{URL: url.URL}
	return long, err
}
