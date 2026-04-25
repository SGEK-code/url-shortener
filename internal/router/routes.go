package router

import (
	"github.com/SGEK-code/url-shortener.git/internal/handler"
	"github.com/go-chi/chi/v5"
)

func addRoutes(
	mux *chi.Mux,
	shortenerHandler *handler.ShortenerHandler,
) {
	mux.Route("/", func(r chi.Router) {
		r.Post("/", shortenerHandler.Main)
		r.Route("/{checksum}", func(r chi.Router) {
			r.Get("/", shortenerHandler.ReturnUrl)
		})
	})
}
