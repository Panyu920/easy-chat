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
	return fmt.Sprintf("%x", pwd), nil
}

func ValidatePasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
