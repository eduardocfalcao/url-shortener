package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const dbDriver = "postgres"

func NewConnection(connectionString string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, connectionString)
	if err != nil {
		return nil, fmt.Errorf("An error happened when trying to create the database connection: %w", err)
	}
	return db, nil
}
