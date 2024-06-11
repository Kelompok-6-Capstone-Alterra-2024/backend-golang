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
	UserID        uint
	User          user.User
	Complaint     complaint.Complaint
	Status        string
	PaymentStatus string
	IsAccepted    bool
	IsActive      bool
	Date          time.Time
	Time          time.Time
}

type ConsultationNotes struct {
	ID              uint
	ConsultationID  uint
	Consultation    Consultation
	MusicID         uint
	Music           music.Music
	ForumID         uint
	Forum           forum.Forum
	MainPoint       string
	NextStep        string
	AdditionalNote  string
	MoodTrackerNote string
}

type ConsultationRepository interface {
	CreateConsultation(consultation *Consultation) (*Consultation, error)
	GetConsultationByID(consultationID int) (*Consultation, error)
	GetAllUserConsultation(metadata *entities.Metadata, userID int) (*[]Consultation, error)
	UpdateStatusConsultation(consultation *Consultation) (*Consultation, error)
	GetAllDoctorConsultation(metadata *entities.Metadata, doctorID int) (*[]Consultation, error)
	GetConsultationByComplaintID(complaintID int) (*Consultation, error)
	CountConsultationByStatus(doctorID int, status string) (int64, error)
	CountConsultationToday(doctorID int) (int64, error)
	CountConsultationByDoctorID(doctorID int) (int64, error)
	CreateConsultationNotes(consultationNotes ConsultationNotes) (ConsultationNotes, error)
	GetConsultationNotesByID(consultationID int) (ConsultationNotes, error)
}

type ConsultationUseCase interface {
	CreateConsultation(consultation *Consultation) (*Consultation, error)
	GetConsultationByID(consultationID int) (*Consultation, error)
	GetAllUserConsultation(metadata *entities.Metadata, userID int) (*[]Consultation, error)
	UpdateStatusConsultation(consultation *Consultation) (*Consultation, error)
	GetAllDoctorConsultation(metadata *entities.Metadata, doctorID int) (*[]Consultation, error)
	GetConsultationByComplaintID(complaintID int) (*Consultation, error)
	CountConsultationByDoctorID(doctorID int) (int64, error)
	CountConsultationToday(doctorID int) (int64, error)
	CountConsultationByStatus(doctorID int, status string) (int64, error)
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
		Date:          r.Date.Format("2006-01-02"),
		Time:          r.Time.Format("15:04"),
	}
}

func (r *Consultation) ToDoctorResponse() *response.ConsultationDoctorResponse {
	return &response.ConsultationDoctorResponse{
		ID:            int(r.ID),
		Status:        r.Status,
		PaymentStatus: r.PaymentStatus,
		IsAccepted:    r.IsAccepted,
		IsActive:      r.IsActive,
		Date:          r.Date.Format("2006-01-02"),
		Time:          r.Time.Format("15:04"),
		Complaint:     r.Complaint.ToResponse(),
	}
}
