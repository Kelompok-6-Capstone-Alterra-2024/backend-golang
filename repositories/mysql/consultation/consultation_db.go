package consultation

import (
	"capstone/entities/consultation"
	"capstone/repositories/mysql/doctor"
	"capstone/repositories/mysql/user"
	"gorm.io/gorm"
	"time"
)

type Consultation struct {
	gorm.Model
	DoctorID      uint          `gorm:"column:doctor_id;not null"`
	Doctor        doctor.Doctor `gorm:"foreignKey:doctor_id;references:id"`
	UserID        int           `gorm:"column:user_id;not null"`
	User          user.User     `gorm:"foreignKey:user_id;references:id"`
	ComplaintID   int           `gorm:"column:complaint_id;unique;default:NULL"`
	Status        string        `gorm:"column:status;not null;default:'pending'"`
	PaymentStatus string        `gorm:"column:payment_status;not null;default:'pending'"`
	IsAccepted    bool          `gorm:"column:is_accepted"`
	IsActive      bool          `json:"is_active"`
	Date          time.Time     `json:"date"`
}

func (receiver Consultation) ToEntities() *consultation.Consultation {
	return &consultation.Consultation{
		ID:            receiver.ID,
		DoctorID:      receiver.DoctorID,
		Doctor:        receiver.Doctor.ToEntities(),
		UserID:        receiver.UserID,
		Status:        receiver.Status,
		PaymentStatus: receiver.PaymentStatus,
		IsAccepted:    receiver.IsAccepted,
		IsActive:      receiver.IsActive,
		Date:          receiver.Date,
	}
}

func ToConsultationModel(request *consultation.Consultation) *Consultation {
	return &Consultation{
		DoctorID:      request.DoctorID,
		UserID:        request.UserID,
		Status:        request.Status,
		PaymentStatus: request.PaymentStatus,
		IsAccepted:    request.IsAccepted,
		IsActive:      request.IsActive,
		Date:          request.Date,
	}
}
