package utils

import (
	"crypto/sha256"
	"fmt"
)

func GenerateHash(data string) string {
	// Intialize new hash
	hash := sha256.New()

	// Write data to the hash
	hash.Write([]byte(data))

	// Get hash bytes
	hashBytes := hash.Sum(nil)

	return fmt.Sprintf("%x", hashBytes)
}
