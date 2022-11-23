package models

import (
	_ "github.com/go-sql-driver/mysql"
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
	Genre    string  `json:"genre,omitempty"`
	Rating   float64 `json:"rating,omitempty"`
	Plot     string  `json:"plot,omitempty"`
	Released bool    `json:"released,omitempty"`
}
