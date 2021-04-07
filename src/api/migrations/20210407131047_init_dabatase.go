package migrations

import (
	"database/sql"
)

func init() {
	defaultMigrator.AddMigration(&Migration{
		Version: "20210407131047",
		Up:      migration_20210407131047_init_dabatase_up,
		Down:    migration_20210407131047_init_dabatase_down,
	})
}

func migration_20210407131047_init_dabatase_up(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id int NOT NULL AUTO_INCREMENT,
			name varchar(255),
			email varchar(500),
			PRIMARY KEY(id)
	);`)
	return err
}

func migration_20210407131047_init_dabatase_down(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS users`)

	return err
}
