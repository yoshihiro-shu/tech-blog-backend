package auth

import (
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

type authConfig struct {
	AccessToken  config.AuthToken
	RefreshToken config.AuthToken
}
