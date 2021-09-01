package main

import (
	"github.com/ozonva/ova_film_api/pkg/service/pkg/service"
	//"github.com/ozonva/ova_film_api/pkg/service/github.com/ozonva/ova_film_api/pkg/service"
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
	service.RegisterMovieServiceServer(s, service.ImplementedMovieServiceServer{})

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
