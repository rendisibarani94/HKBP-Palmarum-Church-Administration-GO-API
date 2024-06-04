package routes

import (
	"RPP_SERVICE/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRppRoutes = func(router *mux.Router) {
	router.HandleFunc("/rpp/", controllers.CreateRpp).Methods("POST")
	router.HandleFunc("/rpp/", controllers.GetRpp).Methods("GET")
	router.HandleFunc("/rpp/{rppId}", controllers.GetRppbyId).Methods("GET")
	router.HandleFunc("/rpp/{rppId}", controllers.UpdateRpp).Methods("PUT")
	router.HandleFunc("/rpp/{rppId}", controllers.DeleteRpp).Methods("DELETE")
}
