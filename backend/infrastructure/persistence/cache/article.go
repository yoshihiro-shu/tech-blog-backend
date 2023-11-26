package cache

import (
	"github.com/yoshihiro-shu/tech-blog-backend/backend/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/domain/repository"
)

type articleCacheAdaptor struct {
	client RedisClient
}

func NewArticleCacheAdaptor(c RedisClient) repository.ArticleCacheRepository {
	return &articleCacheAdaptor{
		client: c,
	}
}

func (c *articleCacheAdaptor) GetArticleDetailById(article *model.Article, id int) error {
	return c.client.GET(GetArticleByIdKey(id), &article)
}

func (c *articleCacheAdaptor) SetArticleDetailById(article model.Article, id int) error {
	return c.client.SET(GetArticleByIdKey(id), article)
}
