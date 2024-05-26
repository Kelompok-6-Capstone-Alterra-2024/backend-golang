package entities

import "capstone/repositories/mysql/doctor"

type Doctor struct {
	ID               uint
	Username         string
	Email            string
	Password         string
	Name             string
	Address          string
	Bio              string
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

func ToDoctorEntities(doctor *doctor.Doctor) *Doctor {
	return &Doctor{
		ID:               doctor.ID,
		Username:         doctor.Username,
		Email:            doctor.Email,
		Password:         doctor.Password,
		Name:             doctor.Name,
		Address:          doctor.Address,
		Bio:              doctor.Bio,
		PhoneNumber:      doctor.PhoneNumber,
		Gender:           doctor.Gender,
		IsAvailable:      doctor.IsAvailable,
		ProfilePicture:   doctor.ProfilePicture,
		Balance:          doctor.Balance,
		Experience:       doctor.Experience,
		Almamater:        doctor.Almamater,
		GraduationYear:   doctor.GraduationYear,
		PracticeLocation: doctor.PracticeLocation,
		PracticeCity:     doctor.PracticeCity,
		PracticeProvince: doctor.PracticeProvince,
		StrNumber:        doctor.StrNumber,
		Fee:              doctor.Fee,
		Specialist:       doctor.Specialist,
	}
}
