package encrypt

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func MD5(str string) string {
	return fmt.Sprintf("%x", md5.New().Sum([]byte(str)))
}

func GneratePasswordHash(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwd), nil
}

func ValidatePasswordHash(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
