package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response struct {
	Code   int      `json:"code,omitempty"`
	Status string   `json:"status,omitempty"`
	Data   []Movies `json:"data"`
}

type SingleResponse struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Data   Movies `json:"data"`
}

type DeleteStatus struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Data   string `json:"data"`
}

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

func ErrorResponse(code int) (resp []byte) {
	response := Response{
		Code:   code,
		Status: "ERROR",
	}

	responseByte, err := json.MarshalIndent(response, "", "    ")

	if err != nil {
		panic(err)
	}

	return responseByte
}

func GetResponse(movie []Movies) (resp Response) {
	var response = Response{
		Code:   200,
		Status: "SUCCESS",
		Data:   movie,
	}
	return response
}

func GetSingleResponse(movie Movies) (resp SingleResponse) {
	var response = SingleResponse{
		Code:   200,
		Status: "SUCCESS",
		Data:   movie,
	}
	return response
}

func MovieDeleted() (resp DeleteStatus) {
	var response = DeleteStatus{
		Code:   200,
		Status: "SUCCESS",
		Data:   "Movie Deleted Successfully",
	}
	return response
}
