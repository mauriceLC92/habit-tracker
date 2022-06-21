package main

import (
	"fmt"
	"log"

	db "github.com/mauriceLC92/habit-tracker/internal/db/habits"

	"database/sql"

	_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger

	Queries *db.Queries
}

func main() {
	fmt.Print("fghj")

}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
