package main

import (
	"log"
	"net/http"

	"github.com/SGEK-code/url-shortener.git/internal/repository/inmemory"
	"github.com/SGEK-code/url-shortener.git/internal/router"
)

func run() error {
	repo := inmemory.NewInMemoryResourceRepo()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.SetupRouter(repo),
	}

	log.Printf("Server starting on %s", srv.Addr)
	err := srv.ListenAndServe()

	return err
}

func main() {
	log.Fatal(run())
}
