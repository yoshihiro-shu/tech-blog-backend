package repository

type RefreshTokenRepository interface {
	Create(userId int, jwtId string) error
}
