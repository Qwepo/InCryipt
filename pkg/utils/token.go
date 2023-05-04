package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomToken() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(randomBytes)[:32], nil
}
