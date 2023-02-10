package helper

import "crypto/sha1"

func GenerateCacheKey(url string) string {
	encodedBytes := sha1.Sum([]byte("ecoindex_" + url))

	return string(encodedBytes[:])
}
