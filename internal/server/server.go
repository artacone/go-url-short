package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/artacone/go-url-short/internal/api"
	"github.com/artacone/go-url-short/internal/repository"
	"github.com/artacone/go-url-short/internal/services"
)

type URLShortenerServer struct {
	shortener *services.Shortener
	pb.UnimplementedURLShortenerServer
}

func Run(mem *string) {
	repo, err := repository.New(mem)
	if err != nil {
		log.Fatalf("failed to run repo: %v", err)
	}

	shortener := services.New(repo)
	server := URLShortenerServer{shortener: shortener}

	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	reflection.Register(grpcServer)

	pb.RegisterURLShortenerServer(grpcServer, server)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to accept: %v", err)
	}
}
