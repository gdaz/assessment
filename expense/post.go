package expense

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func AddNewExpense(c *gin.Context) {
	fmt.Println("AddNewExpense" + c.ClientIP())
	var expenseRequest ExpenseBody

	if err := c.BindJSON(&expenseRequest); err != nil {
		return
	}

	sqlStmt := "INSERT INTO expenses (title, amount, note, tags) VALUES($1, $2, $3, $4) RETURNING id"
	rows := db.QueryRow(sqlStmt, expenseRequest.Title, expenseRequest.Amount, expenseRequest.Note, pq.Array(expenseRequest.Tags))

	err := rows.Scan(&expenseRequest.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
	}

	c.JSON(http.StatusCreated, expenseRequest)
}
