package consultation

import (
	"capstone/constants"
	consultationEntities "capstone/entities/consultation"
	"gorm.io/gorm"
)

type ConsultationRepo struct {
	db *gorm.DB
}

func NewConsultationRepo(db *gorm.DB) consultationEntities.ConsultationRepository {
	return &ConsultationRepo{
		db: db,
	}
}

func (repository *ConsultationRepo) CreateConsultation(consultation *consultationEntities.Consultation) (*consultationEntities.Consultation, error) {
	consultationRequest := ToConsultationModel(consultation)

	if err := repository.db.Create(&consultationRequest).Error; err != nil {
		return nil, constants.ErrInsertDatabase
	}

	consultationResult := consultationRequest.ToEntities()
	return consultationResult, nil
}

func (repository *ConsultationRepo) GetConsultationByID(consultationID int) (consultation *consultationEntities.Consultation, err error) {
	var consultationResult Consultation
	if err := repository.db.Preload("User").Preload("Doctor").First(&consultationResult, consultationID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	return consultationResult.ToEntities(), nil
}

func (repository *ConsultationRepo) GetAllConsultation(userID int) (*[]consultationEntities.Consultation, error) {
	//TODO implement me
	panic("implement me")
}
