package routes

import (
	"BAPTIS_SERVICE/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBaptisRoutes = func(router *mux.Router) {
	router.HandleFunc("/baptis/", controllers.CreateBaptis).Methods("POST")
	router.HandleFunc("/baptis/", controllers.GetBaptis).Methods("GET")
	router.HandleFunc("/baptis/{baptisId}", controllers.GetBaptisbyId).Methods("GET")
	router.HandleFunc("/baptis/{baptisId}", controllers.UpdateBaptis).Methods("PUT")
	router.HandleFunc("/baptis/{baptisId}", controllers.DeleteBaptis).Methods("DELETE")
}
