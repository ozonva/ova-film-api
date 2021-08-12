package main

import (
	"fmt"
	"github.com/ozonva/ova_film_api/internal/ova_film_api"
	"github.com/ozonva/ova_film_api/internal/utils"
)

func main() {
	fmt.Printf(ova_film_api.HelloWorldText())

	fmt.Println()
	fmt.Println("--Separating slice--")
	slice := []string{"aa", "bb", "cc", "dd", "e", "f", "gg", "hh"}
	var size int = 3
	var slices [][]string = utils.Split(slice, size)
	for _, slice := range slices {
		fmt.Println(slice)
	}

	fmt.Println()
	fmt.Println("--Turning map--")
	sourceMap := map[string]string{
		"first":   "Mercury",
		"second":  "Venus",
		"third":   "Earth",
		"fourth":  "Mars",
		"fifth":   "Jupiter",
		"sixth":   "Saturn",
		"seventh": "Uranus",
		"eight":   "Neptune",
	}
	turnedMap := utils.TurnMap(sourceMap)
	for key, value := range turnedMap {
		fmt.Printf("%s is %s\n", key, value)
	}

	fmt.Println()
	fmt.Println("--Filtering slice--")
	sourceSlice := []string{"one", "two", "three", "four", "three", "four", "five"}
	filteredSlice := utils.Filter(sourceSlice, "three")
	fmt.Println(filteredSlice)

}
