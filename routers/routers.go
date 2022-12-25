package routers

import (
	"github.com/gdaz/assessment/expense"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	expenses := r.Group("/expenses")
	{
		expenses.POST("", expense.AddNewExpense)
		expenses.GET("/:id", expense.GetExpense)
		expenses.GET("", expense.GetExpenses)
		expenses.PUT("/:id", expense.UpdateNewExpense)
	}

	return r
}
