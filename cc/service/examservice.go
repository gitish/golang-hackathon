package service

import (
	"fmt"
	"os"

	"shl/bnk/cc/web"

	"github.com/labstack/echo/v4"
)

func Examservice(e *echo.Echo) {
	//super admin
	e.GET("/post-transaction/:arrangmentId", web.PostTransaction)
	e.GET("/pending-transaction/:arrangmentId", web.PendingTransaction)

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
