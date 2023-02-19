package auth

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AccessToken struct {
	UserId    int
	ExpiredAt time.Time
}

func NewAccessToken(userId int) *AccessToken {
	return &AccessToken{
		UserId:    userId,
		ExpiredAt: time.Now().Add(conf.AccessToken.Expires),
	}
}

func (a AccessToken) JwtToken() string {
	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(a.UserId),
		"exp":     a.ExpiredAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(conf.AccessToken.SecretKey))

	return tokenString
}
