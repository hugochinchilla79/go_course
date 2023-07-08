package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"gopkg.in/mgo.v2"
)

var movies = GetMovieList()

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://root:rootpassword@127.0.0.1:27017/go_course")
	if err != nil {
		panic(err)	
	}
	return session
}	

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To get in touch, please send an email to")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}	

func GetMovieList() Movies {
	movies := Movies{
		Movie{Id: 1, Name: "The Matrix", Year: 1999, Director: "Wachowski Brothers"},
		Movie{Id: 2, Name: "The Matrix Reloaded", Year: 2003, Director: "Wachowski Brothers"},
		Movie{Id: 3, Name: "The Matrix Revolutions", Year: 2003, Director: "Wachowski Brothers"},
	}
	return movies
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id, _ := strconv.Atoi(params["id"]) // Convert string to int

	for _, item := range movies {
		if item.Id == movie_id {
			json.NewEncoder(w).Encode(item)
		}	
	}
	json.NewEncoder(w).Encode(&Movie{})
}


func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if(err != nil) {
		panic(err)
	}

	defer r.Body.Close()
	
	session := getSession()
	session.DB("go_course").C("movies").Insert(movie_data)
	session.Close()
	json.NewEncoder(w).Encode(movies)
}