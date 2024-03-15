package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shublakhan-kaur/Golang_learning/mongoapi/controller"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/movie", controller.InsertOneMovie).Methods("POST")
	r.HandleFunc("movie/{id}", controller.MarkedAsWatched).Methods("PUT")
	r.HandleFunc("movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("movies", controller.DeleteAllMovies).Methods("DELETE")
	fmt.Println(http.ListenAndServe(":4000", r))
}
