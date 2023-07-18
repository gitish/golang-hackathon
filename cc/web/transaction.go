package web

import (
	"fmt"
	"net/http"

	"shl/bnk/cc/dao"
	"shl/bnk/cc/model"

	"github.com/labstack/echo/v4"
)

func PostTransaction(c echo.Context) error {
	trans := model.PendingTransResponse{}
	arrangmentId := c.Param("arrangmentId")
	trans.Cid = arrangmentId
	trans.Comment = "PostTransaction"

	return c.JSON(http.StatusOK, trans)
}

func PendingTransaction(c echo.Context) error {
	trans := model.PendingTransResponse{}
	arrangmentId := c.Param("arrangmentId")
	trans.Cid = arrangmentId
	trans.Comment = "PendingTransaction"
	fmt.Println("Calling downstream")
	dpResponse, err1 := dao.DpPendingRequest()
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("after rutruning from downstream " + dpResponse.Status())
	trans.Comment = string(dpResponse.Body())
	return c.JSON(http.StatusOK, trans)
}
