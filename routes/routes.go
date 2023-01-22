package routes

import (
	"ayush-gupta-01/golang-api-tut/controllers"

	"github.com/gorilla/mux"
)

func NewRouter(router *mux.Router) {
	router.HandleFunc("/", controllers.SayHello).Methods("Get")
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("Get")
	router.HandleFunc("/book", controllers.CreateBook).Methods("Post")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("Get")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("Delete")
}
