package main

import (
	"log"

	"github.com/SGEK-code/url-shortener.git/internal/router"
)

func run() error {
	return router.StartServer("localhost:8080")
}

func main() {
	log.Fatal(run())
}
