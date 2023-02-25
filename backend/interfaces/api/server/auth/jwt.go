package auth

import (
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
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
