package migrations

import (
	"database/sql"
	"errors"
)

func init() {
	migrations.AddMigration(&Migration{
		Version: "{{.Version}}",
		Up:      migration_{{.Version}}_{{.Name}}_up,
		Down:    migration_{{.Version}}_{{.Name}}_down,
	})
}

func migration_{{.Version}}_{{.Name}}_up(db *sql.Tx) error {
	return errors.New("Not Implemented")
}

func migration_{{.Version}}_{{.Name}}_down(db *sql.Tx) error {
	return errors.New("Not Implemented")
}
