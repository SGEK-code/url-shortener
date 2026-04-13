package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/SGEK-code/url-shortener.git/internal/service/shortener"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
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

	shortened, err := shortener.ShortenURL(string(body))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := "http://" + r.Host + "/" + shortened
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result))
}

func ReturnUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	checksum := r.PathValue("checksum")

	url, err := shortener.GetUrl(checksum)
	if err != nil {
		if errors.Is(err, shortener.ErrNoResultFound) {
			http.Error(w, "unregistered url", http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte("Location: " + url))
}
