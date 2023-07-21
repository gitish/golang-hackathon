package dao

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// DB set up
func SetupDB() *sql.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT, portErr := strconv.Atoi(os.Getenv("DB_PORT"))
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	if DB_HOST == "" {
		DB_HOST = "localhost"
	}
	if portErr != nil {
		DB_PORT = 5432
	}
	if DB_USER == "" {
		DB_USER = "postgres"
	}
	if DB_PASSWORD == "" {
		DB_PASSWORD = "postgres"
	}
	if DB_NAME == "" {
		DB_NAME = "workshop"
	}
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	fmt.Println("dbInfo: " + dbinfo)
	db, err := sql.Open("postgres", dbinfo)

	CheckErr(err)

	return db
}

// Function for handling errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
