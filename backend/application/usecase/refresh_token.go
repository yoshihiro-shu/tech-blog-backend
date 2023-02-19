package usecase

import (
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
)

type AuthUseCase interface {
}

type authUseCase struct {
	refreshTokenRepo repository.RefreshTokenRepository
}

func NewAuthUseCase(rfRepo repository.RefreshTokenRepository) AuthUseCase {
	return &authUseCase{refreshTokenRepo: rfRepo}
}

func (au *authUseCase) CreateRefreshToken(userId int, jwtId string) error {
	return au.refreshTokenRepo.Create(userId, jwtId)
}
