package main

import (
	"SIDI_SERVICE/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterSidiRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9050")
	log.Fatal(http.ListenAndServe("localhost:9050", r))
}
