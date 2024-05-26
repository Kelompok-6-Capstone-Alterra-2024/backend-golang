package doctor

import (
	"capstone/entities"
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Username         string `gorm:"type:varchar(100);unique;not null"`
	Email            string `gorm:"type:varchar(100);unique;not null"`
	Password         string `gorm:"type:varchar(100);not null"`
	Name             string `gorm:"type:varchar(100);not null"`
	Address          string `gorm:"type:text"`
	Bio              string `gorm:"type:text"`
	PhoneNumber      string `gorm:"type:varchar(100)"`
	Gender           string `gorm:"type:ENUM('pria', 'wanita')"`
	IsAvailable      bool   `gorm:"type:boolean;default:true"`
	ProfilePicture   string `gorm:"type:varchar(255)"`
	Balance          int    `gorm:"type:int;default:0"`
	Experience       int    `gorm:"type:int"`
	Almamater        string `gorm:"type:varchar(100)"`
	GraduationYear   int    `gorm:"type:int"`
	PracticeLocation string `gorm:"type:text"`
	PracticeCity     string `gorm:"type:varchar(100)"`
	PracticeProvince string `gorm:"type:varchar(100)"`
	StrNumber        string `gorm:"type:varchar(100)"`
	Fee              int    `gorm:"type:int"`
	Specialist       string `gorm:"type:varchar(100)"`
}

func ToDoctorModel(doctor *entities.Doctor) *Doctor {
	return &Doctor{
		Username:         doctor.Username,
		Email:            doctor.Email,
		Password:         doctor.Password,
		Name:             doctor.Name,
		Address:          doctor.Address,
		Bio:              doctor.Bio,
		PhoneNumber:      doctor.PhoneNumber,
		Gender:           doctor.Gender,
		ProfilePicture:   doctor.ProfilePicture,
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
