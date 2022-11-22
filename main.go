package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Model for movies
type MoviesData struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   *Data  `json:"data"`
	Error  string `json:"error""`
}

type Data struct {
	Movie *Movies `json:"movie"`
}
type Movies struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Genre    string  `json:"genre"`
	Rating   float64 `json:"rating"`
	Plot     string  `json:"plot"`
	Released bool    `json:"released"`
}

// fake database
var movies = []Movies{{Id: 1, Name: "Silicon Valley", Genre: "Comedy", Rating: 4.5, Plot: "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.", Released: true}}

// to check if id is not empty
func isEmpty(m Movies) bool {
	return m.Id == 0 && m.Name == ""
}

func serveMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hiii to the api world</h1>"))
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all the movies")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getOneMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one movie")
	w.Header().Set("Content-Type", "application/json")

	//id from request
	params := mux.Vars(r)
	fmt.Println(params)

	// looping through movies, find matching id and return the response
	for _, movie := range movies {
		movieId, err := strconv.Atoi(params["id"])
		if err != nil {
			errors.New("strconv failure")
		}
		if movie.Id == movieId {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	// if not found anything
	json.NewEncoder(w).Encode(MoviesData{Error: "No movie found with the id"})
	//json.NewEncoder(w).Encode("No movie found with the id")
	return
}

func createOneMovieData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one movie")
	w.Header().Set("Content-Type", "application/json")

	//body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Plesae enter atlease one movie data")
	}

	// if response is {}
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	if isEmpty(movie) {
		json.NewEncoder(w).Encode("no data inside")
		return
	}

	// generate random id
	//rand.Seed(time.Now().UnixNano())

	//movie.Id = 1
	movies = append(movies, movie)

	var responseOutput MoviesData
	if movie.Id == 2 {

		responseOutput = MoviesData{
			Code:   http.StatusForbidden,
			Status: "FAILURE",
			Data:   nil,
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&responseOutput)
		return
	}

	responseOutput = MoviesData{
		Code:   200,
		Status: "SUCCESS",
		Data:   &Data{Movie: &movie},
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&responseOutput)
	return
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update movie")
	w.Header().Set("Content-Type", "application/json")

	// id from request
	params := mux.Vars(r)
	for index, movie := range movies {
		movieId, err := strconv.Atoi(params["id"])
		if err != nil {
			panic(err)
		}
		if movie.Id == movieId {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movies
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = movieId
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteOneMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete movie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		movieId, err := strconv.Atoi(params["id"])
		if err != nil {
			panic(err)
		}
		if movie.Id == movieId {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
func main() {
	fmt.Println("API")
	r := mux.NewRouter()
	//out := add(2, 3)
	//fmt.Println(out)
	movies = append(movies, Movies{Id: 1, Name: "Silicon Valley", Genre: "Comedy", Rating: 4.5, Plot: "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.", Released: true})
	fmt.Println(movies)
	r.HandleFunc("/", serveMovie).Methods("GET")
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getOneMovie).Methods("GET")
	r.HandleFunc("/movie", createOneMovieData).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteOneMovie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}
