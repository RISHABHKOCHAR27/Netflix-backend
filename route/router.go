package route

import (
	"github.com/gorilla/mux"
	"main.go/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreatreMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/movie", controller.DeleteAllMovies).Methods("DELETE")

	return r
}
