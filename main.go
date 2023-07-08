package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"title`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname  string `json:"firstname"`
	Lastname string `json:"lastname"` 
}

var movies []Movie

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "433248", Title: "Airforce One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "468082", Title: "Welcome To Nigeria", Director: &Director{Firstname: "Ayodele", Lastname: "Adebanjo"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}