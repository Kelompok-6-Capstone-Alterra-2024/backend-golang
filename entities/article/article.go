package article

import (
	"capstone/controllers/article/response"
	"capstone/entities"
	"capstone/entities/doctor"
	"time"
)

type Article struct {
	ID        uint
	Title     string
	Content   string
	Date      time.Time
	ImageUrl  string
	ViewCount int
	DoctorID  uint
	Doctor    doctor.Doctor
	IsLiked   bool
}

type ArticleRepositoryInterface interface {
	CreateArticle(article *Article) (*Article, error)
	GetAllArticle(metadata entities.Metadata, userId int) ([]Article, error)
	GetArticleById(articleId int, userId int) (Article, error)
	GetLikedArticle(metadata entities.Metadata, userId int) ([]Article, error)
}

type ArticleUseCaseInterface interface {
	CreateArticle(article *Article) (*Article, error)
	GetAllArticle(metadata entities.Metadata, userId int) ([]Article, error)
	GetArticleById(articleId int, userId int) (Article, error)
	GetLikedArticle(metadata entities.Metadata, userId int) ([]Article, error)
}

func (ar *Article) ToResponse() response.ArticleCreatedResponse {
	return response.ArticleCreatedResponse{
		ID:        ar.ID,
		DoctorID:  ar.DoctorID,
		Title:     ar.Title,
		Content:   ar.Content,
		ImageUrl:  ar.ImageUrl,
		Date:      ar.Date,
		ViewCount: ar.ViewCount,
		IsLiked:   ar.IsLiked,
	}
}
