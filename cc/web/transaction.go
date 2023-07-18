package web

import (
	"fmt"
	"net/http"
	"time"

	"shl/bnk/cc/dao"
	"shl/bnk/cc/model"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	//trans := model.PendingTransResponse{}
	db := dao.SetupDB()
	fmt.Println("Getting transactions...")
	// var response []JsonResponse
	var transactions []model.Transaction
	acc_id := c.Param("acc_id")
	fmt.Println(acc_id)
	status := c.QueryParam("status")
	// Get all movies from movies table that don't have movieID = "1"
	query := "SELECT tx_id, acc_id, tx_ts, status, amount, merchantname, merchant_id, tx_type, tx_details FROM transaction.transactions where acc_id='" + acc_id + "'"
	if status != "" {
		query = query + " AND status='" + status + "'"
	}
	rows, err := db.Query(query)

	// check errors
	dao.CheckErr(err)

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
		dao.CheckErr(err)

		transactions = append(transactions, model.Transaction{TransactionID: transactionID, AccountID: accountID, TimeStamp: timestamp, Status: status, Amount: amount, MerchantName: merchantName, MerchantID: merchantId, Type: txnType, Details: details})
	}

	var response = model.JsonResponse{Type: "success", Data: transactions}
	return c.JSON(http.StatusOK, response)
}

func TransactionSearch(c echo.Context) error {
	tx_id := c.Param("tx_id")
	//trans := model.PendingTransResponse{}
	db := dao.SetupDB()
	fmt.Println("Getting transactions...")
	// var response []JsonResponse
	var transactions []model.Transaction
	acc_id := c.Param("acc_id")
	fmt.Println(acc_id)
	// Get all movies from movies table that don't have movieID = "1"
	query := "SELECT tx_id, acc_id, tx_ts, status, amount, merchantname, merchant_id, tx_type, tx_details FROM transaction.transactions where tx_id='" + tx_id + "'"

	rows, err := db.Query(query)

	// check errors
	dao.CheckErr(err)

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
		dao.CheckErr(err)

		transactions = append(transactions, model.Transaction{TransactionID: transactionID, AccountID: accountID, TimeStamp: timestamp, Status: status, Amount: amount, MerchantName: merchantName, MerchantID: merchantId, Type: txnType, Details: details})
	}

	var response = model.JsonResponse{Type: "success", Data: transactions}
	return c.JSON(http.StatusOK, response)
}
