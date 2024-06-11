package otp

import (
	"capstone/constants"
	otpEntities "capstone/entities/otp"
	"time"

	"gorm.io/gorm"
)

type OtpRepo struct {
	db *gorm.DB
}

func NewOtpRepo(db *gorm.DB) *OtpRepo {
	return &OtpRepo{
		db: db,
	}
}

func (otpRepo *OtpRepo) SendOTP(otp otpEntities.Otp) error {
	var otpDB Otp
	otpDB.Email = otp.Email
	otpDB.Code = otp.Code
	otpDB.ExpiredAt = otp.ExpiredAt
	return otpRepo.db.Create(&otpDB).Error
}

func (otpRepo *OtpRepo) VerifyOTP(otp otpEntities.Otp) error {
	var otpDB Otp
	
	err := otpRepo.db.Where("email = ? AND code = ?", otp.Email, otp.Code).First(&otpDB).Error
	if err != nil {
		return constants.ErrInvalidOTP
	}

	if otpDB.ExpiredAt.Before(time.Now()) {
		return constants.ErrExpiredOTP
	}

	return nil
}