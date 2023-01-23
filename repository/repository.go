package repository

import (
	"ayush-gupta-01/golang-api-tut/models"
	"log"

	"gorm.io/gorm"
)

type Repo interface {
	GetAllBooks(m []models.Books, gormDB *gorm.DB) []models.Books
	GetBookById(m models.Books, id int64, gormDB *gorm.DB) models.Books
}

func GetAllBooks(m []models.Books, gormDB *gorm.DB) []models.Books {
	if err := gormDB.Find(&m); err != nil {
		log.Fatal("Error")
	}
	return m
}

func GetBookById(m models.Books, id int64, gormDB *gorm.DB) models.Books {
	if err := gormDB.Where("ID=?", id).Find(&m); err != nil {
		log.Fatal("Error")
	}
	return m
}
