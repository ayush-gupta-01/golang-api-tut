package routes

import (
	"ayush-gupta-01/golang-api-tut/controllers"

	"github.com/gorilla/mux"
)

func NewRouter(router *mux.Router) {
	router.HandleFunc("/", controllers.SayHello).Methods("Get")
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("Get")

}
