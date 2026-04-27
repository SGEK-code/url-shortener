package shortener

import (
	"testing"

	"github.com/SGEK-code/url-shortener.git/internal/repository/inmemory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortenURLNoSave(t *testing.T) {
	// testcases for crc32.ChecksumIEEE
	tests := []struct {
		url          string
		expectedHash string
		expectedErr  error
	}{
		{
			url:          "http://meow.ru",
			expectedHash: "9FB40793",
			expectedErr:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			result, err := ShortenURLNoSave(test.url)
			assert.Equal(t, test.expectedHash, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestShortenURLGetURL(t *testing.T) {
	tests := []struct {
		urlReq        string
		hashReq       string
		urlExp        string
		hashExp       string
		expectedError error
	}{
		{
			urlReq:        "http://meow.ru",
			hashReq:       shortenURLNoSaveNoErr("http://meow.ru"),
			urlExp:        "http://meow.ru",
			hashExp:       shortenURLNoSaveNoErr("http://meow.ru"),
			expectedError: nil,
		},
		{
			urlReq:        "http://UnknownUrl.com",
			hashReq:       "UnknownH",
			urlExp:        "",
			hashExp:       shortenURLNoSaveNoErr("http://UnknownUrl.com"),
			expectedError: ErrNoResultFound,
		},
	}

	for _, test := range tests {
		t.Run(test.urlReq, func(t *testing.T) {
			repo := inmemory.NewInMemoryResourceRepo()
			require.NotNil(t, repo)

			service := NewResourceService(repo)
			require.NotNil(t, service)

			resultHash, err := service.ShortenURL(test.urlReq)
			require.NoError(t, err)
			assert.Equal(t, test.hashExp, resultHash)

			resultUrl, err := service.GetUrl(test.hashReq)
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.urlExp, resultUrl)
		})
	}
}

func shortenURLNoSaveNoErr(url string) string {
	result, err := ShortenURLNoSave(url)
	if err != nil {
		panic(err)
	}
	return result
}
