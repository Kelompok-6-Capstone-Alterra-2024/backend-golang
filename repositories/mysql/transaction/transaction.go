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

func (repository *TransactionRepo) FindByID(ID string) (*transactionEntities.Transaction, error) {
	transactionDB := new(Transaction)
	if err := repository.db.Preload("Consultation").Preload("Consultation.Doctor").First(&transactionDB, "id LIKE ?", ID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}
	return transactionDB.ToEntities(), nil
}

func (repository *TransactionRepo) FindByConsultationID(consultationID uint) (*transactionEntities.Transaction, error) {
	transactionDB := new(Transaction)
	if err := repository.db.Preload("Consultation", "id = ?", consultationID).Preload("Consultation.Doctor").First(&transactionDB).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}
	return transactionDB.ToEntities(), nil
}

func (repository *TransactionRepo) FindAll(metadata *entities.Metadata, userID uint) (*[]transactionEntities.Transaction, error) {
	transactionDB := new([]Transaction)
	if err := repository.db.
		Joins("JOIN consultations ON consultations.id = transactions.consultation_id").
		Joins("JOIN users ON consultations.user_id = users.id").
		Where("users.id LIKE ?", userID).
		Preload("Consultation").
		Preload("Consultation.Doctor").
		Limit(metadata.Limit).
		Offset(metadata.Offset()).
		Find(&transactionDB).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}
	var transactions []transactionEntities.Transaction
	for _, transaction := range *transactionDB {
		transactions = append(transactions, *transaction.ToEntities())
	}
	return &transactions, nil
}

func (repository *TransactionRepo) Update(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	transactionDB := ToTransactionModel(transaction)
	if err := repository.db.Model(&Transaction{}).Where("id LIKE ?", transactionDB.ID).Update("status", transactionDB.Status).Error; err != nil {
		return nil, constants.ErrInsertDatabase
	}
	return transactionDB.ToEntities(), nil
}

func (repository *TransactionRepo) Delete(ID uint) error {
	//TODO implement me
	panic("implement me")
}
