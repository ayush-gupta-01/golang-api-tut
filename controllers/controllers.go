package controllers

import (
	"ayush-gupta-01/golang-api-tut/connect"
	"ayush-gupta-01/golang-api-tut/models"
	"ayush-gupta-01/golang-api-tut/services"
	"ayush-gupta-01/golang-api-tut/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var NewServiceRepository = services.NewBookService()

type response struct {
	Msg string `json:"msg"`
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	res := response{"Hello World"}
	utils.WriteJsonData(w, res, 200)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	// body : [123 34 105 100 34 58 49 52 44 34 98 111 111 107 110 97 109 101 34 58
	// 34 97 115 100 97 115 100 97 115 100 34 44 34 112 117 98 108 105 99 97
	// 116 105 111 110 115 34 58 34 115 97 100 107 108 106 97 115 108 107 100 106 97 115 108 107 106 100 97 108 107 115 106 100 108 107 97 115 106
	// 100 97 115 34 125]
	if err != nil {
		utils.WriteJsonData(w, err, 400)
	}
	details := models.Books{}
	if err := json.Unmarshal(body, &details); err != nil {
		utils.WriteJsonData(w, err, 400)
	}
	// details by unmarshal : {12 asdasdasd sadkljaslkdjaslkjdalksjdlkasjdas}
	data, _, status := NewServiceRepository.CreateBook(w, details)
	utils.WriteJsonData(w, data, status)
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

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var GormDB = connect.GormDB()
	params := mux.Vars(r)
	id := params["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJsonData(w, err, 400)
	}
	details := models.Books{}
	if err := json.Unmarshal(body, &details); err != nil {
		utils.WriteJsonData(w, err, 400)
	}
	book := models.Books{}
	result := GormDB.Where("Id=?", id).Find(&book)
	if result.RowsAffected == 0 {
		response := models.Response{
			ResponseMsg:  "failed",
			ResponseCode: 400,
			Data:         "Book not found",
			Error:        nil,
		}
		utils.WriteJsonData(w, response, 400)
	}
	book.BookName = details.BookName
	book.Publications = details.Publications
	result = GormDB.Save(&book)
	response := models.Response{
		ResponseMsg:  "success",
		ResponseCode: 200,
		Data:         "Book updated successfully",
		Error:        nil,
	}
	utils.WriteJsonData(w, response, 200)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var GormDB = connect.GormDB()
	params := mux.Vars(r)
	id := params["id"]
	book := models.Books{}
	result := GormDB.Where("Id=?", id).Delete(&book)
	fmt.Println(result)
	response := models.Response{
		ResponseMsg:  "success",
		ResponseCode: 200,
		Data:         "Book deleted successfully",
		Error:        nil,
	}
	utils.WriteJsonData(w, response, 200)
}
