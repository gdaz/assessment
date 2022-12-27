package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: r,
	}

	go func() {
		log.Println("Routines")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	} else {
		log.Println("Server connection gracefully closed")
	}

	expense.CloseDB()

	log.Println("Server exiting")
}
