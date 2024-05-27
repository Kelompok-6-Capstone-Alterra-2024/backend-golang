package consultation

import (
	"capstone/entities/doctor"
	"capstone/entities/user"
)

type Consultation struct {
	ID            uint
	DoctorID      uint
	Doctor        doctor.Doctor
	UserID        int
	User          user.User
	Status        string
	PaymentStatus string
	IsAccepted    bool
	IsActive      bool
	Date          string
}

type ConsultationRepository interface {
	CreateConsultation(consultation *Consultation) (*Consultation, error)
	GetConsultationByID(consultationID int) (*Consultation, error)
	GetAllConsultation(userID int) (*[]Consultation, error)
}

type ConsultationUseCase interface {
	CreateConsultation(consultation *Consultation) (*Consultation, error)
	GetConsultationByID(consultationID int) (*Consultation, error)
	GetAllConsultation(userID int) (*[]Consultation, error)
}
