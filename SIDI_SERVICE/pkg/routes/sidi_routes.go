package routes

import (
	"SIDI_SERVICE/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterSidiRoutes = func(router *mux.Router) {
	router.HandleFunc("/sidi/", controllers.CreateSidi).Methods("POST")
	router.HandleFunc("/sidi/", controllers.GetSidi).Methods("GET")
	router.HandleFunc("/sidi/{sidiId}", controllers.GetSidibyId).Methods("GET")
	router.HandleFunc("/sidi/{sidiId}", controllers.UpdateSidi).Methods("PUT")
	router.HandleFunc("/sidi/{sidiId}", controllers.DeleteSidi).Methods("DELETE")
}
