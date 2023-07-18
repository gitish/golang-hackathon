package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	//trans := model.PendingTransResponse{}
	acc_id := c.Param("acc_id")

	return c.JSON(http.StatusOK, acc_id)
}

func TransactionSearch(c echo.Context) error {
	tx_id := c.Param("tx_id")
	return c.JSON(http.StatusOK, tx_id)
}
