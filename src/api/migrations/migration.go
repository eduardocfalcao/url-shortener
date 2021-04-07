package migrations

import (
	"database/sql"
	"fmt"
)

type Migration struct {
	Version string
	Up      func(db *sql.Tx) error
	Down    func(db *sql.Tx) error
	done    bool
}

func InitMigrator(db *sql.DB) (*Migrator, error) {
	defaultMigrator.db = db

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		version varchar(255)
	);`); err != nil {
		err = fmt.Errorf("Unable to create `schema_migrations` table: %w", err)
		return defaultMigrator, err
	}

	rows, err := db.Query("SELECT version from `schema_migrations;`")
	if err != nil {
		return defaultMigrator, err
	}

	defer rows.Close()

	for rows.Next() {
		var version string
		err := rows.Scan(&version)
		if err != nil {
			return defaultMigrator, err
		}

		if defaultMigrator.Migrations[version] != nil {
			defaultMigrator.Migrations[version].done = true
		}
	}

	return defaultMigrator, nil
}
