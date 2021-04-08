package migrations

import (
	"context"
	"database/sql"
	"fmt"
)

type Migrator struct {
	db         *sql.DB
	Versions   []string
	Migrations map[string]*Migration
}

var defaultMigrator = &Migrator{
	Migrations: make(map[string]*Migration),
}

func (m *Migrator) AddMigration(migration *Migration) {
	m.Migrations[migration.Version] = migration
	index := 0
	for index < len(m.Versions) {
		if m.Versions[index] > migration.Version {
			break
		}
		index++
	}

	m.Versions = append(m.Versions, migration.Version)
	copy(m.Versions[index+1:], m.Versions[index:])
	m.Versions[index] = migration.Version
}

func (m *Migrator) Up(ctx context.Context) error {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	for _, v := range m.Versions {
		migration := m.Migrations[v]

		if migration.done {
			continue
		}

		fmt.Println("Running migration ", migration.Version)
		if err := migration.Up(tx); err != nil {
			tx.Rollback()
			return fmt.Errorf("Error trying to run migration %s. %w", migration.Version, err)
		}

		if _, err := tx.Exec("Insert INTO schema_migrations VALUES($1)", migration.Version); err != nil {
			tx.Rollback()
			return fmt.Errorf("Error trying to insert version %s in the schema_migrations table. %w", migration.Version, err)
		}
		fmt.Println("Finished running migration", migration.Version)
	}
	tx.Commit()
	return nil
}

func (m *Migrator) Down(ctx context.Context, step int) error {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	count := 0
	for _, v := range reverse(m.Versions) {
		if step > 0 && count == step {
			fmt.Println("caiu no if")
			break
		}

		migration := m.Migrations[v]
		if !migration.done {
			continue
		}

		fmt.Println("Reverting Migration", migration.Version)
		if err := migration.Down(tx); err != nil {
			tx.Rollback()
			return fmt.Errorf("Error trying to revert migration %s. %w", migration.Version, err)
		}

		if _, err := tx.Exec("DELETE FROM schema_migrations WHERE version = $1", migration.Version); err != nil {
			tx.Rollback()
			return fmt.Errorf("Error trying to delete version %s in the schema_migrations table. %w", migration.Version, err)
		}

		fmt.Println("Finished reverting migration", migration.Version)
		count++
	}
	tx.Commit()
	return nil
}

func reverse(args []string) []string {
	rev := make([]string, len(args))
	n := 0
	for i := len(args) - 1; i >= 0; i-- {
		rev[n] = args[i]
		n++
	}
	return rev
}
