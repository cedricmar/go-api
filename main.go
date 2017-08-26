package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weebagency/go-api/controllers"
	"github.com/weebagency/go-api/migrate"
)

func main() {
	// Controllers
	hello := &controllers.HelloController{}
	info := &controllers.InfoController{}

	r := mux.NewRouter()

	migrate.Up()

	// Routes
	r.HandleFunc("/", hello.Index).Methods("GET")
	r.HandleFunc("/api/v1/info", info.Info).Methods("GET")
	http.Handle("/", r)

	port := ":8080"
	log.Printf("Listening to %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
