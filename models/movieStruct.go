package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movies struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Gener    string  `json:"gener,omitempty"`
	Rating   float64 `json:"rating,omitempty"`
	Plot     string  `json:"plot,omitempty"`
	Released bool    `json:"released,omitempty"`
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
	moviesByte, err := json.MarshalIndent(movies, "", "    ")

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./data/movieData.json", moviesByte, 0644)

	if err != nil {
		panic(err)
	}
}
