package consultation

import (
	"capstone/controllers/consultation/response"
	"capstone/entities"
	"capstone/entities/complaint"
	"capstone/entities/doctor"
	"capstone/entities/forum"
	"capstone/entities/music"
	"capstone/entities/user"
	"time"
)

type Consultation struct {
	ID            uint
	DoctorID      uint
	Doctor        *doctor.Doctor
	UserID        int
	User          user.User
	Complaint     complaint.Complaint
	Status        string
	PaymentStatus string
	IsAccepted    bool
	IsActive      bool
	Date          time.Time
}

type ConsultationNotes struct {
	ID        uint
	ConsultationID uint
	Consultation   Consultation
	MusicID   uint
	Music     music.Music
	ForumID   uint
	Forum     forum.Forum
	MainPoint string
	NextStep  string
	AdditionalNote string
	MoodTrackerNote string
}

type ConsultationRepository interface {
	CreateConsultation(consultation *Consultation) (*Consultation, error)
	GetConsultationByID(consultationID int) (*Consultation, error)
	GetAllUserConsultation(metadata *entities.Metadata, userID int) (*[]Consultation, error)
	UpdateStatusConsultation(consultation *Consultation) (*Consultation, error)
	GetAllDoctorConsultation(metadata *entities.Metadata, doctorID int) (*[]Consultation, error)
	CreateConsultationNotes(consultationNotes ConsultationNotes) (ConsultationNotes, error)
	GetConsultationNotesByID(consultationID int) (ConsultationNotes, error)
}

type ConsultationUseCase interface {
	CreateConsultation(consultation *Consultation) (*Consultation, error)
	GetConsultationByID(consultationID int) (*Consultation, error)
	GetAllUserConsultation(metadata *entities.Metadata, userID int) (*[]Consultation, error)
	UpdateStatusConsultation(consultation *Consultation) (*Consultation, error)
	GetAllDoctorConsultation(metadata *entities.Metadata, doctorID int) (*[]Consultation, error)
	CreateConsultationNotes(consultationNotes ConsultationNotes) (ConsultationNotes, error)
	GetConsultationNotesByID(consultationID int) (ConsultationNotes, error)
}

func (r *Consultation) ToUserResponse() *response.ConsultationUserResponse {
	return &response.ConsultationUserResponse{
		ID:            int(r.ID),
		Doctor:        r.Doctor.ToDoctorResponse(),
		Status:        r.Status,
		PaymentStatus: r.PaymentStatus,
		IsAccepted:    r.IsAccepted,
		IsActive:      r.IsActive,
		Date:          r.Date,
	}
}

func (r *Consultation) ToDoctorResponse() *response.ConsultationDoctorResponse {
	return &response.ConsultationDoctorResponse{
		ID:            int(r.ID),
		Status:        r.Status,
		PaymentStatus: r.PaymentStatus,
		IsAccepted:    r.IsAccepted,
		IsActive:      r.IsActive,
		Date:          r.Date,
		Complaint:     r.Complaint.ToResponse(),
	}
}
