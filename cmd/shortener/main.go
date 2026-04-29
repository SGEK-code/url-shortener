package main

import (
	"log"
	"net/http"

	"github.com/SGEK-code/url-shortener.git/internal/config"
	"github.com/SGEK-code/url-shortener.git/internal/repository/inmemory"
	"github.com/SGEK-code/url-shortener.git/internal/router"
)

func run(cfg *config.Config) error {
	repo := inmemory.NewInMemoryResourceRepo()

	srv := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: router.SetupRouter(repo, cfg),
	}

	log.Printf("Server starting on %s", srv.Addr)
	err := srv.ListenAndServe()

	return err
}

func main() {
	cfg := config.ParseConfig()
	log.Fatal(run(cfg))
}
