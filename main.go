package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"apps/api/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Controllers
	hello := &controllers.HelloController{}
	info := &controllers.InfoController{}

	r := mux.NewRouter()

	db, err := sql.Open("mysql", "admin:130779@/api")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", db)

	// Routes
	r.HandleFunc("/", hello.Index).Methods("GET")
	r.HandleFunc("/api/v1/info", info.Info).Methods("GET")
	http.Handle("/", r)

	port := ":8080"
	log.Printf("Listening to %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
