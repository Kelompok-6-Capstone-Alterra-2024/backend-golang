package article

import (
	"capstone/constants"
	articleEntities "capstone/entities/article"
)

type ArticleUseCase struct {
	articleRepository articleEntities.ArticleRepositoryInterface
}

func NewArticleUseCase(articleRepository articleEntities.ArticleRepositoryInterface) *ArticleUseCase {
	return &ArticleUseCase{
		articleRepository: articleRepository,
	}
}

func (useCase *ArticleUseCase) CreateArticle(article *articleEntities.Article) (*articleEntities.Article, error) {
	if article.Title == "" || article.Content == "" {
		return nil, constants.ErrEmptyInputArticle
	}

	createdArticle, err := useCase.articleRepository.CreateArticle(article)
	if err != nil {
		return nil, err
	}

	return createdArticle, nil
}

func (useCase *ArticleUseCase) GetAllArticle() ([]*articleEntities.Article, error) {
	return useCase.articleRepository.GetAllArticle()
}
