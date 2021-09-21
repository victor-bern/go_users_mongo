package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}
