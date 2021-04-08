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
		CREATE TABLE IF NOT EXISTS short_urls (
			id int(11) NOT NULL AUTO_INCREMENT,
			name varchar(255),
			shorturl varchar(30),
			url varchar(500),
			PRIMARY KEY(id),
			UNIQUE KEY(shorturl)
	);`)
	return err
}

func migration_20210407131047_init_dabatase_down(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS short_urls`)

	return err
}
