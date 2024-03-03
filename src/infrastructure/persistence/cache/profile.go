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

func (c *profileCacheAdaptor) GetResume() ([]byte, error) {
	var res []byte
	err := c.client.GET(ResumeKey(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *profileCacheAdaptor) SetResume(resume []byte) error {
	return c.client.SET(ResumeKey(), resume)
}
