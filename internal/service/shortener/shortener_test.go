package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
