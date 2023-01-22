package connect

import (
	"ayush-gupta-01/golang-api-tut/config"
	"database/sql"
	"fmt"

	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB
var gormDB *gorm.DB

func SQLDB() *sql.DB {
	return db
}

func GormDB() *gorm.DB {
	return gormDB
}

func InitDBConnection() {
	sqldb, err := sql.Open(config.DB_DRIVER, config.ConnectionString())
	if err != nil {
		log.Fatalf("Unable to connect to database, exiting: %v", err)
	}
	db = sqldb

	//Configure Database connection pool
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	//Reuse the above connection pool in gorm
	gormDatabase, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	fmt.Println(gormDatabase)
	if err != nil {
		log.Fatalf("Unable to connect to Gorm, exiting: %v", err)
	}
	gormDB = gormDatabase
}
