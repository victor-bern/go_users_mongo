package helpers

import "golang.org/x/crypto/bcrypt"

func GenerateHash(pass []byte) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword(pass, 8)
	if err != nil {
		return nil, err
	}

	return hash, nil

}
