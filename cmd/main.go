package main

import (
	"fmt"
	"github.com/ozonva/ova_film_api/internal/movies"
	"github.com/ozonva/ova_film_api/internal/utils"
	"time"
)

func main() {
	fmt.Printf(utils.HelloWorldText())

	{
		fmt.Println()
		fmt.Println("--Separating slice--")
		slice := []string{"aa", "bb", "cc", "dd", "e", "f", "gg", "hh"}
		var size int = 3
		var slices [][]string = utils.Split(slice, size)
		for _, slice := range slices {
			fmt.Println(slice)
		}
	}

	{
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
	}

	{
		fmt.Println()
		fmt.Println("--Filtering slice--")
		sourceSlice := []string{"one", "two", "three", "four", "three", "four", "five"}
		filteredSlice := utils.Filter(sourceSlice, "three")
		fmt.Println(filteredSlice)
	}

	{
		fmt.Println()
		fmt.Println("--Opening file 5 times with closing using defer key word--")
		utils.ReadFile()
	}

	{
		fmt.Println()
		fmt.Println("--Movie struct--")
		movie1 := movies.New(1, 25, "Titanic", 1997)
		fmt.Println(movie1.String())
	}

	{
		fmt.Println()
		fmt.Println("--SplitToBulks function for Movie entities--")
		entities := []movies.Movie{
			*movies.New(1, 25, "Titanic", 1997),
			*movies.New(2, 25, "Green Mile", 1999),
			*movies.New(3, 25, "Forrest Gump", 1999),
			*movies.New(3, 31, "WALL-E", 2008),
			*movies.New(3, 31, "Back to the Future", 1985),
			*movies.New(3, 31, "Matrix, The", 1999),
			*movies.New(3, 36, "Fight Club", 1999),
			*movies.New(3, 36, "Fight Club", 1999),
			*movies.New(3, 15, "Dirty Dancing", 1987),
		}
		var slices [][]movies.Movie = utils.SplitToBulks(entities, 2)
		for _, slice := range slices {
			fmt.Println(slice)
		}
	}

	{
		fmt.Println()
		fmt.Println("--Timer--")
		timer := time.NewTimer(2 * time.Second)
		fmt.Println("timer created")
		fmt.Println(<-timer.C)
	}

}
