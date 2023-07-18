package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "workshop"
)

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

type Transaction struct {
	TransactionID string    `json:"transactionId"`
	AccountID     string    `json:"accountId"`
	TimeStamp     time.Time `json:"timestamp"`
	Status        string    `json:"status"`
	Amount        int       `json:"amount"`
	MerchantName  string    `json:"merchantName"`
	MerchantID    string    `json:"merchantId"`
	Type          string    `json:"type"`
	Details       string    `json:"details"`
}

type JsonResponse struct {
	Type    string        `json:"type"`
	Data    []Transaction `json:"data"`
	Message string        `json:"message"`
}

// Main function
func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all transactions
	router.HandleFunc("/transactions/", GetTransactions).Methods("GET")

	// serve the app
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Get all movies

// response and request handlers
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting transactions...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT tx_id, acc_id, tx_ts, status, amount, merchantname, merchant_id, tx_type, tx_details FROM transaction.transactions")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var transactions []Transaction

	// Foreach movie
	for rows.Next() {
		var transactionID string
		var accountID string
		var timestamp time.Time
		var status string
		var amount int
		var merchantName string
		var merchantId string
		var txnType string
		var details string

		err = rows.Scan(&transactionID, &accountID, &timestamp, &status, &amount, &merchantName, &merchantId, &txnType, &details)

		// check errors
		checkErr(err)

		transactions = append(transactions, Transaction{TransactionID: transactionID, AccountID: accountID, TimeStamp: timestamp, Status: status, Amount: amount, MerchantName: merchantName, MerchantID: merchantId, Type: txnType, Details: details})
	}

	var response = JsonResponse{Type: "success", Data: transactions}

	json.NewEncoder(w).Encode(response)
}
