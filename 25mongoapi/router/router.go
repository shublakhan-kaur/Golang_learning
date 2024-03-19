package router

import (
	"github.com/gorilla/mux"
	"github.com/shublakhan-kaur/Golang_learning/mongoapi/controller"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie", controller.InsertOneMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
