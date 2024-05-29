package transaction

import (
	"capstone/entities"
	midtransEntities "capstone/entities/midtrans"
	transactionEntities "capstone/entities/transaction"
	"fmt"
)

type Transaction struct {
	transactionRepository transactionEntities.TransactionRepository
	midtransUseCase       midtransEntities.MidtransUseCase
}

func NewTransactionUseCase(transactionRepository transactionEntities.TransactionRepository, midtransUseCase midtransEntities.MidtransUseCase) transactionEntities.TransactionRepository {
	return &Transaction{
		transactionRepository: transactionRepository,
		midtransUseCase:       midtransUseCase,
	}
}

func (usecase *Transaction) Insert(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	newTransaction, err := usecase.midtransUseCase.GenerateSnapURL(transaction)
	if err != nil {
		return nil, err
	}
	fmt.Println(newTransaction.SnapURL)
	response, err := usecase.transactionRepository.Insert(newTransaction)
	if err != nil {
		return nil, err
	}
	return response, nil
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
