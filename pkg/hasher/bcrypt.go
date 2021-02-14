package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

// HashPassword hash password before saving
func bcryptHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

// CheckPasswordHash verify the password
func bcryptCheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}
