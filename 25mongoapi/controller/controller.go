package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shublakhan-kaur/Golang_learning/mongoapi/model"
	"github.com/shublakhan-kaur/Golang_learning/mongoapi/service"
)

func InsertOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	service.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	service.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	allMovies := service.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	service.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	service.DeleteAllMovies()
	json.NewEncoder(w).Encode("All movies deleted")
}
