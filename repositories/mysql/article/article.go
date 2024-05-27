package article

import (
	articleEntities "capstone/entities/article"

	"gorm.io/gorm"
)

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{
		db: db,
	}
}

func (repository *ArticleRepo) CreateArticle(article *articleEntities.Article) (*articleEntities.Article, error) {
	articleDB := Article{
		DoctorID: article.DoctorID,
		Title:    article.Title,
		Content:  article.Content,
		ImageURL: article.ImageURL,
	}

	if err := repository.db.Create(&articleDB).Error; err != nil {
		return nil, err
	}

	articleEntity := articleDB.ToEntities()
	return articleEntity, nil
}

func (repository *ArticleRepo) GetAllArticle() ([]*articleEntities.Article, error) {
	var articles []*articleEntities.Article
	if err := repository.db.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}