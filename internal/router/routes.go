package router

import (
	"net/http"

	"github.com/SGEK-code/url-shortener.git/internal/handler"
)

func addRoutes(
	mux *http.ServeMux,
) {
	mux.Handle("/", http.HandlerFunc(handler.MainHandler))
}
