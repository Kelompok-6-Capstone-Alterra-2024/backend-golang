package consultation

import (
	"capstone/entities/consultation"
	"capstone/repositories/mysql/complaint"
	"capstone/repositories/mysql/doctor"
	"capstone/repositories/mysql/user"
	"gorm.io/gorm"
	"time"
)

type Consultation struct {
	gorm.Model
	DoctorID      uint                `gorm:"column:doctor_id;not null"`
	Doctor        doctor.Doctor       `gorm:"foreignKey:doctor_id;references:id"`
	UserID        uint                `gorm:"column:user_id;not null"`
	User          user.User           `gorm:"foreignKey:user_id;references:id"`
	ComplaintID   int                 `gorm:"column:complaint_id;unique;default:NULL"`
	Complaint     complaint.Complaint `gorm:"foreignKey:complaint_id;references:id"`
	Status        string              `gorm:"column:status;not null;default:'pending';type:enum('pending', 'rejected', 'incoming', 'active', 'done')"`
	PaymentStatus string              `gorm:"column:payment_status;not null;type:enum('pending', 'paid', 'canceled');default:'pending'"`
	IsAccepted    bool                `gorm:"column:is_accepted"`
	IsActive      bool                `gorm:"column:is_active"`
	Date          time.Time           `gorm:"column:date;type:date;not null"`
	Time          string              `gorm:"column:time;type:time(3);not null"`
}

func (receiver Consultation) ToEntities() (*consultation.Consultation, error) {
	consultationTime, err := time.Parse("15:04:05.000", receiver.Time)
	if err != nil {
		return nil, err
	}
	return &consultation.Consultation{
		ID:            receiver.ID,
		DoctorID:      receiver.DoctorID,
		Doctor:        receiver.Doctor.ToEntities(),
		UserID:        receiver.UserID,
		Complaint:     *receiver.Complaint.ToEntities(),
		Status:        receiver.Status,
		PaymentStatus: receiver.PaymentStatus,
		IsAccepted:    receiver.IsAccepted,
		IsActive:      receiver.IsActive,
		Date:          receiver.Date,
		Time:          consultationTime,
	}, nil
}

func ToConsultationModel(request *consultation.Consultation) *Consultation {
	return &Consultation{
		Model:         gorm.Model{ID: request.ID},
		DoctorID:      request.DoctorID,
		UserID:        request.UserID,
		Status:        request.Status,
		PaymentStatus: request.PaymentStatus,
		IsAccepted:    request.IsAccepted,
		IsActive:      request.IsActive,
		Date:          request.Date,
		Time:          request.Time.Format("15:04"),
	}
}
