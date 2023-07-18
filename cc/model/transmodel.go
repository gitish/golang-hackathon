package model

import "fmt"

func init() {
	fmt.Println("Model package initialized")
}

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "workshop"
)
