package cmd

import (
	"database/sql"
	"fmt"
	"log"
)

//dbInit creates DB connection and runs migration
func dbInit(usr, pwd, dbHost, dbPort, dbName string) *sql.DB {
	if len(dbHost) == 0 {
		log.Fatalf("invalid POSTGRES_HOST")
	}
	if len(dbName) == 0 {
		log.Fatalf("invalid POSTGRES_PORT")
	}
	if len(dbPort) == 0 {
		log.Fatalf("invalid POSTGRES_DB")
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", usr, pwd, dbHost, dbPort, dbName)
	log.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("invalid POSTGRES_DB open", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("invalid POSTGRES_DB Ping failed", err)
	}
	mustHandleMigrations(db, true)
	return db
}
