package dao

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	_ "github.com/lib/pq"
	"database/sql"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "workshop"
)

func main1() {

	db := setupDB()

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

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {

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
