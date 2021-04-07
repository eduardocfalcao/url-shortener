package main

import (
	"github.com/eduardocfalcao/url-shortener/src/api/cmd"
)

func main() {
	cmd.StartHttpServer(":8080")
}
