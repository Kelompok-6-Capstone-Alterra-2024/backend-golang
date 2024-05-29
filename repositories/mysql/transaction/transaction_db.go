package transaction

import (
	"capstone/entities/transaction"
	"capstone/repositories/mysql/consultation"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ConsultationID uint                      `gorm:"column:consultation_id;not null"`
	Consultation   consultation.Consultation `gorm:"foreignKey:consultation_id;references:id"`
	Price          int                       `gorm:"column:price;not null"`
	SnapURL        string                    `gorm:"column:snap_url"`
	Status         string                    `gorm:"column:status;not null;type:enum('pending','success','failed');default:'pending'"`
}

func (receiver Transaction) ToEntities() *transaction.Transaction {
	return &transaction.Transaction{
		ID:             receiver.ID,
		ConsultationID: receiver.ConsultationID,
		Consultation:   *receiver.Consultation.ToEntities(),
		Price:          receiver.Price,
		SnapURL:        receiver.SnapURL,
		Status:         receiver.Status,
		CreatedAt:      receiver.CreatedAt,
		UpdatedAt:      receiver.UpdatedAt,
	}
}

func ToTransactionModel(transaction *transaction.Transaction) *Transaction {
	return &Transaction{
		Model:          gorm.Model{ID: transaction.ID, CreatedAt: transaction.CreatedAt, UpdatedAt: transaction.UpdatedAt},
		ConsultationID: transaction.ConsultationID,
		Consultation:   *consultation.ToConsultationModel(&transaction.Consultation),
		Price:          transaction.Price,
		SnapURL:        transaction.SnapURL,
		Status:         transaction.Status,
	}
}
