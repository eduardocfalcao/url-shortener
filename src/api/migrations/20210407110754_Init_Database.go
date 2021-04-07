package migrations

import (
	"database/sql"
	"errors"
)

func init() {
	defaultMigrator.AddMigration(&Migration{
		Version: "20210407110754",
		Up:      migration_20210407110754_Init_Database_up,
		Down:    migration_20210407110754_Init_Database_down,
	})
}

func migration_20210407110754_Init_Database_up(db *sql.Tx) error {
	return errors.New("Not Implemented")
}

func migration_20210407110754_Init_Database_down(db *sql.Tx) error {
	return errors.New("Not Implemented")
}
