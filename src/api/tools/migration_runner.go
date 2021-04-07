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

	migrateCmd.AddCommand(migrateCreateCmd, migrateUpCmd)

	rootCmd.AddCommand(migrateCmd)
}

func main() {
	// var args struct {
	// 	ConnString string `short:"c" long:"connstring" required:"true" name:"Connection String"`
	// }

	// _, err := flags.Parse(&args)
	// if err != nil {
	// 	os.Stderr.WriteString(err.Error() + "\n")
	// 	os.Exit(1)
	// }

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

		db := database.NewConnection(connString)
		migrator, err := migrations.InitMigrator(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Up(cmd.Context())
		if err != nil {
			fmt.Println("Unable to run `up` migrations")
		}
	},
}
