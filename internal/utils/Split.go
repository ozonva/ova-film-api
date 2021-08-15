package utils

import "github.com/ozonva/ova_film_api/internal/movies"

func Split(slice []string, size int) [][]string {
	var result [][]string
	for i := 0; i < len(slice); i += size {
		result = append(result, slice[i:min(i+size, len(slice))])
	}
	return result
}

func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func SplitToBulks(entities []movies.Movie, butchSize int) [][]movies.Movie {
	var result [][]movies.Movie
	for i := 0; i < len(entities); i += butchSize {
		result = append(result, entities[i:min(i+butchSize, len(entities))])
	}
	return result
}
