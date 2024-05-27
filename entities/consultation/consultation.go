package consultation

import (
	"capstone/controllers/consultation/response"
	"capstone/entities"
	"capstone/entities/doctor"
	"capstone/entities/user"
	"time"
)

type Consultation struct {
	ID            uint
	DoctorID      uint
	Doctor        *doctor.Doctor
	UserID        int
	User          user.User
	Status        string
	PaymentStatus string
	IsAccepted    bool
	IsActive      bool
	Date          time.Time
}

type ConsultationRepository interface {
	CreateConsultation(consultation *Consultation) (*Consultation, error)
	GetConsultationByID(consultationID int) (*Consultation, error)
	GetAllConsultation(metadata *entities.Metadata, userID int) (*[]Consultation, error)
}

type ConsultationUseCase interface {
	CreateConsultation(consultation *Consultation) (*Consultation, error)
	GetConsultationByID(consultationID int) (*Consultation, error)
	GetAllConsultation(metadata *entities.Metadata, userID int) (*[]Consultation, error)
}

func (r *Consultation) ToResponse() *response.ConsultationResponse {
	return &response.ConsultationResponse{
		ID:            int(r.ID),
		Doctor:        r.Doctor.ToDoctorResponse(),
		Status:        r.Status,
		PaymentStatus: r.PaymentStatus,
		IsAccepted:    r.IsAccepted,
		IsActive:      r.IsActive,
		Date:          r.Date,
	}
}
