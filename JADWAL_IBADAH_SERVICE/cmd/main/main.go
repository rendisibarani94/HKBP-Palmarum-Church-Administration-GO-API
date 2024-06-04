package main

import (
	"JADWAL_IBADAH_SERVICE/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterIbadahRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9060")
	log.Fatal(http.ListenAndServe("localhost:9060", r))
}
