package post

import "capstone/entities/user"

type Post struct {
	ID       uint
	ForumId  uint
	UserId   uint
	Content  string
	ImageUrl string
	User     user.User
}

type RepositoryInterface interface {
}

type UseCaseInterface interface {
}