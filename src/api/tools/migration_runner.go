package main

import (
	"fmt"
	"log"

	"github.com/eduardocfalcao/url-shortener/src/api/database"
	"github.com/eduardocfalcao/url-shortener/src/api/migrations"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Application Description",
}

func init() {
	migrateCreateCmd.Flags().StringP("name", "n", "", "Name for the migration")

	migrateUpCmd.Flags().StringP("connString", "c", "", "Database connection string to run the migration")

	migrateDownCmd.Flags().StringP("connString", "c", "", "Database connection string to run the migration")
	migrateDownCmd.Flags().IntP("step", "s", 0, "Database connection string to run the migration")

	migrateCmd.AddCommand(migrateCreateCmd, migrateUpCmd, migrateDownCmd)

	rootCmd.AddCommand(migrateCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migrations tool",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new empty migrations file",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Unable to read flag `name`", err.Error())
			return
		}

		if err := migrations.Create(name); err != nil {
			fmt.Println("Unable to create migration", err.Error())
			return
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Run up migrations",
	Run: func(cmd *cobra.Command, args []string) {
		connString, err := cmd.Flags().GetString("connString")
		if err != nil {
			fmt.Println("Unable to read flag `connString`")
			return
		}

		db, err := database.NewConnection(connString)
		if err != nil {
			fmt.Printf("Unable to fetch the database connection: %s", err.Error())
			return
		}

		migrator, err := migrations.InitMigrator(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator", err)
			return
		}

		err = migrator.Up(cmd.Context())
		if err != nil {
			fmt.Println("Unable to run `up` migrations", err)
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback migrations",
	Run: func(cmd *cobra.Command, args []string) {
		connString, err := cmd.Flags().GetString("connString")
		if err != nil {
			fmt.Println("Unable to read flag `connString`")
			return
		}

		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db, err := database.NewConnection(connString)
		if err != nil {
			fmt.Printf("Unable to fetch the database connection: %s", err.Error())
			return
		}

		migrator, err := migrations.InitMigrator(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator", err)
			return
		}

		err = migrator.Down(cmd.Context(), step)
		if err != nil {
			fmt.Println("Unable to run `down` migrations", err)
		}
	},
}
