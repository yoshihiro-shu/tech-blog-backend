package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

const (
	UserKey = "userID"
)

var (
	conf authConfig
)

func Init(accessTokenConf, refreshTokenConf config.AuthToken) {
	conf = authConfig{
		AccessToken:  accessTokenConf,
		RefreshToken: refreshTokenConf,
	}
}

func CreateAccessToken(id int) string {
	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(id),
		"exp":     time.Now().Add(conf.AccessToken.Expires).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(conf.AccessToken.SecretKey))

	return tokenString
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.AccessToken.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
