package router

import (
	"net/http"

	"github.com/SGEK-code/url-shortener.git/internal/handler"
	"github.com/SGEK-code/url-shortener.git/internal/repository/inmemory"
	"github.com/SGEK-code/url-shortener.git/internal/service/shortener"
)

func StartServer(addr string) error {
	mux := http.NewServeMux()
	repo := inmemory.NewInMemoryResourceRepo()
	shortener := shortener.NewResourceService(repo)
	handler := handler.NewShortenerHandler(shortener)
	addRoutes(mux, handler)
	return http.ListenAndServe(addr, mux)
}
