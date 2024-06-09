package transaction

import (
	"capstone/controllers/transaction/response"
	"capstone/entities/consultation"
	"capstone/entities/doctor"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID             uuid.UUID
	ConsultationID uint `validate:"required"`
	Consultation   consultation.Consultation
	DoctorID       uint
	Doctor         doctor.Doctor
	Price          int
	Status         string
	PaymentType    string
	PaymentLink    string
	Bank           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (r Transaction) ToResponse() *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:           r.ID.String(),
		Consultation: *r.Consultation.ToUserResponse(),
		Price:        r.Price,
		PaymentType:  r.PaymentType,
		PaymentLink:  r.PaymentLink,
		Bank:         r.Bank,
		Status:       r.Status,
		CreatedAt:    r.CreatedAt.String(),
		UpdatedAt:    r.UpdatedAt.String(),
	}
}
