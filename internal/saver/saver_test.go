package saver_test

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonva/ova_film_api/internal/flusher"
	"github.com/ozonva/ova_film_api/internal/movies"
	"github.com/ozonva/ova_film_api/internal/saver"
	"github.com/ozonva/ova_film_api/mocks"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSaver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Saver Suite")
}

var _ = Describe("Saver", func() {
	var ctrl *gomock.Controller
	var movieRepo *mocks.MockMovieRepo

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		movieRepo = mocks.NewMockMovieRepo(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Close method", func() {
		Context("When 3 movies are added", func() {
			It("should flush all movies 1 time", func() {

				movieRepo.
					EXPECT().
					AddMovies(gomock.Eq([]movies.Movie{createMovie(1), createMovie(2), createMovie(3)})).
					Times(1)

				flusher := flusher.NewFlusher(5, movieRepo)
				saver := saver.NewSaver(10, flusher, 5*time.Second)

				saver.Save(createMovie(1))
				saver.Save(createMovie(2))
				saver.Save(createMovie(3))
				saver.Close()

			})
		})
		Context("When movies are added several times with delays", func() {
			It("should flush all movies 4 times", func() {

				movieRepo.
					EXPECT().
					AddMovies(gomock.Any()).
					Times(4)

				flusher := flusher.NewFlusher(5, movieRepo)
				saver := saver.NewSaver(10, flusher, 1*time.Second)

				timer := time.NewTimer(500 * time.Millisecond)
				<-timer.C

				for i := 0; i < 4; i++ {
					var m uint64 = 0
					m++
					saver.Save(createMovie(m))
					timer := time.NewTimer(1 * time.Second)
					<-timer.C
				}

			})
		})
	})
})

func createMovie(id uint64) movies.Movie {
	return *movies.New(id, 5, "Movie", 2000)
}
