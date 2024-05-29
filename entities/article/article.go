package article

import "capstone/controllers/article/response"

type Article struct {
	ID       uint
	DoctorID uint
	Title    string
	Content  string
	ImageURL string
}

type ArticleRepositoryInterface interface {
	CreateArticle(article *Article) (*Article, error)
	GetAllArticle() ([]*Article, error)
}

type ArticleUseCaseInterface interface {
	CreateArticle(article *Article) (*Article, error)
	GetAllArticle() ([]*Article, error)
}

func (ar *Article) ToResponse() response.ArticleCreatedResponse {
	return response.ArticleCreatedResponse{
		ID:       ar.ID,
		DoctorID: ar.DoctorID,
		Title:    ar.Title,
		Content:  ar.Content,
		ImageURL: ar.ImageURL,
	}
}
