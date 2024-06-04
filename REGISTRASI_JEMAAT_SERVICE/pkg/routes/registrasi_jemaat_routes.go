package routes

import (
	"REGISTRASI_JEMAAT_SERVICE/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterJemaatRoutes = func(router *mux.Router) {
	router.HandleFunc("/reg_jemaat/", controllers.CreateJemaat).Methods("POST")
	router.HandleFunc("/reg_jemaat/", controllers.GetJemaat).Methods("GET")
	router.HandleFunc("/reg_jemaat/{jemaatId}", controllers.GetJemaatbyId).Methods("GET")
	router.HandleFunc("/reg_jemaat/{jemaatId}", controllers.UpdateJemaat).Methods("PUT")
	router.HandleFunc("/reg_jemaat/{jemaatId}", controllers.DeleteJemaat).Methods("DELETE")
}
