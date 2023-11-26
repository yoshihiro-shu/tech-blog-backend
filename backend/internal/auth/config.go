package auth

import (
	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/config"
)

type authConfig struct {
	AccessToken  config.AuthToken
	RefreshToken config.AuthToken
}
