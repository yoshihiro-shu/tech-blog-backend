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

func VerifyRefeshToken(tokenString string) (*RefreshToken, error) {
	jwtToken, err := verifyToken(tokenString, conf.RefreshToken.SecretKey)
	if err != nil {
		return nil, err
	}

	claims := jwtToken.Claims.(jwt.MapClaims)

	jwt_id := claims["jwt_id"]
	if jwt_id == nil {
		return nil, ErrJwtIdIsMissingAtRefreshToken
	}

	exp := claims["exp"]
	if jwt_id == nil {
		return nil, ErrExpiresIsMissingAtRefreshToken
	}

	user_id := claims["user_id"]
	if user_id == nil {
		return nil, ErrUserIdIsMissingAtRefreshToken
	}

	userId, err := strconv.Atoi(user_id.(string))
	if err != nil {
		return nil, err
	}

	jwtId := jwt_id.(string)

	expfloat := exp.(float64)
	expires := time.Unix(int64(expfloat), 0)

	return &RefreshToken{
		UserId:    userId,
		JwtId:     jwtId,
		ExpiredAt: expires,
	}, nil
}

func (r RefreshToken) JwtToken() string {
	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(r.UserId),
		"jwt_id":  r.JwtId,
		"exp":     r.ExpiredAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(conf.AccessToken.SecretKey))
	if err != nil {
		panic(err)
	}

	return tokenString
}
