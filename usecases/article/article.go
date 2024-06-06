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

func (useCase *ArticleUseCase) GetArticleByIdForDoctor(articleId int) (articleEntities.Article, error) {
	articles, err := useCase.articleRepository.GetArticleByIdForDoctor(articleId)
	if err != nil {
		return articleEntities.Article{}, err
	}
	return articles, nil
}

func (useCase *ArticleUseCase) GetAllArticleByDoctorId(metadata entities.MetadataFull, doctorId int) ([]articleEntities.Article, error) {
	articles, err := useCase.articleRepository.GetAllArticleByDoctorId(metadata, doctorId)
	if err != nil {
		return []articleEntities.Article{}, err
	}
	return articles, nil
}

func (useCase *ArticleUseCase) CountArticleByDoctorId(doctorId int) (int, error) {
	count, err := useCase.articleRepository.CountArticleByDoctorId(doctorId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (useCase *ArticleUseCase) CountArticleLikesByDoctorId(doctorId int) (int, error) {
	count, err := useCase.articleRepository.CountArticleLikesByDoctorId(doctorId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (useCase *ArticleUseCase) CountArticleViewByDoctorId(doctorId int) (int, error) {
	count, err := useCase.articleRepository.CountArticleViewByDoctorId(doctorId)
	if err != nil {
		return 0, err
	}
	return count, nil
}
