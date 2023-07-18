package service

import (
	"shl/bnk/cc/web"

	"github.com/labstack/echo/v4"
)

func TransService(e *echo.Echo) {
	//super admin
	e.GET("/transactions/account/:acc_id", web.GetAccount)
	e.GET("/transactions/:tx_id", web.TransactionSearch)

}
