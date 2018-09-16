package main

import (
	"log"
	"net/http"

	"github.com/cedricmar/go-api/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// Controllers
	hello := &controllers.HelloController{}
	info := &controllers.InfoController{}

	r := mux.NewRouter()

	// @TODO - make a command of that
	//migrate.Up()

	// Routes
	r.HandleFunc("/", hello.Index).Methods("GET")
	r.HandleFunc("/api/v1/info", info.Info).Methods("GET")
	http.Handle("/", r)

	port := ":8080"
	log.Printf("Listening to %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
