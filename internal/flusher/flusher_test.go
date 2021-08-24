package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova_film_api/internal/repo"
	"github.com/ozonva/ova_film_api/mocks"
)

var _ = Describe("Flusher", func() {
	var ctrl *gomock.Controller
	var repo repo.MovieRepo

	BeforeEach(func() {
		ctrl = gomock.NewController(nil)
		repo = mocks.NewMockMovieRepo(ctrl)
	})
})
