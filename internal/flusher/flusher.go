package flusher

import (
	"github.com/ozonva/ova_film_api/internal/movies"
	"github.com/ozonva/ova_film_api/internal/repo"
	"github.com/ozonva/ova_film_api/internal/utils"
	"log"
)

type Flusher interface {
	Flush(entities []movies.Movie) []movies.Movie
}

func NewFlusher(chunkSize int, entityRepo repo.MovieRepo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		movieRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize int
	movieRepo repo.MovieRepo
}

func (r *flusher) Flush(source []movies.Movie) []movies.Movie {
	var bulks [][]movies.Movie = utils.SplitToBulks(source, r.chunkSize)
	unflushed := make([]movies.Movie, 0)
	for _, bulk := range bulks {
		if err := r.movieRepo.AddMovies(bulk); err != nil {
			unflushed = append(unflushed, bulk...)
			log.Printf("Unable to flush chunk of movies to a repository due to an error: %v", err)
		}
	}
	if len(unflushed) == 0 {
		return nil
	}
	return unflushed
}
