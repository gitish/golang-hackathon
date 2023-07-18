package model

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Model package initialized")
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