package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	//trans := model.PendingTransResponse{}
	acc_id := c.Param("acc_id")
	//trans.Cid = arrangmentId
	//trans.Comment = "PostTransaction"

	return c.JSON(http.StatusOK, acc_id)
}

func PendingTransaction(c echo.Context) error {
	acc_id := c.Param("acc_id")
	return c.JSON(http.StatusOK, acc_id)
}
