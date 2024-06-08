package rating

import (
	"capstone/entities"
	"capstone/entities/doctor"
	"capstone/entities/user"
)

type Rating struct {
	Id uint
	UserId uint
	User user.User
	DoctorId uint
	Doctor doctor.Doctor
	Rate int
	Message string
	Date string
}

type RepositoryInterface interface {
	SendFeedback(rating Rating) (Rating, error)
	GetAllFeedbacks(metadata entities.Metadata, doctorId uint) ([]Rating, error)
}

type UseCaseInterface interface {
	SendFeedback(rating Rating) (Rating, error)
	GetAllFeedbacks(metadata entities.Metadata, doctorId uint) ([]Rating, error)
}