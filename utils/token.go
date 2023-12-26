package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const tokenLifespan = 1 * time.Hour
const tokenSecret = "6a3USUhktLyrV8Z-95ecO9eSwFmKU3M9O4M4z5Res2E="

// GenerateToken generates a token with a lifespan of 1 hour and makes a claim about the username
func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(tokenLifespan)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tokenSecret))
}

func TokenValid(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// GenerateRandomKey generates a random 32-Bit Key (e.g. as token secret)
func GenerateRandomKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)

	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}
