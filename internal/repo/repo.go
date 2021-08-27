package repo

import "github.com/ozonva/ova_film_api/internal/movies"

type MovieRepo interface {
	AddMovies(entities []movies.Movie) error
	ListMovies(limit, offset uint64) ([]movies.Movie, error)
	DescribeMovie(entityId uint64) (*movies.Movie, error)
}
