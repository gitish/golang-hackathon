package service

import (
	"fmt"
	"os"

	"shl/bnk/cc/web"

	"github.com/labstack/echo/v4"
)

func Transactionservice(e *echo.Echo) {
	//super admin

	e.GET("/transactions/account", web.Accounts)
	e.GET("/transactions/account/:acc_id", web.AccountSearch)
	e.GET("/transactions/:tx_id", web.TransactionSearch)


	DEFAULT_STATIC_DIR := "./static"
	STATIC_DIR := os.Getenv("STATIC_DIR")
	if STATIC_DIR == "" {
		fmt.Println("STATIC_DIR is not set as Environment Variable")
		STATIC_DIR = DEFAULT_STATIC_DIR
	} else {
		fmt.Println("Content DIR: " + STATIC_DIR)
	}
	paths := []string{
		"/",
	}
	for _, url := range paths {
		e.Static(url, STATIC_DIR)
	}

	e.Static("/images", web.GetDataDir()+"/images")
}
