package main

import "github.com/himanshuBamoriya/MovieApiTrainingPorj/models"

func main() {
	a := models.GetMovies()
	a = append(a, models.Movies{
		Id:     1,
		Name:   "luck",
		Gener:  "drama",
		Rating: 4,
		Plot:   "abc",
	})
	models.SaveMovies(a)
}
