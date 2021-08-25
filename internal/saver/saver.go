package saver

import (
	"github.com/ozonva/ova_film_api/internal/flusher"
	"github.com/ozonva/ova_film_api/internal/movies"
)

type Saver interface {
	Save(movie movies.Movie)
	Close()
}

func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
) Saver {
	return &saver{
		movies:  make([]movies.Movie, capacity),
		flusher: flusher,
	}
}

type saver struct {
	movies  []movies.Movie
	flusher flusher.Flusher
}

func (s saver) Save(movie movies.Movie) {
	s.movies = append(s.movies, movie)
}

func (s saver) Close() {
	s.movies = s.flusher.Flush(s.movies)
}
