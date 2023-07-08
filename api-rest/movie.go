package main

type Movie struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Year int `json:"year"`
	Director string `json:"director"`
}

type Movies []Movie
