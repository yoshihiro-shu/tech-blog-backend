package auth

import "errors"

var (
	ErrInvalidPassword                = errors.New("error invalid password")
	ErrUserIdIsMissingAtRefreshToken  = errors.New("failed at get user id from refresh token")
	ErrJwtIdIsMissingAtRefreshToken   = errors.New("failed at get jwt id from refresh token")
	ErrExpiresIsMissingAtRefreshToken = errors.New("failed at get exp from refresh token")
)
