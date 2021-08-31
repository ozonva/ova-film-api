package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
)

// ImplementedMovieServiceServer must be embedded to have forward compatible implementations.
type ImplementedMovieServiceServer struct {
}

func (ImplementedMovieServiceServer) CreateMovie(context.Context, *Movie) (*empty.Empty, error) {
	log.Print("CreateMovie performed")
	return &empty.Empty{}, nil
}
func (ImplementedMovieServiceServer) DescribeMovie(context.Context, *DescribeMovieMessage) (*Movie, error) {
	log.Print("DescribeMovie performed")
	return &Movie{}, nil
}
func (ImplementedMovieServiceServer) RemoveMovie(context.Context, *RemoveMovieMessage) (*empty.Empty, error) {
	log.Print("RemoveMovie performed")
	return &empty.Empty{}, nil
}
func (ImplementedMovieServiceServer) ListMovies(context.Context, *MovieListRequest) (*MovieList, error) {
	log.Print("ListMovies performed")
	return &MovieList{}, nil
}
func (ImplementedMovieServiceServer) mustEmbedUnimplementedMovieServiceServer() {}
