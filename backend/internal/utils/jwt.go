package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserUUID string `json:"userUuid"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func SignAccessToken(secret, issuer, userUUID, email string, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{
		UserUUID: userUUID,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secret))
}

func ParseToken(secret, tokenStr string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	c, ok := t.Claims.(*Claims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return c, nil
}
