package main

import (
	"PELAYAN_SERVICE/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterPelayanRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9070")
	log.Fatal(http.ListenAndServe("localhost:9070", r))
}
