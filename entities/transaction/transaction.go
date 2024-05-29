package transaction

import (
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
