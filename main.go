package main

import (
	"database/sql"
	"log"
	"net/http"

	internal "github.com/himanshuBamoriya/MovieApiTrainingPorj/internals/handlers"
	internals2 "github.com/himanshuBamoriya/MovieApiTrainingPorj/internals/store"

	"github.com/gorilla/mux"
)

var Db, err = sql.Open(
	"mysql",
	"user:password@tcp(localhost:3308)/mysql",
)

func main() {
	RequestHandler()
}

func RequestHandler() {
	moviesRouter := mux.NewRouter().StrictSlash(true)

	//db, err := sql.Open("mysql", "user:password@tcp(localhost:3308)/")
	//
	//log.Println(err, db.Ping())
	sender := internals2.Store{Db}

	datasender := internal.New(&sender)
	moviesRouter.HandleFunc("/", datasender.GetAllMovies).Methods("GET")
	//moviesRouter.HandleFunc("/movies/{id}", handlers.SaveMoviesById).Methods("GET")
	//moviesRouter.HandleFunc("/movies", handlers.SaveMovies).Methods("POST")
	//moviesRouter.HandleFunc("/movies/{id}", handlers.DeleteMovieById).Methods("DELETE")
	//moviesRouter.HandleFunc("/movies/{id}", handlers.UpdateMovieById).Methods("PUT")

	moviesRouter.NotFoundHandler = http.HandlerFunc(datasender.HandleErrors)
	moviesRouter.MethodNotAllowedHandler = http.HandlerFunc(datasender.HandleMethodErrors)

	log.Fatal(http.ListenAndServe(":8080", moviesRouter))
}
