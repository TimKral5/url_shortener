package util

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func GenerateSHA256Hex(source string) string {
	hash := sha256.New()
	hash.Write([]byte(source))
	hashStr := strings.ToUpper(hex.EncodeToString(hash.Sum(nil))[:10])

	return hashStr
}
