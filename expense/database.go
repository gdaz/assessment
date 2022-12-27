package expense

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func InitDB() {
	// const url = "postgres://zcqexosc:GK-Gjf0HlysSDAe1vhfeOxwEq3p2uYlD@tiny.db.elephantsql.com/zcqexosc"

	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connection gracefully closed")
	}
}
