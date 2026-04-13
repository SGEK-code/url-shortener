package shortener

import (
	"fmt"
	"hash/crc32"
)

var urlToHash = make(map[string]string)
var HashToUrl = make(map[string]string)

func ShortenURL(url string) (string, error) {
	checksumInt := crc32.ChecksumIEEE([]byte(url))
	checksum := fmt.Sprintf("%08X", checksumInt)
	urlToHash[url] = checksum
	HashToUrl[checksum] = url

	return checksum, nil
}
