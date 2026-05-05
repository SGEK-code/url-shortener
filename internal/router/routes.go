package router

import (
	"net/http"

	"github.com/SGEK-code/url-shortener.git/internal/handler"
	"github.com/SGEK-code/url-shortener.git/internal/logger"
	"github.com/go-chi/chi/v5"
)

func addRoutes(
	mux *chi.Mux,
	shortenerHandler *handler.ShortenerHandler,
) {
	mux.Route("/", func(r chi.Router) {
		r.Handle("/", logger.RequestLogger(http.HandlerFunc(shortenerHandler.Main)))
		r.Route("/{checksum}", func(r chi.Router) {
			r.Handle("/", logger.RequestLogger(http.HandlerFunc(shortenerHandler.ReturnUrl)))
		})
	})
}
