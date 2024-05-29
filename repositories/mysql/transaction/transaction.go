package transaction

import (
	"capstone/constants"
	"capstone/entities"
	transactionEntities "capstone/entities/transaction"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) transactionEntities.TransactionRepository {
	return &TransactionRepo{db}
}

func (repository *TransactionRepo) Insert(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	transactionDb := ToTransactionModel(transaction)
	if err := repository.db.Create(&transactionDb).Error; err != nil {
		return nil, constants.ErrInsertDatabase
	}

	if err := repository.db.Preload("Consultation").Preload("Consultation.Doctor").First(&transactionDb, transactionDb.ID).Error; err != nil {
		return nil, constants.ErrInsertDatabase
	}

	return transactionDb.ToEntities(), nil
}

func (repository *TransactionRepo) FindByID(ID uint) (*transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *TransactionRepo) FindByConsultationID(consultationID uint) (*transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *TransactionRepo) FindAll(metadata *entities.Metadata, userID uint) (*[]transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *TransactionRepo) Update(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *TransactionRepo) Delete(ID uint) error {
	//TODO implement me
	panic("implement me")
}
