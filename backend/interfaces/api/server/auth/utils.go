package auth

import "golang.org/x/crypto/bcrypt"

func GenerateBcryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func IsVerifyPassword(password1, password2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	if err != nil {
		return false
	}
	return true
}
