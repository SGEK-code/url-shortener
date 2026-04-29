package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SGEK-code/url-shortener.git/internal/config"
	"github.com/SGEK-code/url-shortener.git/internal/repository/inmemory"
	"github.com/SGEK-code/url-shortener.git/internal/router"
	"github.com/SGEK-code/url-shortener.git/internal/service/shortener"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainFormat(t *testing.T) {
	meowUrl := "http://meow.ru"
	posTestEx, err := shortener.ShortenURLNoSave(meowUrl)
	require.NoError(t, err)

	repo := inmemory.NewInMemoryResourceRepo()
	cfg := &config.Config{ListenAddr: "notUsed", BaseURL: "http://testBaseUrl.ru"}
	srv := httptest.NewServer(router.SetupRouter(repo, cfg))
	defer srv.Close()

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

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			req := resty.New().R()
			req.Method = test.method
			req.URL = srv.URL
			req.SetBody(strings.NewReader(meowUrl))
			req.Header.Set("Content-Type", test.contentTypeSent)

			resp, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")

			assert.Equal(t, test.code, resp.StatusCode())
			assert.Equal(t, test.contentTypeExpected, resp.Header().Get("Content-Type"))
			assert.Equal(t, test.body, string(resp.Body()))
		})
	}
}

func TestReturnUrlFormat(t *testing.T) {
	meowUrl := "http://meow.ru"
	posTestEx, err := shortener.ShortenURLNoSave(meowUrl)
	require.NoError(t, err)

	repo := inmemory.NewInMemoryResourceRepo()
	cfg := &config.Config{ListenAddr: "notUsed", BaseURL: "notUsed"}
	srv := httptest.NewServer(router.SetupRouter(repo, cfg))
	defer srv.Close()

	reqProperties := struct {
		contentType string
		body        string
	}{
		contentType: "text/plain",
		body:        meowUrl,
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", reqProperties.contentType).
		SetBody(strings.NewReader(reqProperties.body)).
		Post(srv.URL)
	require.Equal(t, http.StatusCreated, resp.StatusCode())

	tests := []struct {
		method           string
		targetAddr       string
		code             int
		expectedLocation string
	}{
		{
			method:           http.MethodPost,
			targetAddr:       posTestEx,
			code:             http.StatusBadRequest,
			expectedLocation: "",
		},
		{
			method:           http.MethodGet,
			targetAddr:       "badCheck",
			code:             http.StatusBadRequest,
			expectedLocation: "",
		},
		{
			method:           http.MethodGet,
			targetAddr:       posTestEx,
			code:             http.StatusTemporaryRedirect,
			expectedLocation: meowUrl,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s/%s", test.method, test.targetAddr), func(t *testing.T) {
			// Здесь resty дает кастомную ошибку auto redirect is disabled при
			// отправке запроса, т.ч. используем стандартный клиент
			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}

			req, err := http.NewRequest(test.method, srv.URL+"/"+test.targetAddr, nil)
			require.NoError(t, err, "failed to create request")

			resp, err := client.Do(req)
			require.NoError(t, err, "error making HTTP request")
			defer resp.Body.Close()

			assert.Equal(t, test.code, resp.StatusCode,
				"Неожиданный статус-код. Запрос: %s %s", test.method, test.targetAddr)
			assert.Equal(t, test.expectedLocation, resp.Header.Get("Location"),
				"Неожиданный Location. Запрос: %s %s", test.method, test.targetAddr)
		})
	}
}
