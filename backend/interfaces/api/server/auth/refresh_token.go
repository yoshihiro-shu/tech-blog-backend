package auth

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type RefreshToken struct {
	UserId    int       `json:"user_id"`
	JwtId     string    `json:"jwt_id"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewRefreshToken(userId int) *RefreshToken {
	return &RefreshToken{
		UserId:    userId,
		JwtId:     GenerateToken(),
		ExpiredAt: time.Now().Add(conf.RefreshToken.Expires),
	}
}

func (r RefreshToken) JwtToken() string {
	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(r.UserId),
		"jwt_id":  r.JwtId,
		"exp":     r.ExpiredAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(conf.AccessToken.SecretKey))

	return tokenString
}
