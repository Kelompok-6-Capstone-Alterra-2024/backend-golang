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
	InsertWithBuiltIn(transaction *Transaction) (*Transaction, error)
	InsertWithCustom(transaction *Transaction) (*Transaction, error)
	FindByID(ID string) (*Transaction, error)
	FindByConsultationID(consultationID uint) (*Transaction, error)
	FindAll(metadata *entities.Metadata, userID uint) (*[]Transaction, error)
	Update(transaction *Transaction) (*Transaction, error)
	Delete(ID uint) error
}
