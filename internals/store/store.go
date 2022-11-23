package internals

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/himanshuBamoriya/MovieApiTrainingPorj/internals/models"
)

type Store struct {
	Db *sql.DB
}

type DataSender interface {
	GetMovies() (movies []models.Movies)
	SaveMovies(movies []models.Movies)
	GetResponse(movie []models.Movies) (resp models.Response)
	GetSingleResponse(movie models.Movies) (resp models.SingleResponse)
	MovieDeleted() (resp models.DeleteStatus)
	ErrorResponse(code int) (resp []byte)
}

func (s *Store) GetMovies() (movies []models.Movies) {
	// user:password@tcp(localhost:5555)/dbname

	ctx := context.Background()

	rows, err2 := s.Db.QueryContext(ctx, "SELECT * FROM movies")
	if err2 != nil {
		panic(err2)
	}
	movieData := make([]models.Movies, 0)
	movie := models.Movies{}
	for rows.Next() {

		err := rows.Scan(
			&movie.Id,
			&movie.Name,
			&movie.Genre,
			&movie.Rating,
			&movie.Plot,
			&movie.Released,
		)
		if err != nil {
			log.Fatalln(err)
		}
		movieData = append(movieData, movie)
	}
	//log.Println(Db.Ping())
	return movieData
}

func (s *Store) SaveMovies(movies []models.Movies) {
	ctx := context.Background()
	for i := range movies {
		row, err := s.Db.ExecContext(
			ctx,
			"INSERT INTO movies VALUES(?,?,?,?,?,?)",
			movies[i].Id,
			movies[i].Name,
			movies[i].Genre,
			movies[i].Rating,
			movies[i].Plot,
			movies[i].Released,
		)
		if err != nil {
			panic(err)
		}

		log.Println(row)
	}

}

func (s *Store) ErrorResponse(code int) (resp []byte) {
	response := models.Response{
		Code:   code,
		Status: "ERROR",
	}

	responseByte, err := json.MarshalIndent(response, "", "    ")

	if err != nil {
		panic(err)
	}

	return responseByte
}

func (s *Store) GetResponse(movie []models.Movies) (resp models.Response) {
	var response = models.Response{
		Code:   200,
		Status: "SUCCESS",
		Data:   movie,
	}
	return response
}

func (s *Store) GetSingleResponse(movie models.Movies) (resp models.SingleResponse) {
	var response = models.SingleResponse{
		Code:   200,
		Status: "SUCCESS",
		Data:   movie,
	}
	return response
}

func (s *Store) MovieDeleted() (resp models.DeleteStatus) {
	var response = models.DeleteStatus{
		Code:   200,
		Status: "SUCCESS",
		Data:   "Movie Deleted Successfully",
	}
	return response
}
