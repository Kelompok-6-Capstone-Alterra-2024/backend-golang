package article

import (
	"gorm.io/gorm"

	articleEntities "capstone/entities/article"
)

type Article struct {
	gorm.Model
	DoctorID uint   `gorm:"type:int;not null"`
	Title    string `gorm:"type:varchar(100);not null"`
	Content  string `gorm:"type:text"`
	ImageURL string `gorm:"type:varchar(255)"`
}

func (article *Article) ToEntities() *articleEntities.Article {
	return &articleEntities.Article{
		ID:       article.ID,
		DoctorID: article.DoctorID,
		Title:    article.Title,
		Content:  article.Content,
		ImageURL: article.ImageURL,
	}
}

func ToArticleModel(request *articleEntities.Article) *Article {
	return &Article{
		DoctorID: request.DoctorID,
		Title:    request.Title,
		Content:  request.Content,
		ImageURL: request.ImageURL,
	}
}
