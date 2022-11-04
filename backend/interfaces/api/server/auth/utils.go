package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateBcryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func IsVerifyPassword(textConplainPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(textConplainPassword))
	if err != nil {
		return err
	}
	return nil
}
