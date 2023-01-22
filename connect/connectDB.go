package connect

import (
	"ayush-gupta-01/golang-api-tut/config"
	"ayush-gupta-01/golang-api-tut/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var GormDB *gorm.DB

func InitDBConnection() {

	GormDB, err := gorm.Open("postgres", config.ConnectionString())
	if err != nil {
		log.Fatalf("Unable to connect to Gorm, exiting : %v", err)
	}

	GormDB.AutoMigrate(&models.Books{})

	GormDB.Close()
}
