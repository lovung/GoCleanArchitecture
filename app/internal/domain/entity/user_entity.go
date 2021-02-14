package entity

import (
	"crypto/rand"
	"time"

	"github.com/lovung/GoCleanArchitecture/pkg/hasher"

	"github.com/oklog/ulid"
)

// User entity
type User struct {
	ID       string
	Username string
	Password string
}

// HashPassword user.Password <- hash(user.Password)
func (e *User) HashPassword() error {
	hashed, err := hasher.HashPassword(e.Password)
	if err != nil {
		return err
	}
	e.Password = hashed
	return nil
}

// GenID to generate ID
func (e *User) GenID() (err error) {
	entropy := ulid.Monotonic(rand.Reader, 0)
	id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
	if err != nil {
		return err
	}
	e.ID = id.String()
	return err
}
