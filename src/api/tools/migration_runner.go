package main

import (
	"log"
	"os"

	"github.com/eduardocfalcao/url-shortener/src/api/migrations"
)

func main() {
	// var args struct {
	// 	ConnString string `short:"c" long:"connstring" required:"true" name:"Connection String"`
	// }

	// _, err := flags.Parse(&args)
	// if err != nil {
	// 	os.Stderr.WriteString(err.Error() + "\n")
	// 	os.Exit(1)
	// }

	if err := migrations.Create("Init_Database"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
