package utils

func TurnMap(source map[string]string) map[string]string {
	result := map[string]string{}
	for key, value := range source {
		result[value] = key
	}
	return result
}
