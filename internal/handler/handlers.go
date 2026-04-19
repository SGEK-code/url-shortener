package handler

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/SGEK-code/url-shortener.git/internal/service/shortener"
)

type ShortenerHandler struct {
	srs *shortener.ResourceService
}

func NewShortenerHandler(srs *shortener.ResourceService) *ShortenerHandler {
	return &ShortenerHandler{srs: srs}
}

func (h *ShortenerHandler) Main(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusBadRequest)
		return
	}
	if r.Header.Get("Content-Type") != "text/plain" {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	shortened, err := h.srs.ShortenURL(string(body))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := "http://" + r.Host + "/" + shortened
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result))
}

func (h *ShortenerHandler) ReturnUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	checksum := strings.TrimPrefix(r.URL.Path, "/")
	checksum = strings.TrimSpace(checksum)

	url, err := h.srs.GetUrl(checksum)
	if err != nil {
		if errors.Is(err, shortener.ErrNoResultFound) {
			http.Error(w, "unregistered url", http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
