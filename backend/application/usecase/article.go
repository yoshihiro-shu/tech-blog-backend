package usecase

import (
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/repository"
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
}

func NewArticleUseCase(articleRepo repository.ArticleRepository) ArticleUseCase {
	return &articleUseCase{articleRepo: articleRepo}
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
	article := &model.Article{Id: id}
	err := au.articleRepo.FindByID(article)
	if err != nil {
		return nil, err
	}
	return article, nil
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
