package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movies struct {
	Id       int
	Name     string
	Gener    string
	Rating   float64
	Plot     string
	released bool
}

func GetMovies() (movies []Movies) {
	fileByte, err := os.ReadFile("./data/movieData.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(fileByte, &movies)

	if err != nil {
		fmt.Println(err)
	}

	return movies
}

func SaveMovies(movies []Movies) {
	moviesByte, err := json.Marshal(movies)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./data/update-movieData.json", moviesByte, 0644)

	if err != nil {
		panic(err)
	}
}
