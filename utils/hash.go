package utils

import (
	"crypto/sha1"
	"fmt"
)

func HashFileContent(content []byte) string {
	hash := sha1.New()
	hash.Write(content)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
