package forum

import (
	"capstone/entities/doctor"
	"capstone/entities/user"
)

type Forum struct {
	ID          uint
	Name        string
	Description string
	ImageUrl    string
	DoctorID    uint
	Doctor      doctor.Doctor
}

type ForumMember struct {
	ID      uint
	ForumID uint
	Forum   Forum
	UserID  uint
	User    user.User
}

type RepositoryInterface interface {
	JoinForum(forumId uint, userId uint) (error)
}

type UseCaseInterface interface {
	JoinForum(forumId uint, userId uint) (error)
}