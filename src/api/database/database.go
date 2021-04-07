package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(connectionString string) (*sql.DB, error) {
	return sql.Open("mysql", connectionString)
}
