package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SGEK-code/url-shortener.git/internal/config"
	"github.com/SGEK-code/url-shortener.git/internal/repository/inmemory"
	"github.com/SGEK-code/url-shortener.git/internal/service/shortener"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainFormat(t *testing.T) {
	meowUrl := "http://meow.ru"
	posTestEx, err := shortener.ShortenURLNoSave(meowUrl)
	require.NoError(t, err)

	cfg := &config.Config{ListenAddr: "notUsed", BaseURL: "http://testBaseUrl.ru"}

	tests := []struct {
		method              string
		contentTypeSent     string
		code                int
		contentTypeExpected string
		body                string
	}{
		{
			method:              http.MethodGet,
			contentTypeSent:     "text/plain",
			code:                http.StatusBadRequest,
			contentTypeExpected: "text/plain; charset=utf-8",
			body:                http.StatusText(http.StatusMethodNotAllowed) + "\n",
		},
		{
			method:              http.MethodPost,
			contentTypeSent:     "notTextPlain",
			contentTypeExpected: "text/plain; charset=utf-8",
			code:                http.StatusBadRequest,
			body:                http.StatusText(http.StatusUnsupportedMediaType) + "\n",
		},
		{
			method:              http.MethodPost,
			contentTypeSent:     "text/plain",
			contentTypeExpected: "text/plain",
			code:                http.StatusCreated,
			body:                cfg.BaseURL + "/" + posTestEx,
		},
	}

	repo := inmemory.NewInMemoryResourceRepo()
	shortener := shortener.NewResourceService(repo)
	handler := NewShortenerHandler(shortener, cfg)

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(test.method, "/", strings.NewReader(meowUrl))
			req.Header.Set("Content-Type", test.contentTypeSent)

			handler.Main(recorder, req)

			response := recorder.Result()
			body, err := io.ReadAll(response.Body)
			require.NoError(t, err)
			err = response.Body.Close()
			require.NoError(t, err)

			assert.Equal(t, test.code, response.StatusCode)
			assert.Equal(t, test.contentTypeExpected, response.Header.Get("Content-Type"))
			assert.Equal(t, test.body, string(body))
		})
	}
}

func TestReturnUrlFormat(t *testing.T) {
	repo := inmemory.NewInMemoryResourceRepo()
	shortener := shortener.NewResourceService(repo)
	cfg := &config.Config{ListenAddr: "notUsed", BaseURL: "notUsed"}
	handler := NewShortenerHandler(shortener, cfg)

	meowUrl := "http://meow.ru"
	posTestEx, err := shortener.ShortenURL(meowUrl)
	require.NoError(t, err)
	value, err := shortener.GetUrl(posTestEx)
	require.NoError(t, err)
	require.Equal(t, meowUrl, value)

	tests := []struct {
		method           string
		targetAddr       string
		code             int
		expectedLocation string
	}{
		{
			method:           http.MethodPost,
			targetAddr:       "/" + posTestEx,
			code:             http.StatusBadRequest,
			expectedLocation: "",
		},
		{
			method:           http.MethodGet,
			targetAddr:       "/badCheck",
			code:             http.StatusBadRequest,
			expectedLocation: "",
		},
		{
			method:           http.MethodGet,
			targetAddr:       "/" + posTestEx,
			code:             http.StatusTemporaryRedirect,
			expectedLocation: meowUrl,
		},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(test.method, test.targetAddr, nil)

			handler.ReturnUrl(recorder, req)

			response := recorder.Result()
			err = response.Body.Close()
			require.NoError(t, err)

			assert.Equal(t, test.code, response.StatusCode)
			assert.Equal(t, test.expectedLocation, response.Header.Get("Location"))
		})
	}
}
