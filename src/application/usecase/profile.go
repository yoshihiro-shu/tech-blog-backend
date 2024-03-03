package usecase

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/repository"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/github_api"
)

type ProfileUseCase interface {
	GetResume() ([]byte, error)
}
type profileUseCase struct {
	token            string
	cacheProfileRepo repository.ProfileCacheRepository
}

func NewProfileUseCase(token string, cacheProfileRepo repository.ProfileCacheRepository) ProfileUseCase {
	return &profileUseCase{
		token:            token,
		cacheProfileRepo: cacheProfileRepo,
	}
}

func (u *profileUseCase) GetResume() ([]byte, error) {
	var res []byte
	if res, err := u.cacheProfileRepo.GetResume(); err == nil {
		return res, nil
	}

	client := github_api.NewClient("yoshihiro-shu", "Resume", u.token)
	res, err := client.GetRepositoryContent("README.md")
	if err != nil {
		return nil, err
	}

	go u.cacheProfileRepo.SetResume(res)

	return res, nil
}
