package main

import (
	"ayush-gupta-01/golang-api-tut/connect"
	"ayush-gupta-01/golang-api-tut/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	connect.InitDBConnection()

	r := mux.NewRouter()

	routes.NewRouter(r)
	server := &http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	log.Printf("Server is runnning on port : %s", server.Addr)
	server.ListenAndServe()
}
