package controllers

import (
	"ayush-gupta-01/golang-api-tut/connect"
	"ayush-gupta-01/golang-api-tut/models"
	"ayush-gupta-01/golang-api-tut/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type response struct {
	Msg string `json:"msg"`
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	res := response{"Hello World"}
	utils.WriteJsonData(w, res, 200)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var GormDB = connect.GormDB()
	var res []models.Books
	result := GormDB.Find(&res)
	fmt.Println(result)
	utils.WriteJsonData(w, res, 200)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	var GormDB = connect.GormDB()
	params := mux.Vars(r)
	id := params["id"]
	var res models.Books
	result := GormDB.Where("Id=?", id).Find(&res)
	fmt.Println(result)
	utils.WriteJsonData(w, res, 200)
}
