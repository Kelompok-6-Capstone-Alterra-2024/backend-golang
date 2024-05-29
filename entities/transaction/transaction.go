package transaction

import (
	"capstone/controllers/transaction/response"
	"capstone/entities/consultation"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID             uuid.UUID
	ConsultationID uint
	Consultation   consultation.Consultation
	Price          int
	SnapURL        string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (r Transaction) ToResponse() *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:           r.ID.String(),
		Consultation: *r.Consultation.ToResponse(),
		Price:        r.Price,
		SnapURL:      r.SnapURL,
		Status:       r.Status,
	}
}
