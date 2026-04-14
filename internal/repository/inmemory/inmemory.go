package inmemory

import (
	"fmt"

	"github.com/SGEK-code/url-shortener.git/internal/model"
	"github.com/SGEK-code/url-shortener.git/internal/service/shortener"
)

type InMemoryResourceRepo struct {
	urlToHash map[string]string
	hashToUrl map[string]string
}

func NewInMemoryResourceRepo() *InMemoryResourceRepo {
	return &InMemoryResourceRepo{
		urlToHash: make(map[string]string),
		hashToUrl: make(map[string]string),
	}
}

var _ shortener.ResourceRep = (*InMemoryResourceRepo)(nil)

func (r *InMemoryResourceRepo) Create(resource *model.Resource) error {
	r.urlToHash[resource.Url] = resource.Hash
	r.hashToUrl[resource.Hash] = resource.Url
	return nil
}

func (r *InMemoryResourceRepo) GetByUrl(url string) (*model.Resource, error) {
	hash, ok := r.urlToHash[url]
	if !ok {
		return nil, fmt.Errorf("hash for url '%s' not found", url)
	}
	return &model.Resource{Url: url, Hash: hash}, nil
}

func (r *InMemoryResourceRepo) GetByHash(hash string) (*model.Resource, error) {
	url, ok := r.hashToUrl[hash]
	if !ok {
		return nil, fmt.Errorf("url for hash '%s' not found", hash)
	}
	return &model.Resource{Url: url, Hash: hash}, nil
}
