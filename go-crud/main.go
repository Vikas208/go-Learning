package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movies struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"string"`
}

var movies []Movies

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var isFound bool = false
	for _, item := range movies {
		if item.ID == params["id"] {
			isFound = true
			json.NewEncoder(w).Encode(item)
			break
		}
	}

	if !isFound {
		json.NewEncoder(w).Encode(&Response{Status: http.StatusNotFound, Message: "Not Found"})
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(&Response{Status: http.StatusOK, Message: "Created"})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var isFound bool = false
	for index, item := range movies {
		if item.ID == params["id"] {
			isFound = true
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	if isFound {
		json.NewEncoder(w).Encode(&Response{Status: http.StatusOK, Message: "Deleted"})
	} else {
		json.NewEncoder(w).Encode(&Response{
			Status:  http.StatusNotFound,
			Message: "Not Found",
		})
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var isFound bool = false
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)

	for index, item := range movies {
		if item.ID == params["id"] {
			isFound = true
			movie.ID = movies[index].ID
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(&Response{
				Status:  http.StatusOK,
				Message: "Updated",
			})
			break
		}
	}

	if !isFound {
		json.NewEncoder(w).Encode(&Response{
			Status:  http.StatusNotFound,
			Message: "Not Found",
		})
	}
}

func seedersData() {
	movies = append(movies, Movies{
		ID:    "35",
		Isbn:  "302555",
		Title: "after",
		Director: &Director{
			Name: "Allie",
		},
	})
	movies = append(movies, Movies{
		ID:    "64",
		Isbn:  "390554",
		Title: "making",
		Director: &Director{
			Name: "Leon",
		},
	})

}

func main() {

	defer func() {
		r := recover()
		if r != nil {
			log.Fatal(r)
		}
	}()
	seedersData()
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
