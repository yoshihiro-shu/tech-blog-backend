package cache

import "github.com/yoshihiro-shu/tech-blog-backend/src/domain/repository"

type profileCacheAdaptor struct {
	client RedisClient
}

func NewProfileCacheAdaptor(c RedisClient) repository.ProfileCacheRepository {
	return &profileCacheAdaptor{
		client: c,
	}
}

func (c *profileCacheAdaptor) GetResume(resume []byte) error {
	return c.client.GET(ResumeKey(), &resume)
}

func (c *profileCacheAdaptor) SetResume(resume []byte) error {
	return c.client.SET(ResumeKey(), resume)
}
