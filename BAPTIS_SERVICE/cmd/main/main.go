package main

import (
	"BAPTIS_SERVICE/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBaptisRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9030")
	log.Fatal(http.ListenAndServe("localhost:9030", r))
}
