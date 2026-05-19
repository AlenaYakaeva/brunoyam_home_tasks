package auth

import (
	"ToDoList/internal/service/errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret-word")

type TokenClaims struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(uid string) (string, error) {
	claims := TokenClaims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(uid string) (string, error) {
	claims := TokenClaims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&TokenClaims{},
		func(t *jwt.Token) (any, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return "", errors.ErrInvalidJWT
	}
	return claims.UID, nil
}
