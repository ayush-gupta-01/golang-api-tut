package config

import "fmt"

var (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_PROTOCOL = "tcp"
	DB_USER     = "postgres"
	DB_PASSWORD = "your_password"
	DB_NAME     = "gobookstore"
	DB_DRIVER   = "pgx"
)

func ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
}
