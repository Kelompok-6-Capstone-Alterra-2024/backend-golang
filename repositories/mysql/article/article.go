package article

import (
	"capstone/entities"
	articleEntities "capstone/entities/article"
	doctorEntities "capstone/entities/doctor"

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
		Date:     article.Date,
		ImageUrl: article.ImageUrl,
	}

	if err := repository.db.Create(&articleDB).Error; err != nil {
		return nil, err
	}

	articleEntity := articleDB.ToEntities()
	return articleEntity, nil
}

func (repository *ArticleRepo) GetAllArticle(metadata entities.Metadata, userId int) ([]articleEntities.Article, error) {
	var articlesDb []Article

	// Pagination
	err := repository.db.Limit(metadata.Limit).Offset((metadata.Page - 1) * metadata.Limit).Find(&articlesDb).Error
	if err != nil {
		return nil, err
	}

	articleLikes := make([]ArticleLikes, len(articlesDb))
	var counter int64
	var isLiked []bool

	// Check if articles are liked by the user
	for i := 0; i < len(articlesDb); i++ {
		articleLikes[i].UserId = uint(userId)
		articleLikes[i].ArticleID = articlesDb[i].ID
		err = repository.db.Model(&articleLikes[i]).Where("user_id = ? AND article_id = ?", articleLikes[i].UserId, articleLikes[i].ArticleID).Count(&counter).Error

		if err != nil {
			return nil, err
		}

		if counter > 0 {
			isLiked = append(isLiked, true)
		} else {
			isLiked = append(isLiked, false)
		}

		counter = 0
	}

	articlesEnt := make([]articleEntities.Article, len(articlesDb))
	for i := 0; i < len(articlesDb); i++ {
		articlesEnt[i] = articleEntities.Article{
			ID:        articlesDb[i].ID,
			Title:     articlesDb[i].Title,
			Content:   articlesDb[i].Content,
			Date:      articlesDb[i].Date,
			ImageUrl:  articlesDb[i].ImageUrl,
			ViewCount: articlesDb[i].ViewCount,
			DoctorID:  articlesDb[i].DoctorID,
			Doctor: doctorEntities.Doctor{
				ID:   articlesDb[i].Doctor.ID,
				Name: articlesDb[i].Doctor.Name,
			},
			IsLiked: isLiked[i],
		}
	}

	return articlesEnt, nil
}
