package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	expHour   = 72
	secretKey = "secret"
	UserKey   = "userID"
)

func CreateAccessToken(id int) string {
	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(id),
		"exp":     time.Now().Add(time.Hour * expHour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
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
