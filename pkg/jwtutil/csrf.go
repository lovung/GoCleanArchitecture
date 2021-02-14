package jwtutil

import (
	"crypto/rand"
	"fmt"
)

const (
	csrfTokenLength = 32
)

// GenerateCSRFToken generates csrf token with random string
func GenerateCSRFToken() (string, error) {
	b := make([]byte, csrfTokenLength)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
