package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"

	models2 "github.com/himanshuBamoriya/MovieApiTrainingPorj/internals/models"
	"github.com/himanshuBamoriya/MovieApiTrainingPorj/models"
)

func TestGetAllMovies(t *testing.T) {
	movie := models.GetMovies()
	want := models.GetResponse(movie)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	GetAllMovies(w, req)
	res := w.Result()

	//resp, err := io.ReadAll(res.Body)
	//if err != nil {
	//	t.Errorf("cannot find data")
	//}
	marshalled, _ := json.Marshal(res.Body)

	jsonresp := models2.Response{}

	err := json.Unmarshal(marshalled, &jsonresp)

	if err != nil {
		t.Errorf("no data found")
	}

	if !reflect.DeepEqual(jsonresp, want) {
		t.Errorf("Mismatch content type ----- %v \n -----%v", jsonresp, want)
	}

}

func TestSaveMovies(t *testing.T) {
	movieTestCase := []models2.Movies{
		{
			1,
			"tarzan",
			"Action",
			4.4,
			"Wonder",
			true,
		},
	}

	postBody := models2.Response{
		Code:   200,
		Status: "SUCCESS",
		Data:   movieTestCase,
	}

	body, _ := json.Marshal(movieTestCase)

	req := httptest.NewRequest("POST", "/movies", bytes.NewReader(body))
	w := httptest.NewRecorder()
	SaveMovies(w, req)
	res := w.Result()

	resp, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("cannot find data")
	}

	jsonresp := models2.Response{}

	err = json.Unmarshal(resp, &jsonresp)

	if err != nil {
		t.Errorf("no data found")
	}

	if !reflect.DeepEqual(postBody, jsonresp) {
		t.Errorf("Error")
	}

}
