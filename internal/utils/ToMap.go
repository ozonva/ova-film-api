package utils

import "github.com/ozonva/ova_film_api/internal/movies"

func ToMap(entities []movies.Movie) (map[uint64]movies.Movie, error) {
	moviesMap := map[uint64]movies.Movie{}
	for _, entity := range entities {
		moviesMap[entity.Id] = entity
	}

	return moviesMap, nil
}
