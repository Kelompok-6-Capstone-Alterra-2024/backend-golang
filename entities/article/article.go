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
	CreateArticle(article *Article, userId int) (*Article, error)
	GetAllArticle(metadata entities.Metadata, userId int) ([]Article, error)
	GetArticleById(articleId int, userId int) (Article, error)
	GetLikedArticle(metadata entities.Metadata, userId int) ([]Article, error)
	LikeArticle(articleId int, userId int) error
}

type ArticleUseCaseInterface interface {
	CreateArticle(article *Article, userId int) (*Article, error)
	GetAllArticle(metadata entities.Metadata, userId int) ([]Article, error)
	GetArticleById(articleId int, userId int) (Article, error)
	GetLikedArticle(metadata entities.Metadata, userId int) ([]Article, error)
	LikeArticle(articleId int, userId int) error
}

func (ar *Article) ToResponse() response.ArticleListResponse {
	return response.ArticleListResponse{
		ID:        ar.ID,
		DoctorID:  ar.DoctorID,
		Title:     ar.Title,
		Content:   ar.Content,
		ImageUrl:  ar.ImageUrl,
		Date:      ar.Date,
		ViewCount: ar.ViewCount,
		IsLiked:   ar.IsLiked,
		Doctor: response.DoctorInfoResponse{
			ID:   ar.Doctor.ID,
			Name: ar.Doctor.Name,
		},
	}
}
