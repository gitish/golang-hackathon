package web

import (
	"fmt"
	"net/http"
	"strconv"

	"shl/bnk/cc/dao"
	"shl/bnk/cc/model"

	"github.com/huandu/go-sqlbuilder"
	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {

	acc_id := c.Param("acc_id")
	fmt.Println(acc_id)

	sb := TransQuery().
		Where("acc_id='" + acc_id + "'")

	status := c.QueryParam("status")
	if status != "" {
		sb.Where("status='" + status + "'")
	}
	/**
	pagination start here. For pagination pageSize is must
	however page is optional and if unavailable then it jsut return first page
	**/
	pageSize := c.QueryParam("pageSize")
	if pageSize != "" {
		limit, _ := strconv.Atoi(pageSize)
		sb.Limit(limit)

		page := c.QueryParam("page")
		if page != "" {
			p, _ := strconv.Atoi(page)
			offset := limit * (p - 1)
			sb.Offset(offset)
		}
	}

	response := getTransactionDetails(sb)
	return c.JSON(http.StatusOK, response)
}

func TransactionSearch(c echo.Context) error {
	tx_id := c.Param("tx_id")

	sb := TransQuery().
		Where("tx_id='" + tx_id + "'")

	response := getTransactionDetails(sb)
	return c.JSON(http.StatusOK, response)

}

func TransQuery() *sqlbuilder.SelectBuilder {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("tx_id",
		"acc_id", "tx_ts", "status",
		"amount", "merchantname",
		"merchant_id", "tx_type", "tx_details").
		From("transaction.transactions")
	return sb
}

func getTransactionDetails(sb *sqlbuilder.SelectBuilder) model.JsonResponse {
	db := dao.SetupDB()
	defer db.Close()
	rows, err := db.Query(sb.String())

	dao.CheckErr(err)

	var transactions []model.Transaction
	for rows.Next() {
		var t model.Transaction

		err = rows.Scan(&t.TransactionID,
			&t.AccountID,
			&t.TimeStamp,
			&t.Status,
			&t.Amount,
			&t.MerchantName,
			&t.MerchantID,
			&t.Type, &t.Details)

		dao.CheckErr(err)
		transactions = append(transactions, t)
	}

	var response = model.JsonResponse{Type: "success", Data: transactions}
	return response
}
