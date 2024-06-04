package routes

import (
	"PELAYAN_SERVICE/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterPelayanRoutes = func(router *mux.Router) {
	router.HandleFunc("/pelayan/", controllers.CreatePelayan).Methods("POST")
	router.HandleFunc("/pelayan/", controllers.GetPelayan).Methods("GET")
	router.HandleFunc("/pelayan/{pelayanId}", controllers.GetPelayanbyId).Methods("GET")
	router.HandleFunc("/pelayan/{pelayanId}", controllers.UpdatePelayan).Methods("PUT")
	router.HandleFunc("/pelayan/{pelayanId}", controllers.DeletePelayan).Methods("DELETE")
}
