package internals

import (
	"encoding/json"
	"net/http"

	internals "github.com/himanshuBamoriya/MovieApiTrainingPorj/internals/store"

	"github.com/himanshuBamoriya/MovieApiTrainingPorj/internals/models"
)

// new handler
type handler struct {
	sender internals.DataSender
}

func New(s internals.DataSender) *handler {
	return &handler{sender: s}
}

func (ds *handler) HandleErrors(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(404)
	writer.Write(ds.sender.ErrorResponse(404))
}

func (ds *handler) HandleMethodErrors(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(405)
	writer.Write(ds.sender.ErrorResponse(405))
}

//func (ds *handler) UpdateMovieById(writer http.ResponseWriter, request *http.Request) {
//	//ctx := context.Background()
//
//	body, err := io.ReadAll(request.Body)
//	if err != nil {
//		panic(err)
//	}
//	var newMovie models2.Movies
//	err = json.Unmarshal(body, &newMovie)
//	if err != nil {
//		writer.WriteHeader(400)
//		fmt.Fprintf(writer, "Bad Request")
//	}
//	movies := ds.sender.GetMovies()
//	vars := mux.Vars(request)
//	key, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		writer.WriteHeader(400)
//		writer.Write([]byte("mismatched key type for id"))
//		return
//	}
//
//	for idx, movie := range movies {
//		if movie.Id == key {
//			movies = removeData(movies, idx)
//			break
//		}
//	}
//	movies = append(movies, newMovie)
//	WriteSingleResponse(newMovie, writer)
//	models.SaveMovies(movies)
//
//}
//
//func DeleteMovieById(writer http.ResponseWriter, request *http.Request) {
//	ctx := context.Background()
//
//	vars := mux.Vars(request)
//	key, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		writer.WriteHeader(400)
//		writer.Write([]byte("mismatched key type for id"))
//		return
//	}
//	_, err = models.Db.ExecContext(ctx, "DELETE from movies where Id = ?", key)
//	if err != nil {
//		panic(err)
//	}
//	writer.WriteHeader(404)
//	writer.Write([]byte("id deleted successfully"))
//}

func removeData(movie []models.Movies, idx int) []models.Movies {
	movie[idx] = movie[len(movie)-1]
	return movie[:len(movie)-1]
}

func (ds *handler) GetAllMovies(writer http.ResponseWriter, request *http.Request) {
	movies := ds.sender.GetMovies()

	ds.WriteResponse(movies, writer)
}

//func (ds *handler) SaveMovies(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("content-type", "application/json")
//	if request.Method == "POST" {
//		body, err := io.ReadAll(request.Body)
//		var movies []models2.Movies
//		err = json.Unmarshal(body, &movies)
//		if err != nil {
//			writer.WriteHeader(400)
//			fmt.Fprintf(writer, "Bad Request")
//		}
//		ds.WriteResponse(movies, writer)
//		models.SaveMovies(movies)
//
//	} else {
//		writer.WriteHeader(405)
//		fmt.Fprintf(writer, "This request Method is not supported")
//	}
//
//}

//func (ds *handler) SaveMoviesById(writer http.ResponseWriter, request *http.Request) {
//	vars := mux.Vars(request)
//	key, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		writer.WriteHeader(400)
//		errorStatus := models.ErrorResponse(400)
//		writer.Write(errorStatus)
//		return
//	}
//	movies := models.GetMovies()
//
//	for _, movie := range movies {
//		if movie.Id == key {
//			WriteSingleResponse(movie, writer)
//			return
//		}
//	}
//	writer.WriteHeader(404)
//	errorStatus := models.ErrorResponse(404)
//	writer.Write(errorStatus)
//}

func (ds *handler) WriteResponse(movies []models.Movies, writer http.ResponseWriter) {
	jsonResp := ds.sender.GetResponse(movies)
	jsonData, err := json.MarshalIndent(jsonResp, "", "    ")
	if err != nil {
		writer.WriteHeader(400)
		errorStatus := ds.sender.ErrorResponse(400)
		writer.Write(errorStatus)
	}
	writer.Write(jsonData)
}

func (ds *handler) WriteSingleResponse(movie models.Movies, writer http.ResponseWriter) {
	jsonResp := ds.sender.GetSingleResponse(movie)
	jsonData, err := json.MarshalIndent(jsonResp, "", "    ")
	if err != nil {
		writer.WriteHeader(400)
		errorStatus := ds.sender.ErrorResponse(400)
		writer.Write(errorStatus)
	}
	writer.Write(jsonData)
}
