package forum

import (
	"capstone/entities"
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
	NumberOfMembers int
	User            []user.User
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
	GetJoinedForum(userId uint, metadata entities.Metadata) ([]Forum, error)
}

type UseCaseInterface interface {
	JoinForum(forumId uint, userId uint) (error)
	GetJoinedForum(userId uint, metadata entities.Metadata) ([]Forum, error)
}