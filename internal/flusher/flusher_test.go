package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova_film_api/internal/flusher"
	"github.com/ozonva/ova_film_api/internal/movies"
	"github.com/ozonva/ova_film_api/mocks"
	"testing"
)

func TestFlusher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flusher")
}

var _ = Describe("Flusher", func() {
	var ctrl *gomock.Controller
	var repo *mocks.MockMovieRepo

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		repo = mocks.NewMockMovieRepo(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Flush method", func() {
		Context("When input slice is empty", func() {
			It("shouldn't perform AddMovies method at all", func() {
				flusher1 := flusher.NewFlusher(5, repo)

				repo.
					EXPECT().
					AddMovies(gomock.Any()).
					Times(0)
				Expect(flusher1.Flush([]movies.Movie{})).To(Equal([]movies.Movie(nil)))
			})
		})

		Context("When input slice is chunk size * n", func() {
			It("should perform AddMovies n times", func() {
				flusher1 := flusher.NewFlusher(5, repo)

				repo.
					EXPECT().
					AddMovies(gomock.Any()).
					Times(3)
				Expect(flusher1.Flush(createMovies(15))).To(Equal([]movies.Movie(nil)))
			})
		})

		Context("When input slice is little bit less than chunk size * n", func() {
			It("should perform AddMovies n times", func() {
				flusher1 := flusher.NewFlusher(5, repo)

				repo.
					EXPECT().
					AddMovies(gomock.Any()).
					Times(3)
				Expect(flusher1.Flush(createMovies(11))).To(Equal([]movies.Movie(nil)))
			})
		})

		Context("When input slice equals chunk size", func() {
			It("should perform AddMovies 1 time", func() {
				flusher1 := flusher.NewFlusher(15, repo)

				repo.
					EXPECT().
					AddMovies(gomock.Any()).
					Times(1)
				Expect(flusher1.Flush(createMovies(15))).To(Equal([]movies.Movie(nil)))
			})
		})

		Context("When input slice is less than chunk size but more than zero", func() {
			It("should perform AddMovies 1 time", func() {
				flusher1 := flusher.NewFlusher(15, repo)

				repo.
					EXPECT().
					AddMovies(gomock.Any()).
					Times(1)
				Expect(flusher1.Flush(createMovies(10))).To(Equal([]movies.Movie(nil)))
			})
		})
	})
})

func createMovie(id uint64) movies.Movie {
	return *movies.New(id, 5, "Movie", 2000)
}

func createMovies(count int) []movies.Movie {
	movieSlice := make([]movies.Movie, count)
	for i := 0; i < count; i++ {
		movieSlice[i] = createMovie(uint64(i))
	}
	return movieSlice
}
