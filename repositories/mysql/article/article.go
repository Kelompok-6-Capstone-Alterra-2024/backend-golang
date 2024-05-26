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

// type ArticleRepoWithDoctorID struct {
// 	db       *gorm.DB
// 	doctorID uint
// }

// func NewArticleRepoWithDoctorID(db *gorm.DB, doctorID uint) *ArticleRepoWithDoctorID {
// 	return &ArticleRepoWithDoctorID{
// 		db:       db,
// 		doctorID: doctorID,
// 	}
// }

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
