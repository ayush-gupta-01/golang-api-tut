package services

import (
	"ayush-gupta-01/golang-api-tut/models"
	"ayush-gupta-01/golang-api-tut/repository"
	"net/http"
)

type BooksService interface{
	CreateBook(http.ResponseWriter ,models.Books) (*models.Books , error , int)
}

type bookService struct{}

func NewBookService() BooksService {
	return &bookService{}
}

var NewBookRepository = repository.NewBookRepository()

func (b *bookService) CreateBook(rw http.ResponseWriter ,details models.Books)(*models.Books , error , int) {
	data , err , status := NewBookRepository.CreateBook(rw , details)
	if err != nil {
		return nil , err , status
	}
	return data , nil , status
}