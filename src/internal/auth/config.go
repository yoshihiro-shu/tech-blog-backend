package auth

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/config"
)

type authConfig struct {
	AccessToken  config.AuthToken
	RefreshToken config.AuthToken
}
