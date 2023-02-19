package auth

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateBcryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func IsVerifyPassword(textConplainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(textConplainPassword))
	return err == nil
}

func GenerateToken() string {
	return uuid.Must(uuid.NewRandom()).String()
}

func getTokenFromHeader(r *http.Request) (*jwt.Token, error) {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	jwtToken, err := verifyToken(token)
	if err != nil {
		return nil, err
	}
	return jwtToken, nil
}
