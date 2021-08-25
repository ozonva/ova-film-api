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
		Context("When 4 movies are added", func() {
			It("should flush all movies", func() {

				movieRepo.
					EXPECT().
					AddMovies(gomock.Any()).
					Times(1)

				flusher := flusher.NewFlusher(5, movieRepo)
				saver := saver.NewSaver(10, flusher, 5*time.Second)

				saver.Save(createMovie(1))
				saver.Save(createMovie(2))
				saver.Save(createMovie(3))
				saver.Save(createMovie(4))
				saver.Close()

			})
		})
	})
})

func createMovie(id uint64) movies.Movie {
	return *movies.New(id, 5, "Movie", 2000)
}
