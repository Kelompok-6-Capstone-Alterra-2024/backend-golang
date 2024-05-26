package doctor

import "capstone/controllers/doctor/response"

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
}

type DoctorRepositoryInterface interface {
	Register(doctor *Doctor) (*Doctor, error)
	Login(doctor *Doctor) (*Doctor, error)
}

type DoctorUseCaseInterface interface {
	Register(doctor *Doctor) (*Doctor, error)
	Login(doctor *Doctor) (*Doctor, error)
}

func (r *Doctor) ToResponse() response.DoctorLoginAndRegisterResponse {
	return response.DoctorLoginAndRegisterResponse{
		ID:    r.ID,
		Token: r.Token,
	}
}
