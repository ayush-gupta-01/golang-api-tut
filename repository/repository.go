package repository

import (
	"ayush-gupta-01/golang-api-tut/connect"
	"ayush-gupta-01/golang-api-tut/models"
	"log"
	"net/http"
)

type BookRepository interface {
	CreateBook(http.ResponseWriter, models.Books) (*models.Books, error,int)
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (b *bookRepository) CreateBook(w http.ResponseWriter, details models.Books) (*models.Books, error ,int) {
	var GormDB = connect.GormDB()
	if err := GormDB.Create(&details).Error; err != nil {
		log.Fatalf("Error while creating book: %v \n", err.Error())
	}
	return &details,nil,200
}
