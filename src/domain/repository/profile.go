package repository

type ProfileCacheRepository interface {
	GetResume(profile []byte) error
	SetResume(profile []byte) error
}
