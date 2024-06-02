package post

import (
	"capstone/entities"
	"capstone/entities/user"
)

type Post struct {
	ID       uint
	ForumId  uint
	UserId   uint
	Content  string
	ImageUrl string
	User     user.User
}

type RepositoryInterface interface {
	GetAllPostsByForumId(forumId uint, metadata entities.Metadata) ([]Post, error)
	GetPostById(postId uint) (Post, error)
}

type UseCaseInterface interface {
	GetAllPostsByForumId(forumId uint, metadata entities.Metadata) ([]Post, error)
	GetPostById(postId uint) (Post, error)
}