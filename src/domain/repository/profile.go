package repository

type ProfileCacheRepository interface {
	GetResume() ([]byte, error)
	SetResume(resume []byte) error
}
