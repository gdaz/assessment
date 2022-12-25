package expense

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() {
	const url = "postgres://zcqexosc:GK-Gjf0HlysSDAe1vhfeOxwEq3p2uYlD@tiny.db.elephantsql.com/zcqexosc"

	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	var err error
	db, err = sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
	}
}
