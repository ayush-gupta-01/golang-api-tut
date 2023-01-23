package services

import (
	"ayush-gupta-01/golang-api-tut/models"
	"ayush-gupta-01/golang-api-tut/repository"

	"gorm.io/gorm"
)

func GetAllBooks(gormDB *gorm.DB) []models.Books {
	var res []models.Books
	data := repository.Repo.GetAllBooks(res, gormDB)
	return data
}
