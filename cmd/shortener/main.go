package main

import (
	"log"
	"net/http"

	"github.com/SGEK-code/url-shortener.git/internal/config"
	"github.com/SGEK-code/url-shortener.git/internal/logger"
	"github.com/SGEK-code/url-shortener.git/internal/repository/inmemory"
	"github.com/SGEK-code/url-shortener.git/internal/router"
)

func run(cfg *config.Config) error {
	err := logger.Initialize(cfg.LogLevel)
	if err != nil {
		return err
	}

	repo := inmemory.NewInMemoryResourceRepo()

	srv := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: router.SetupRouter(repo, cfg),
	}

	log.Printf("Server starting on %s, base URL is %s", srv.Addr, cfg.BaseURL)
	err = srv.ListenAndServe()

	return err
}

func main() {
	cfg := config.ParseConfig()
	log.Fatal(run(cfg))
}
