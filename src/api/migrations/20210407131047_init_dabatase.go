package migrations

import (
	"database/sql"
	"errors"
)

func init() {
	defaultMigrator.AddMigration(&Migration{
		Version: "20210407131047",
		Up:      migration_20210407131047_init_dabatase_up,
		Down:    migration_20210407131047_init_dabatase_down,
	})
}

func migration_20210407131047_init_dabatase_up(db *sql.Tx) error {
	return errors.New("Not Implemented")
}

func migration_20210407131047_init_dabatase_down(db *sql.Tx) error {
	return errors.New("Not Implemented")
}
