package dao

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func DpPendingRequest() (*resty.Response, error) {
	client := resty.New()
	fmt.Println("Calling mock service..")
	resp, err := client.R().
		EnableTrace().
		Get("http://localhost:9080/dp/fs/test")

	return resp, err
}
