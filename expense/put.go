package expense

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func UpdateExpense(c *gin.Context) {
	var expenseRequest ExpenseBody

	if err := c.BindJSON(&expenseRequest); err != nil {
		return
	}

	sqlStmt := `UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1 RETURNING id`
	// stmt, err := db.Prepare(sqlStmt)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// stmt.Exec(&expenseRequest.Id, expenseRequest.Title, expenseRequest.Amount, expenseRequest.Note, pq.Array(expenseRequest.Tags))
	rows := db.QueryRow(sqlStmt, expenseRequest.Id, expenseRequest.Title, expenseRequest.Amount, expenseRequest.Note, pq.Array(expenseRequest.Tags))

	err := rows.Scan(&expenseRequest.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
	}

	c.JSON(http.StatusOK, expenseRequest)
}
