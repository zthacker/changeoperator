package coAPI

import (
	"database/sql"
	"log"
	"os"
)

type ChangeOperator struct {
	PostgresClient *sql.DB
}

func setupPostgres() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("PG_CONN"))
	if err != nil {
		log.Fatalf("Could not setup PostgresDB: %s", err)
	}
	return db
}

func NewClients() *ChangeOperator {
	pgDB := setupPostgres()

	return &ChangeOperator{PostgresClient: pgDB}
}
