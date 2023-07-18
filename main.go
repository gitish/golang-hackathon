package main

import (
	"fmt"
	"os"

	service "shl/bnk/cc/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Hello, world.")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//exams services exposed here
	service.TransService(e)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8085"
	}
	e.Logger.Fatal(e.Start(":" + httpPort))
}
