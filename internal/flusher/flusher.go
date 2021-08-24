package flusher

import (
	"github.com/ozonva/ova_film_api/internal/movies"
	"github.com/ozonva/ova_film_api/internal/repo"
)

type Flusher interface {
	Flush(entities []movies.Movie) []movies.Movie
}

func NewFlusher(chunkSize int, entityRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func (r *flusher) Flush(movies []movies.Movie) []movies.Movie {
	err := r.entityRepo.AddEntities(movies)
	if err != nil {
		return movies
	}
	panic("Something went wrong!")
}
