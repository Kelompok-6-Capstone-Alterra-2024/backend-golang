package story

import (
	"capstone/entities"
	"capstone/entities/doctor"
	"time"
)

type Story struct {
	Id       uint
	Title    string
	Content  string
	Date     time.Time
	ImageUrl string
	ViewCount int
	DoctorId uint
	Doctor   doctor.Doctor
	IsLiked bool
}

type RepositoryInterface interface {
	GetAllStories(metadata entities.Metadata, userId int) ([]Story, error)
	GetStoryById(storyId int, userId int) (Story, error)
	GetLikedStories(metadata entities.Metadata, userId int) ([]Story, error)
	LikeStory(storyId int, userId int) error
}

type UseCaseInterface interface {
	GetAllStories(metadata entities.Metadata, userId int) ([]Story, error)
	GetStoryById(storyId int, userId int) (Story, error)
	GetLikedStories(metadata entities.Metadata, userId int) ([]Story, error)
	LikeStory(storyId int, userId int) error
}