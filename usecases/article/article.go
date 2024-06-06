package article

import (
	"capstone/constants"
	"capstone/entities"
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

func (useCase *ArticleUseCase) CreateArticle(article *articleEntities.Article, userId int) (*articleEntities.Article, error) {
	if article.Title == "" || article.Content == "" {
		return nil, constants.ErrEmptyInputArticle
	}

	createdArticle, err := useCase.articleRepository.CreateArticle(article, userId)
	if err != nil {
		return nil, err
	}

	return createdArticle, nil
}

func (useCase *ArticleUseCase) GetAllArticle(metadata entities.Metadata, userId int) ([]articleEntities.Article, error) {
	articles, err := useCase.articleRepository.GetAllArticle(metadata, userId)
	if err != nil {
		return []articleEntities.Article{}, err
	}
	return articles, nil
}

func (useCase *ArticleUseCase) GetArticleById(articleId int, userId int) (articleEntities.Article, error) {
	article, err := useCase.articleRepository.GetArticleById(articleId, userId)
	if err != nil {
		return articleEntities.Article{}, err
	}
	return article, nil
}

func (useCase *ArticleUseCase) GetLikedArticle(metadata entities.Metadata, userId int) ([]articleEntities.Article, error) {
	articles, err := useCase.articleRepository.GetLikedArticle(metadata, userId)
	if err != nil {
		return []articleEntities.Article{}, err
	}
	return articles, nil
}

func (useCase *ArticleUseCase) LikeArticle(articleId int, userId int) error {
	err := useCase.articleRepository.LikeArticle(articleId, userId)
	if err != nil {
		return err
	}
	return nil
}
