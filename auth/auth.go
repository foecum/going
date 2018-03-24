package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtUser contains token claims map values
type JwtUser struct {
	Name string
	ID   int64
}

// New  creates a new auth token
func New(claims jwt.Claims) (string, error) {
	secret := os.Getenv("JWT_TOKEN_SECRET")
	if secret == "" {
		secret = "secret"
		// TODO: Uncomment this code
		// return "", fmt.Errorf("\"JWT_TOKEN_SECRET\" enviroment variable not set")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = claims

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateSHA256Hash ...
func GenerateSHA256Hash(raw string) string {
	hasher := sha256.New224()
	hasher.Write([]byte(raw))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
