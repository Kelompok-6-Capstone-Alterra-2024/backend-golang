package request

import "capstone/entities/transaction"

type TransactionRequest struct {
	ConsultationID uint `json:"consultation_id" binding:"required"`
	Price          int  `json:"price" binding:"required"`
}

func (r TransactionRequest) ToEntities() *transaction.Transaction {
	return &transaction.Transaction{
		ConsultationID: r.ConsultationID,
		Price:          r.Price,
	}
}
