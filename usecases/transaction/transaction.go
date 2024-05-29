package transaction

import (
	"capstone/entities"
	transactionEntities "capstone/entities/transaction"
)

type Transaction struct {
	transactionRepository transactionEntities.TransactionRepository
}

func NewTransaction(transactionRepository transactionEntities.TransactionRepository) transactionEntities.TransactionRepository {
	return &Transaction{transactionRepository}
}

func (usecase *Transaction) Insert(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *Transaction) FindByID(ID uint) (*transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *Transaction) FindByConsultationID(consultationID uint) (*transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *Transaction) FindAll(metadata *entities.Metadata, userID uint) (*[]transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *Transaction) Update(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *Transaction) Delete(ID uint) error {
	//TODO implement me
	panic("implement me")
}
