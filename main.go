package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/himanshuBamoriya/MovieApiTrainingPorj/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	RequestHandler()
}

func RequestHandler() {
	moviesRouter := mux.NewRouter().StrictSlash(true)

	moviesRouter.HandleFunc("/", GetAllMovies).Methods("GET")
	moviesRouter.HandleFunc("/movies/{id}", SaveMoviesById).Methods("GET")
	moviesRouter.HandleFunc("/movies", SaveMovies).Methods("POST")
	moviesRouter.HandleFunc("/movies/{id}", DeleteMovieById).Methods("DELETE")
	moviesRouter.HandleFunc("/movies/{id}", UpdateMovieById).Methods("PUT")

	moviesRouter.NotFoundHandler = http.HandlerFunc(handleErrors)
	moviesRouter.MethodNotAllowedHandler = http.HandlerFunc(handleMethodErrors)

	log.Fatal(http.ListenAndServe(":8080", moviesRouter))
}

func handleErrors(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(404)
	writer.Write(models.ErrorResponse(404))
}
func handleMethodErrors(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(405)
	writer.Write(models.ErrorResponse(405))
}

func UpdateMovieById(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var newMovie models.Movies
	err = json.Unmarshal(body, &newMovie)
	if err != nil {
		writer.WriteHeader(400)
		fmt.Fprintf(writer, "Bad Request")
	}
	movies := models.GetMovies()
	vars := mux.Vars(request)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("mismatched key type for id"))
		return
	}

	for idx, movie := range movies {
		if movie.Id == key {
			movies = removeData(movies, idx)
			break
		}
	}
	movies = append(movies, newMovie)
	WriteSingleResponse(newMovie, writer)
	models.SaveMovies(movies)

}

func DeleteMovieById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("mismatched key type for id"))
		return
	}

	movies := models.GetMovies()
	for idx, movie := range movies {
		if movie.Id == key {
			movies = removeData(movies, idx)
			movieByte, err := json.MarshalIndent(movies, "", "    ")
			if err != nil {
				panic(err)
			}
			os.WriteFile(
				"./data/movieData.json",
				movieByte,
				0644,
			)
			jsonResp := models.MovieDeleted()
			jsonData, err := json.MarshalIndent(jsonResp, "", "\t")
			writer.Write(jsonData)
			return
		}
	}
	writer.WriteHeader(404)
	writer.Write([]byte("id does not exist"))
}

func removeData(movie []models.Movies, idx int) []models.Movies {
	movie[idx] = movie[len(movie)-1]
	return movie[:len(movie)-1]
}

func GetAllMovies(writer http.ResponseWriter, _ *http.Request) {
	movies := models.GetMovies()

	WriteResponse(movies, writer)
}

func SaveMovies(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, err := io.ReadAll(request.Body)
		var movies []models.Movies
		err = json.Unmarshal(body, &movies)
		if err != nil {
			writer.WriteHeader(400)
			fmt.Fprintf(writer, "Bad Request")
		}
		WriteResponse(movies, writer)
		models.SaveMovies(movies)

	} else {
		writer.WriteHeader(405)
		fmt.Fprintf(writer, "This request Method is not supported")
	}

}

func SaveMoviesById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		writer.WriteHeader(400)
		errorStatus := models.ErrorResponse(400)
		writer.Write(errorStatus)
		return
	}
	movies := models.GetMovies()

	for _, movie := range movies {
		if movie.Id == key {
			WriteSingleResponse(movie, writer)
			return
		}
	}
	writer.WriteHeader(404)
	errorStatus := models.ErrorResponse(404)
	writer.Write(errorStatus)
}

func WriteResponse(movies []models.Movies, writer http.ResponseWriter) {
	jsonResp := models.GetResponse(movies)
	jsonData, err := json.MarshalIndent(jsonResp, "", "    ")
	if err != nil {
		writer.WriteHeader(400)
		errorStatus := models.ErrorResponse(400)
		writer.Write(errorStatus)
	}
	writer.Write(jsonData)
}

func WriteSingleResponse(movie models.Movies, writer http.ResponseWriter) {
	jsonResp := models.GetSingleResponse(movie)
	jsonData, err := json.MarshalIndent(jsonResp, "", "    ")
	if err != nil {
		writer.WriteHeader(400)
		errorStatus := models.ErrorResponse(400)
		writer.Write(errorStatus)
	}
	writer.Write(jsonData)
}
