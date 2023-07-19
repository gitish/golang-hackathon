package web

import (
	"fmt"
	"net/http"

	"shl/bnk/cc/dao"
	"shl/bnk/cc/model"

	"github.com/huandu/go-sqlbuilder"
	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {

	acc_id := c.Param("acc_id")
	fmt.Println(acc_id)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("tx_id",
		"acc_id", "tx_ts", "status",
		"amount", "merchantname",
		"merchant_id", "tx_type", "tx_details").
		From("transaction.transactions").
		Where("acc_id='" + acc_id + "'")

	status := c.QueryParam("status")
	if status != "" {
		sb.Where("status='" + status + "'")
	}

	response := getTransactionDetails(sb)
	return c.JSON(http.StatusOK, response)
}

func TransactionSearch(c echo.Context) error {
	tx_id := c.Param("tx_id")

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("tx_id",
		"acc_id", "tx_ts", "status",
		"amount", "merchantname",
		"merchant_id", "tx_type", "tx_details").
		From("transaction.transactions").
		Where("tx_id='" + tx_id + "'")

	response := getTransactionDetails(sb)
	return c.JSON(http.StatusOK, response)

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
