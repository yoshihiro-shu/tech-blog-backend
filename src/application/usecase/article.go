package usecase

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/repository"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/pager"
)

type ArticleUseCase interface {
	Create(title, content string, userId, categoryId int) (*model.Article, error)
	FindByID(id int) (*model.Article, error)
	GetArticles(articles *[]model.Article, limit, offset, currentPage int) error
	GetPager(currentPage, offset int) (*pager.Pager, error)
	Update(id int, title, content string) (*model.Article, error)
	Delete(id int) error
}

type articleUseCase struct {
	articleRepo       repository.ArticleRepository
	cacheArticleRepo  repository.ArticleCacheRepository
	cacheArticlesRepo repository.ArticlesCacheRepository
}

func NewArticleUseCase(articleRepo repository.ArticleRepository, cacheArticleRepo repository.ArticleCacheRepository, cacheArticlesRepo repository.ArticlesCacheRepository) ArticleUseCase {
	return &articleUseCase{
		articleRepo:       articleRepo,
		cacheArticleRepo:  cacheArticleRepo,
		cacheArticlesRepo: cacheArticlesRepo,
	}
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
	if err := au.cacheArticleRepo.GetArticleDetailById(&article, id); err == nil {
		return &article, nil
	}

	err := au.articleRepo.FindByID(&article, id)
	if err != nil {
		return nil, err
	}

	if err := au.cacheArticleRepo.SetArticleDetailById(article, id); err != nil {
		// logのみを出力するエラーハンドリングに変えたい。
		return nil, err
	}
	return &article, nil
}

func (au *articleUseCase) GetArticles(articles *[]model.Article, limit, offset, currentPage int) error {
	if err := au.cacheArticlesRepo.GetLastest(articles, currentPage); err == nil {
		return nil
	}
	err := au.articleRepo.GetArticles(articles, limit, offset)
	if err != nil {
		return err
	}

	if err := au.cacheArticlesRepo.SetLastest(articles, currentPage); err != nil {
		// logのみを出力するエラーハンドリングに変えたい。
		return err
	}
	return nil
}

func (au *articleUseCase) GetPager(currentPage, offset int) (*pager.Pager, error) {
	numOfArticles, err := au.articleRepo.GetPager()
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
