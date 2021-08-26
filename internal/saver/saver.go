package saver

import (
	"github.com/ozonva/ova_film_api/internal/flusher"
	"github.com/ozonva/ova_film_api/internal/movies"
	"time"
)

type Saver interface {
	Save(movie movies.Movie)
	Close()
}

func NewSaver(
	capacity int,
	flusher flusher.Flusher,
	timeout time.Duration,
) Saver {
	saver := saver{
		movies:   make([]movies.Movie, 0),
		flusher:  flusher,
		capacity: capacity,
	}

	go func() {
		ticker := time.NewTicker(timeout)
		for {
			<-ticker.C
			saver.Close()
		}
	}()

	return &saver
}

type saver struct {
	movies   []movies.Movie
	flusher  flusher.Flusher
	capacity int
}

func (s *saver) Save(movie movies.Movie) {
	s.movies = append(s.movies, movie)
	if len(s.movies) == s.capacity {
		s.Close()
	}
}

func (s *saver) Close() {
	s.movies = s.flusher.Flush(s.movies)
}
