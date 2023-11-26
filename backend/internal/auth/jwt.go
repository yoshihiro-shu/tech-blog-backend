package auth

import (
	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/config"
)

// TODO configで設定できるようにする
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
