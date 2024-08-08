package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
