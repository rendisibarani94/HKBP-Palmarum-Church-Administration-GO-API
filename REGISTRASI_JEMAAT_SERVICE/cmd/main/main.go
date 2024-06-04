package main

import (
	"REGISTRASI_JEMAAT_SERVICE/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterJemaatRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9020")
	log.Fatal(http.ListenAndServe("localhost:9020", r))
}
