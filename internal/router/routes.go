package router

import (
	"net/http"

	"github.com/SGEK-code/url-shortener.git/internal/handler"
)

func addRoutes(
	mux *http.ServeMux,
	shortenerHandler *handler.ShortenerHandler,
) {
	mux.HandleFunc("/", shortenerHandler.Main)
	mux.HandleFunc("/{checksum}", shortenerHandler.ReturnUrl)
}
