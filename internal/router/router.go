package router

import (
	"net/http"

	"github.com/SGEK-code/url-shortener.git/internal/config"
	"github.com/SGEK-code/url-shortener.git/internal/handler"
	"github.com/SGEK-code/url-shortener.git/internal/service/shortener"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(repo shortener.ResourceRep, cfg *config.Config) http.Handler {
	mux := chi.NewRouter()
	shortener := shortener.NewResourceService(repo)
	handler := handler.NewShortenerHandler(shortener, cfg)
	addRoutes(mux, handler)
	return mux
}
