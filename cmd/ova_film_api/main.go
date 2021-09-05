package main

import (
	"github.com/ozonva/ova_film_api/pkg/generated/api"
	"github.com/ozonva/ova_film_api/pkg/ova_film_api"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterMovieServiceServer(s, ova_film_api.ImplementedMovieServiceServer{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
