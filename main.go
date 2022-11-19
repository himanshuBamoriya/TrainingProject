package main

import (
	"encoding/json"
	"github.com/himanshuBamoriya/MovieApiTrainingPorj/models"
	"net/http"
)

func main() {
	http.HandleFunc("/", HttpHandler)
	http.ListenAndServe(":8080", nil)
}

func HttpHandler(writer http.ResponseWriter, request *http.Request) {
	movies := models.GetMovies()
	movieByte, err := json.Marshal(movies)

	if err != nil {
		panic(err)
	}

	writer.Write(movieByte)
}
