package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Secret key (must match middleware)
var secretKey = []byte("supersecret")

// Expected issuer and audience
const (
	issuer   = "my-app"
	audience = "my-audience"
)

// GenerateJWT creates a new token for a given username
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,                             // Store username (e.g., "kunal") as subject
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Expires in 1 hour
		"iss": issuer,                               // ✅ Added issuer
		"aud": audience,                             // ✅ Added audience
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
