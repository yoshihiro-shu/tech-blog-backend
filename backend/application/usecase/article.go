package usecase

import (
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/repository"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/cache"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pager"
)

type ArticleUseCase interface {
	Create(title, content string, userId, categoryId int) (*model.Article, error)
	FindByID(id int) (*model.Article, error)
	GetArticles(articles *[]model.Article, limit, offset int) error
	GetPager(currentPage, offset int) (*pager.Pager, error)
	Update(id int, title, content string) (*model.Article, error)
	Delete(id int) error
}

type articleUseCase struct {
	articleRepo repository.ArticleRepository
	cacheRepo   cache.RedisClient
}

func NewArticleUseCase(articleRepo repository.ArticleRepository, cacheRepo cache.RedisClient) ArticleUseCase {
	return &articleUseCase{
		articleRepo: articleRepo,
		cacheRepo:   cacheRepo}
}

func (au *articleUseCase) Create(title, content string, userId, categoryId int) (*model.Article, error) {
	article := &model.Article{
		Title:      title,
		Content:    content,
		UserId:     userId,
		CategoryId: categoryId,
	}

	createdArticle, err := au.articleRepo.Create(article)
	if err != nil {
		return nil, err
	}

	return createdArticle, nil
}

func (au *articleUseCase) FindByID(id int) (*model.Article, error) {
	var article model.Article
	if err := au.cacheRepo.GET(cache.GetArticleByIdKey(id), &article); err == nil {
		return &article, nil
	}

	err := au.articleRepo.FindByID(&article, id)
	if err != nil {
		return nil, err
	}

	if err := au.cacheRepo.SET(cache.GetArticleByIdKey(id), article); err != nil {
		// logのみを出力するエラーハンドリングに変えたい。
		return nil, err
	}
	return &article, nil
}

func (au *articleUseCase) GetArticles(articles *[]model.Article, limit, offset int) error {
	return au.articleRepo.GetArticles(articles, limit, offset)
}

func (au *articleUseCase) GetPager(currentPage, offset int) (*pager.Pager, error) {
	var a model.Article

	numOfArticles, err := au.articleRepo.GetPager(&a)
	if err != nil {
		return nil, err
	}

	pager := pager.New(currentPage)
	pager.SetLastPage(offset, numOfArticles)

	return pager, nil
}

func (au *articleUseCase) Update(id int, title, content string) (*model.Article, error) {
	return &model.Article{}, nil
}

func (au *articleUseCase) Delete(id int) error {
	return nil
}
