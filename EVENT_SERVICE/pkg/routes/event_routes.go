package routes

import (
	"EVENT_SERVICE/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterEventRoutes = func(router *mux.Router) {
	router.HandleFunc("/event/", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/event/", controllers.GetEvent).Methods("GET")
	router.HandleFunc("/event/{eventId}", controllers.GetEventbyId).Methods("GET")
	router.HandleFunc("/event/{eventId}", controllers.UpdateEvent).Methods("PUT")
	router.HandleFunc("/event/{eventId}", controllers.DeleteEvent).Methods("DELETE")
}
