package shortener

import (
	"errors"
	"fmt"
	"hash/crc32"
)

var urlToHash = make(map[string]string)
var HashToUrl = make(map[string]string)
var ErrNoResultFound = errors.New("no results found")

func ShortenURL(url string) (string, error) {
	checksumInt := crc32.ChecksumIEEE([]byte(url))
	checksum := fmt.Sprintf("%08X", checksumInt)
	urlToHash[url] = checksum
	HashToUrl[checksum] = url

	return checksum, nil
}

func GetUrl(checksum string) (string, error) {
	if val, ok := HashToUrl[checksum]; !ok {
		return "", ErrNoResultFound
	} else {
		return val, nil
	}

}
