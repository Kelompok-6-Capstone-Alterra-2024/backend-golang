package transaction

import (
	"capstone/constants"
	"capstone/entities"
	"capstone/entities/consultation"
	doctorEntities "capstone/entities/doctor"
	midtransEntities "capstone/entities/midtrans"
	transactionEntities "capstone/entities/transaction"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	transactionRepository  transactionEntities.TransactionRepository
	midtransUseCase        midtransEntities.MidtransUseCase
	consultationRepository consultation.ConsultationRepository
	doctorRepository       doctorEntities.DoctorRepositoryInterface
	validate               *validator.Validate
}

func NewTransactionUseCase(
	transactionRepository transactionEntities.TransactionRepository,
	midtransUseCase midtransEntities.MidtransUseCase,
	consultationRepository consultation.ConsultationRepository,
	doctorRepository doctorEntities.DoctorRepositoryInterface,
	validate *validator.Validate,
) transactionEntities.TransactionUseCase {
	return &Transaction{
		transactionRepository:  transactionRepository,
		midtransUseCase:        midtransUseCase,
		consultationRepository: consultationRepository,
		doctorRepository:       doctorRepository,
		validate:               validate,
	}
}

func (usecase *Transaction) InsertWithBuiltInInterface(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
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

func (usecase *Transaction) InsertWithCustomInterface(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	var err error
	if err = usecase.validate.Struct(transaction); err != nil {
		return nil, err
	}

	_, err = usecase.consultationRepository.GetConsultationByID(int(transaction.ConsultationID))
	if err != nil {
		return nil, err
	}
	var newTransaction, response *transactionEntities.Transaction
	if transaction.PaymentType == constants.BankTransfer {
		newTransaction, err = usecase.midtransUseCase.BankTransfer(transaction)
		if err != nil {
			return nil, err
		}
	}
	if transaction.PaymentType == constants.GoPay {
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

func (usecase *Transaction) FindAllByUserID(metadata *entities.Metadata, userID uint) (*[]transactionEntities.Transaction, error) {
	newTransaction, err := usecase.transactionRepository.FindAllByUserID(metadata, userID)
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

func (usecase *Transaction) ConfirmedPayment(id string, transactionStatus string) (*transactionEntities.Transaction, error) {
	transaction, err := usecase.transactionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(transaction)
	if *transaction == (transactionEntities.Transaction{}) {
		return nil, constants.ErrDataNotFound
	}

	if transaction.Status == transactionStatus {
		return transaction, nil
	}

	// Update Transaction Status
	transaction.Status = transactionStatus
	transactionResponse, err := usecase.transactionRepository.Update(transaction)

	// Add Balance Doctor
	doctorDB, err := usecase.doctorRepository.GetDoctorByID(int(transaction.Consultation.DoctorID))
	if err != nil {
		return nil, err
	}

	if transaction.Status == constants.Success {
		totalBalance := doctorDB.Amount + transaction.Price
		err = usecase.doctorRepository.UpdateAmount(transaction.Consultation.DoctorID, totalBalance)
		if err != nil {
			return nil, err
		}
	}
	return transactionResponse, nil
}

func (usecase *Transaction) FindAllByDoctorID(metadata *entities.Metadata, doctorID uint) (*[]transactionEntities.Transaction, error) {
	newTransaction, err := usecase.transactionRepository.FindAllByDoctorID(metadata, doctorID)
	if err != nil {
		return nil, err
	}
	if len(*newTransaction) == 0 {
		return nil, constants.ErrDataEmpty
	}
	return newTransaction, nil
}

func (usecase *Transaction) CountTransactionByDoctorID(doctorID uint) (int, error) {
	count, err := usecase.transactionRepository.CountTransactionByDoctorID(doctorID)
	if err != nil {
		return 0, err
	}
	return count, nil
}
