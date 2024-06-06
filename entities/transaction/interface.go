package transaction

import "capstone/entities"

type TransactionRepository interface {
	Insert(transaction *Transaction) (*Transaction, error)
	FindByID(ID string) (*Transaction, error)
	FindByConsultationID(consultationID uint) (*Transaction, error)
	FindAll(metadata *entities.Metadata, userID uint) (*[]Transaction, error)
	Update(transaction *Transaction) (*Transaction, error)
	Delete(ID uint) error
}

type TransactionUseCase interface {
	InsertWithBuiltInInterface(transaction *Transaction) (*Transaction, error)
	InsertWithCustomInterface(transaction *Transaction) (*Transaction, error)
	ConfirmedPayment(id string, transactionStatus string) (*Transaction, error)
	FindByID(ID string) (*Transaction, error)
	FindByConsultationID(consultationID uint) (*Transaction, error)
	FindAll(metadata *entities.Metadata, userID uint) (*[]Transaction, error)
	Update(transaction *Transaction) (*Transaction, error)
	Delete(ID uint) error
}
