package transaction

import (
	"capstone/entities/consultation"
	"time"
)

type Transaction struct {
	ID             uint
	ConsultationID uint
	Consultation   consultation.Consultation
	Price          int
	SnapURL        string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
