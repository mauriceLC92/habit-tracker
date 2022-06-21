package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	dbPackage "github.com/mauriceLC92/habit-tracker/internal/db/habits"

	"database/sql"

	_ "github.com/lib/pq"
)

type config struct {
	httpAddr string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger

	Queries *dbPackage.Queries
}

func main() {
	var conf config

	flag.StringVar(&conf.httpAddr, "addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://postgres:mysecretpassword@localhost/habit-tracker?sslmode=disable", "PostgreSQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	queries := dbPackage.New(db)

	app := &application{
		Queries:  queries,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:     conf.httpAddr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s \n", conf.httpAddr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
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
