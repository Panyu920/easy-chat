package token

import (
	"fmt"

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
func VerifyToken(key, token string) (string, error) {
	parseFunc := func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return []byte(key), nil
	}
	tokenJwt, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, parseFunc)
	if err != nil {
		return "", err
	}

	if !tokenJwt.Valid {
		return "", fmt.Errorf("token is not valid")
	}
	claims, ok := tokenJwt.Claims.(*jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("unexpected claims type: %v", tokenJwt.Claims)
	}
	userID, ok := (*claims)[UserID]
	if !ok {
		return "", fmt.Errorf("unexpected user id type: %v", (*claims)[UserID])
	}

	return userID.(string), nil
}
