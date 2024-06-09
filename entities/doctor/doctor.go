package doctor

import (
	"capstone/controllers/doctor/response"
	"capstone/entities"
	"context"
)

type Doctor struct {
	ID               uint
	Username         string
	Email            string
	Password         string
	Name             string
	Address          string
	PhoneNumber      string
	Gender           string
	IsAvailable      bool
	ProfilePicture   string
	Balance          int
	Experience       int
	Almamater        string
	GraduationYear   int
	PracticeLocation string
	PracticeCity     string
	PracticeProvince string
	StrNumber        string
	Fee              int
	Specialist       string
	Token            string
	IsOauth          bool
}

type DoctorRepositoryInterface interface {
	Register(doctor *Doctor) (*Doctor, error)
	Login(doctor *Doctor) (*Doctor, error)
	GetDoctorByID(doctorID int) (*Doctor, error)
	GetAllDoctor(metadata *entities.Metadata) (*[]Doctor, error)
	GetActiveDoctor(metadata *entities.Metadata) (*[]Doctor, error)
	Create(email string, picture string, name string, username string) (Doctor, error)
	OauthFindByEmail(email string) (Doctor, int, error)
}

type DoctorUseCaseInterface interface {
	Register(doctor *Doctor) (*Doctor, error)
	Login(doctor *Doctor) (*Doctor, error)
	GetDoctorByID(doctorID int) (*Doctor, error)
	GetAllDoctor(metadata *entities.Metadata) (*[]Doctor, error)
	GetActiveDoctor(metadata *entities.Metadata) (*[]Doctor, error)
	HandleGoogleLogin() string
	HandleGoogleCallback(ctx context.Context, code string) (Doctor, error)
}

func (r *Doctor) ToResponse() response.DoctorLoginAndRegisterResponse {
	return response.DoctorLoginAndRegisterResponse{
		ID:    r.ID,
		Token: r.Token,
	}
}

func (r *Doctor) ToDoctorResponse() *response.DoctorResponse {
	return &response.DoctorResponse{
		ID:               r.ID,
		Username:         r.Username,
		Email:            r.Email,
		Name:             r.Name,
		Address:          r.Address,
		PhoneNumber:      r.PhoneNumber,
		Gender:           r.Gender,
		IsAvailable:      r.IsAvailable,
		ProfilePicture:   r.ProfilePicture,
		Balance:          r.Balance,
		Experience:       r.Experience,
		Almamater:        r.Almamater,
		GraduationYear:   r.GraduationYear,
		PracticeLocation: r.PracticeLocation,
		PracticeCity:     r.PracticeCity,
		PracticeProvince: r.PracticeProvince,
		StrNumber:        r.StrNumber,
		Fee:              r.Fee,
		Specialist:       r.Specialist,
	}
}
