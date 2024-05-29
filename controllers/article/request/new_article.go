package request

import (
	"capstone/entities/article"
)

type CreateArticleRequest struct {
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	ImageURL string `json:"image_url" form:"image_url"`
}

func (r *CreateArticleRequest) ToArticleEntities() *article.Article {
	return &article.Article{
		Title:    r.Title,
		Content:  r.Content,
		ImageURL: r.ImageURL,
	}
}
