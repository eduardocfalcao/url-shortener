package main

import (
	"flag"
	"log"

	"github.com/eduardocfalcao/url-shortener/src/api/cmd"
	"github.com/eduardocfalcao/url-shortener/src/api/config"
)

func main() {
	envFile := flag.String("env-file", "", "Specify whether environment should be loaded from file or not")
	flag.Parse()

	if *envFile != "" {
		err := config.SetupConfigFile(".", *envFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	config := config.GetConfiguration()

	cmd.StartHttpServer(":8080", config)
}
