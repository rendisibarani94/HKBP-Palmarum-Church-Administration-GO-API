package routes

import (
	"JADWAL_IBADAH_SERVICE/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterIbadahRoutes = func(router *mux.Router) {
	router.HandleFunc("/ibadah/", controllers.CreateIbadah).Methods("POST")
	router.HandleFunc("/ibadah/", controllers.GetIbadah).Methods("GET")
	router.HandleFunc("/ibadah/{ibadahId}", controllers.GetIbadahbyId).Methods("GET")
	router.HandleFunc("/ibadah/{ibadahId}", controllers.UpdateIbadah).Methods("PUT")
	router.HandleFunc("/ibadah/{ibadahId}", controllers.DeleteIbadah).Methods("DELETE")
}
