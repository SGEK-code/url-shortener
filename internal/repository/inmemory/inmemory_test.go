package inmemory

import (
	"testing"

	"github.com/SGEK-code/url-shortener.git/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryResourceRepo(t *testing.T) {
	repo := NewInMemoryResourceRepo()
	require.NotNil(t, repo)

	tests := []struct {
		create        *model.Resource
		provide       *model.Resource
		expected      *model.Resource
		expectedError error
	}{
		{
			create: &model.Resource{
				Url:  "testUrl1",
				Hash: "testHash1",
			},
			provide: &model.Resource{
				Url:  "testUrl1",
				Hash: "testHash1",
			},
			expected: &model.Resource{
				Url:  "testUrl1",
				Hash: "testHash1",
			},
			expectedError: nil,
		},
		{
			create: &model.Resource{
				Url:  "testUrlError",
				Hash: "testHashError",
			},
			provide: &model.Resource{
				Url:  "testUrl",
				Hash: "testHash",
			},
			expected:      nil,
			expectedError: ErrNoResultFound,
		},
	}

	for _, test := range tests {
		t.Run(test.create.Url, func(t *testing.T) {
			err := repo.Create(test.create)
			assert.NoError(t, err)

			result, err := repo.GetByUrl(test.provide.Url)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.expectedError, err)

			result, err = repo.GetByHash(test.provide.Hash)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
