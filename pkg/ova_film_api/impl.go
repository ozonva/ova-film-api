package ova_film_api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ozonva/ova_film_api/pkg/generated/api"
	"github.com/rs/zerolog/log"
)

// ImplementedMovieServiceServer must be embedded to have forward compatible implementations.
type ImplementedMovieServiceServer struct {
	api.MovieServiceServer
}

func (ImplementedMovieServiceServer) CreateMovie(context.Context, *api.Movie) (*empty.Empty, error) {
	log.Print("CreateMovie performed")
	return &empty.Empty{}, nil
}
func (ImplementedMovieServiceServer) DescribeMovie(context.Context, *api.DescribeMovieMessage) (*api.Movie, error) {
	log.Print("DescribeMovie performed")
	return &api.Movie{}, nil
}
func (ImplementedMovieServiceServer) RemoveMovie(context.Context, *api.RemoveMovieMessage) (*empty.Empty, error) {
	log.Print("RemoveMovie performed")
	return &empty.Empty{}, nil
}
func (ImplementedMovieServiceServer) ListMovies(context.Context, *api.MovieListRequest) (*api.MovieList, error) {
	log.Print("ListMovies performed")
	return &api.MovieList{}, nil
}
func (ImplementedMovieServiceServer) mustEmbedUnimplementedMovieServiceServer() {}
