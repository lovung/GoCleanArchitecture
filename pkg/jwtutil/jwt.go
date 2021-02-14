package jwtutil

import (
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	once         sync.Once
	jwtSecret    []byte
	aExpDuration time.Duration
	rExpDuration time.Duration
)

// Error constant for jwt auth
var (
	ErrInvalidAlg = fmt.Errorf("invalid alg")
)

// InitJWTSession initializes variables for JWT Session
func InitJWTSession(secret string, aExpDur, rExpDur time.Duration) {
	once.Do(func() {
		jwtSecret = []byte(secret)
		aExpDuration = aExpDur
		rExpDuration = rExpDur
	})
}

// GenerateToken generates JWT for authentication
func GenerateToken(claims AuthClaims) (aTokenStr, rTokenStr string, err error) {
	aTokenStr, err = generateJWT(claims, aExpDuration)
	if err != nil {
		return "", "", err
	}

	rTokenStr, err = generateJWT(claims, rExpDuration)

	return aTokenStr, rTokenStr, err
}

// TODO: consider using Redis to revoke token per user
func generateJWT(claims AuthClaims, expDur time.Duration) (tokenStr string, err error) {
	claims.StandardClaims = jwt.StandardClaims{
		IssuedAt:  claims.IssueTime.Unix(),
		ExpiresAt: claims.IssueTime.Add(expDur).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(jwtSecret)

	return tokenStr, err
}

// VerifyToken validates JWT and extract userId and officeID
func VerifyToken(tokenStr string) (AuthClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidAlg
		}
		return jwtSecret, nil
	})

	if err != nil {
		return AuthClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return AuthClaims{}, err
	}

	authClaims := AuthClaims{
		UserID: uint64(claims["user_id"].(float64)),
	}

	return authClaims, nil
}
