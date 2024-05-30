package mood

import (
	"mime/multipart"
)

type Mood struct {
	ID         uint
	UserId     uint
	MoodTypeId uint
	MoodType   MoodType
	Message    string
	Date       string
	ImageUrl   string
}

type MoodType struct {
	ID   uint
	Name string
}

type RepositoryInterface interface {
	SendMood(mood Mood) (Mood, error)
}

type UseCaseInterface interface {
	SendMood(file *multipart.FileHeader, mood Mood) (Mood, error)
}