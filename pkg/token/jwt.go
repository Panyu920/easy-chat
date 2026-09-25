package token

import (
	"github.com/golang-jwt/jwt"
)

const (
	UserID = "user_id"
)

func GenerateToken(key, userID string, now, duration int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims[UserID] = userID
	claims["exp"] = now + duration
	claims["iat"] = now

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}
