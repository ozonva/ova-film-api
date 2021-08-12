package utils

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
