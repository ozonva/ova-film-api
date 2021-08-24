package repo

import "github.com/ozonva/ova_film_api/internal/movies"

type Repo interface {
	AddEntities(entities []movies.Movie) error
	ListEntities(limit, offset uint64) ([]movies.Movie, error)
	DescribeEntity(entityId uint64) (*movies.Movie, error)
}
