package model

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("Model package initialized")
}

type Transaction struct {
	TransactionID string    `json:"transactionId" db:"tx_id"`
	AccountID     string    `json:"accountId"  db:"acc_id"`
	TimeStamp     time.Time `json:"timestamp"  db:"tx_ts"`
	Status        string    `json:"status"  db:"status"`
	Amount        int       `json:"amount"  db:"amount"`
	MerchantName  string    `json:"merchantName"  db:"merchantname"`
	MerchantID    string    `json:"merchantId"  db:"merchant_id"`
	Type          string    `json:"type"  db:"tx_type"`
	Details       string    `json:"details"  db:"tx_details"`
}

type JsonResponse struct {
	Type    string        `json:"type"`
	Data    []Transaction `json:"data"`
	Message string        `json:"message"`
}
