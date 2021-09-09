package ova_film_api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ozonva/ova_film_api/internal/movies"
	"github.com/ozonva/ova_film_api/internal/repo"
	"github.com/ozonva/ova_film_api/pkg/generated/api"
	"github.com/rs/zerolog/log"
)

func NewServer(r repo.MovieRepo) api.MovieServiceServer {
	return &implementedMovieServiceServer{
		repo: r,
	}
}

type implementedMovieServiceServer struct {
	api.UnimplementedMovieServiceServer
	repo repo.MovieRepo
}

func (s implementedMovieServiceServer) CreateMovie(c context.Context, m *api.Movie) (*empty.Empty, error) {
	err := s.repo.AddMovie(mapBack(m))
	log.Print("CreateMovie performed")
	return &empty.Empty{}, err
}
func (s implementedMovieServiceServer) DescribeMovie(c context.Context, m *api.DescribeMovieMessage) (*api.Movie, error) {
	entity, err := s.repo.DescribeMovie(m.MovieId)
	log.Printf("DescribeMovie performed %s", entity)
	return mapp(entity), err
}
func (s implementedMovieServiceServer) RemoveMovie(c context.Context, m *api.RemoveMovieMessage) (*empty.Empty, error) {
	err := s.repo.RemoveMovie(m.MovieId)
	log.Print("RemoveMovie performed")
	return &empty.Empty{}, err
}
func (s implementedMovieServiceServer) ListMovies(context.Context, *api.MovieListRequest) (*api.MovieList, error) {
	list, _ := s.repo.ListMovies(10, 0)

	responseSlice := make([]*api.Movie, 0)

	for _, movie := range list {
		responseSlice = append(responseSlice, mapp(&movie))
	}

	response := &api.MovieList{Movies: responseSlice}

	log.Printf("ListMovies performed %s", response)
	return response, nil
}

func mapp(source *movies.Movie) *api.Movie {
	return &api.Movie{
		Id:     source.Id,
		UserId: source.UserId,
		Name:   source.Name,
		Year:   source.Year,
	}
}

func mapBack(source *api.Movie) *movies.Movie {
	return movies.New(source.Id, source.UserId, source.Name, source.Year)
}
