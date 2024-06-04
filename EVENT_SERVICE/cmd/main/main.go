package main

import (
	"EVENT_SERVICE/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterEventRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:7010")
	log.Fatal(http.ListenAndServe("localhost:7010", r))
}
