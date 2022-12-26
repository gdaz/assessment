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

}
