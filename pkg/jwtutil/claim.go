package jwtutil

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// const key for middleware
const (
	JWTClaimsKey = "jwtClaims"
)

// AuthClaims the claim for authentication
type AuthClaims struct {
	jwt.StandardClaims
	UserID    uint64    `json:"user_id,omitempty"`
	IssueTime time.Time `json:"-"`
}

// JWTClaims is the information inside Claims
type JWTClaims struct {
	UserID uint64
}

// ExtractClaims extracts AuthClaims from context
func ExtractClaims(ctx context.Context) (JWTClaims, error) {
	claims := ctx.Value(JWTClaimsKey)
	if claims == nil {
		// We will never have this error because AuthenticateAccess raises error before accessing here.
		// It means jwtClaims always exists that AuthenticateAccess doesn't raise error.
		return JWTClaims{}, errors.New("failed to extract jwtClaims")
	}

	return claims.(JWTClaims), nil
}
