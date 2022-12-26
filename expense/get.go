package expense

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func GetExpense(c *gin.Context) {
	fmt.Println("GetExpense" + c.ClientIP())
	fmt.Println("GetExpense" + c.Param("id"))
	rows := db.QueryRow("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1", c.Param("id"))

	expeense := ExpenseBody{}
	err := rows.Scan(&expeense.Id, &expeense.Title, &expeense.Amount, &expeense.Note, pq.Array(&expeense.Tags))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
	}

	c.JSON(http.StatusOK, expeense)
}

func GetExpenses(c *gin.Context) {
	fmt.Println("GetExpenses" + c.ClientIP())

	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query user statment:" + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Err{Message: "can't query all users:" + err.Error()})
	}

	expenses := []ExpenseBody{}

	for rows.Next() {
		e := ExpenseBody{}
		// tags := []string{}
		err := rows.Scan(&e.Id, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
		if err != nil {
			c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
		}
		// e.Tags = tags
		expenses = append(expenses, e)
	}

	c.JSON(http.StatusOK, expenses)
}
