package doctor

import (
	"capstone/constants"
	doctorEntities "capstone/entities/doctor"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DoctorRepo struct {
	db *gorm.DB
}

func NewDoctorRepo(db *gorm.DB) *DoctorRepo {
	return &DoctorRepo{
		db: db,
	}
}

func (repository *DoctorRepo) Register(doctor *doctorEntities.Doctor) (*doctorEntities.Doctor, error) {

	doctorDb := ToDoctorModel(doctor)
	if err := repository.db.Model(&doctorDb).First(&doctorDb, "username = ?", doctorDb.Username).Error; err == nil {
		return nil, constants.ErrUsernameAlreadyExist
	}

	if err := repository.db.Model(&doctorDb).First(&doctorDb, "email = ?", doctorDb.Email).Error; err == nil {
		return nil, constants.ErrEmailAlreadyExist
	}

	if err := repository.db.Create(&doctorDb).Error; err != nil {
		return nil, constants.ErrInsertDatabase
	}

	doctorResult := doctorDb.ToEntities()
	return doctorResult, nil
}

func (repository *DoctorRepo) Login(doctor *doctorEntities.Doctor) (*doctorEntities.Doctor, error) {
	doctorDb := ToDoctorModel(doctor)

	doctorPassword := doctorDb.Password
	if err := repository.db.First(&doctorDb, "username LIKE ? OR email LIKE ?", doctorDb.Username, doctorDb.Email).Error; err != nil {
		return nil, constants.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(doctorDb.Password), []byte(doctorPassword)); err != nil {
		return nil, constants.ErrUserNotFound
	}

	result := doctorDb.ToEntities()

	return result, nil

}
