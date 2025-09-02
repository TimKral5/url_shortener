// Package hash handles utility functions for hashing.
package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// GenerateSHA256Hex generates a SHA256 from source.
func GenerateSHA256Hex(source string) string {
	hash := sha256.New()
	hash.Write([]byte(source))
	hashStr := strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))

	return hashStr
}
