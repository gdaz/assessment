package routers

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/gdaz/assessment/expense"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// f, _ := os.Create("D:\\gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	expenses := r.Group("/expenses", basicAuth())
	{
		expenses.POST("", expense.AddNewExpense)
		expenses.GET("/:id", expense.GetExpense)
		expenses.GET("", expense.GetExpenses)
<<<<<<< HEAD
		// expenses.PUT("/:id", expense.UpdateNewExpense)
=======
		expenses.PUT("/:id", expense.UpdateExpense)
>>>>>>> exp03
	}

	return r
}

func basicAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basic := ctx.Request.Header.Get("Authorization")
		log.Println("basic ", basic)

		if basic == "" {
			respWithUnauthorized(ctx)
			return
		}

		auth := strings.SplitN(basic, " ", 2)
		log.Println("auth ", auth)

		if len(auth) != 2 && auth[0] != "Basic" {
			respWithUnauthorized(ctx)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		log.Println("payload ", payload)
		log.Println("pair ", pair)
		// fmt.Sprint("pair: %s", pair)

		if len(pair) != 2 || !checkAllow(pair[0], pair[1]) {
			respWithUnauthorized(ctx)
			return
		}

		ctx.Next()
	}
}

func checkAllow(username, password string) bool {
	if username == "admin" && password == "secret" {
		return true
	}
	return false
}

func respWithUnauthorized(ctx *gin.Context) {
	resp := map[string]string{"error": "Unauthorized"}
	ctx.JSON(http.StatusUnauthorized, resp)
	ctx.Abort()
}
