package routes

import (
	"JEMAAT_PINDAH_SERVICE/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterJemaatPindahRoutes = func(router *mux.Router) {
	router.HandleFunc("/pindah/", controllers.CreateJemaatPindah).Methods("POST")
	router.HandleFunc("/pindah/", controllers.GetJemaatPindah).Methods("GET")
	router.HandleFunc("/pindah/{pindahId}", controllers.GetJemaatPindahbyId).Methods("GET")
	router.HandleFunc("/pindah/{pindahId}", controllers.UpdateJemaatPindah).Methods("PUT")
	router.HandleFunc("/pindah/{pindahId}", controllers.DeleteJemaatPindah).Methods("DELETE")
}
