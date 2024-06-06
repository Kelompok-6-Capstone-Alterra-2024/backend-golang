package transaction

import (
	"capstone/constants"
	"capstone/entities"
	midtransEntities "capstone/entities/midtrans"
	transactionEntities "capstone/entities/transaction"
	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	transactionRepository transactionEntities.TransactionRepository
	midtransUseCase       midtransEntities.MidtransUseCase
	validate              *validator.Validate
}

func NewTransactionUseCase(
	transactionRepository transactionEntities.TransactionRepository,
	midtransUseCase midtransEntities.MidtransUseCase,
	validate *validator.Validate,
) transactionEntities.TransactionUseCase {
	return &Transaction{
		transactionRepository: transactionRepository,
		midtransUseCase:       midtransUseCase,
		validate:              validate,
	}
}

func (usecase *Transaction) InsertWithBuiltIn(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	if err := usecase.validate.Struct(transaction); err != nil {
		return nil, err
	}

	newTransaction, err := usecase.midtransUseCase.GenerateSnapURL(transaction)
	if err != nil {
		return nil, err
	}

	response, err := usecase.transactionRepository.Insert(newTransaction)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (usecase *Transaction) InsertWithCustom(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	var err error
	if err = usecase.validate.Struct(transaction); err != nil {
		return nil, err
	}

	var newTransaction, response *transactionEntities.Transaction
	if transaction.PaymentType == "bank_transfer" {
		newTransaction, err = usecase.midtransUseCase.BankTransfer(transaction)
		if err != nil {
			return nil, err
		}
	}
	if transaction.PaymentType == "ewallet" {
		newTransaction, err = usecase.midtransUseCase.EWallet(transaction)
		if err != nil {
			return nil, err
		}
	}

	response, err = usecase.transactionRepository.Insert(newTransaction)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (usecase *Transaction) FindByID(ID string) (*transactionEntities.Transaction, error) {
	newTransaction, err := usecase.transactionRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return newTransaction, nil
}

func (usecase *Transaction) FindByConsultationID(consultationID uint) (*transactionEntities.Transaction, error) {
	newTransaction, err := usecase.transactionRepository.FindByConsultationID(consultationID)
	if err != nil {
		return nil, err
	}
	return newTransaction, nil
}

func (usecase *Transaction) FindAll(metadata *entities.Metadata, userID uint) (*[]transactionEntities.Transaction, error) {
	newTransaction, err := usecase.transactionRepository.FindAll(metadata, userID)
	if err != nil {
		return nil, err
	}
	if len(*newTransaction) == 0 {
		return nil, constants.ErrDataEmpty
	}
	return newTransaction, nil
}

func (usecase *Transaction) Update(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *Transaction) Delete(ID uint) error {
	//TODO implement me
	panic("implement me")
}
