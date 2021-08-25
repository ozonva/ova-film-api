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
	capacity uint,
	flusher flusher.Flusher,
	timeout time.Duration,
) Saver {
	saver := saver{
		movies:   make([]movies.Movie, capacity),
		flusher:  flusher,
		capacity: capacity,
		peek:     0,
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
	capacity uint
	peek     uint
}

func (s *saver) Save(movie movies.Movie) {
	s.movies[s.peek] = movie
	s.peek++
	if s.peek == s.capacity {
		s.Close()
	}
}

func (s *saver) Close() {
	s.movies = s.flusher.Flush(s.movies[:s.peek])
	s.peek = 0
}
