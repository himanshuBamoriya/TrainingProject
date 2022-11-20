package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/himanshuBamoriya/MovieApiTrainingPorj/models"
	"io"
	"net/http"
)

func main() {
	RequestHandler()
}

func RequestHandler() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/", HttpGetMoviesHandler).Methods("GET")
	muxRouter.HandleFunc("/Movies", HttpSaveMoviesHandler).Methods("POST")
	muxRouter.HandleFunc("/Movies", HttpSaveMoviesHandler2).Methods("GET")
	http.ListenAndServe(":8080", muxRouter)
}

func HttpGetMoviesHandler(writer http.ResponseWriter, request *http.Request) {
	movies := models.GetMovies()

	movieByte, err := json.Marshal(movies)

	if err != nil {
		panic(err)
	}

	writer.Write(movieByte)
}

func HttpSaveMoviesHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		var movies []models.Movies
		err = json.Unmarshal(body, &movies)
		if err != nil {
			writer.WriteHeader(400)
			fmt.Fprintf(writer, "Bad Request")
		}
		models.SaveMovies(movies)

	} else {
		writer.WriteHeader(405)
		fmt.Fprintf(writer, "This request Method is not supported")
	}

}

func HttpSaveMoviesHandler2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "testGet endPoint")
}
