package request

import (
	"capstone/entities/transaction"
	"github.com/google/uuid"
)

type TransactionRequest struct {
	ConsultationID uint `json:"consultation_id" binding:"required"`
	Price          int  `json:"price" binding:"required"`
	Bank           string
	PaymentType    string
}

func (r TransactionRequest) ToEntities() *transaction.Transaction {
	return &transaction.Transaction{
		ID:             uuid.New(),
		ConsultationID: r.ConsultationID,
		Price:          r.Price,
		Bank:           r.Bank,
		PaymentType:    r.PaymentType,
	}
}
