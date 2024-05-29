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
	// GetStoryById(id int) (*Stories, error)
}

type UseCaseInterface interface {
	GetAllStories(metadata entities.Metadata, userId int) ([]Story, error)
	// GetStoryById(id int) (*Stories, error)
}