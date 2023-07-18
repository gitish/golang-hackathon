package model

import "fmt"

func init() {
	fmt.Println("Model package initialized")
}

type PendingTransResponse struct {
	Cid     string `json:"cid"`
	Comment string `json:"comment"`
}

type PostTransResponse struct {
	PendingTransResponse
}
