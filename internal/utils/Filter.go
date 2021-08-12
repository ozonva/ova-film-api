package utils

func Filter(where []string, what string) []string {
	var result []string
	for _, value := range where {
		if value != what {
			result = append(result, value)
		}
	}
	return result
}
