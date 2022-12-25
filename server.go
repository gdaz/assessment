package main

import (
	"fmt"
	"os"

	"github.com/gdaz/assessment/expense"

	"github.com/gdaz/assessment/routers"
	_ "github.com/lib/pq"
)

func main() {
	expense.InitDB()

	fmt.Println("Please use server.go for main file")
	fmt.Println("PORT: ", os.Getenv("PORT"))

	r := routers.InitRouter()

	r.Run(os.Getenv("PORT"))
}
