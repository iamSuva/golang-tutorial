package router

import (
	"github.com/gorilla/mux"
	"github.com/iamSuva/mongoapi/controllers"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api/movies",controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controllers.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall",controllers.DeleteAllMovie).Methods("DELETE")
	
	return router;
}
