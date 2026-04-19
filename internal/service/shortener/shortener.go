package shortener

import (
	"errors"
	"fmt"
	"hash/crc32"

	"github.com/SGEK-code/url-shortener.git/internal/model"
)

var ErrNoResultFound = errors.New("no results found")

type ResourceRep interface {
	Create(resource *model.Resource) error
	GetByUrl(url string) (*model.Resource, error)
	GetByHash(hash string) (*model.Resource, error)
}

type ResourceService struct {
	repo ResourceRep
}

func NewResourceService(repo ResourceRep) *ResourceService {
	return &ResourceService{
		repo: repo,
	}
}

func (s *ResourceService) ShortenURL(url string) (string, error) {
	checksum, err := ShortenURLNoSave(url)
	if err != nil {
		return "", err
	}
	resource := &model.Resource{Url: url, Hash: checksum}
	if err := s.repo.Create(resource); err != nil {
		return "", err
	}

	return checksum, nil
}

func ShortenURLNoSave(url string) (string, error) {
	checksumInt := crc32.ChecksumIEEE([]byte(url))
	checksum := fmt.Sprintf("%08X", checksumInt)
	return checksum, nil
}

func (s *ResourceService) GetUrl(hash string) (string, error) {
	resource, err := s.repo.GetByHash(hash)
	if err != nil {
		return "", ErrNoResultFound
	}

	return resource.Url, nil
}
