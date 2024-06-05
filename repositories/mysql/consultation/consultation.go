package consultation

import (
	"capstone/constants"
	"capstone/entities"
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
	if err = repository.db.Preload("User").Preload("Doctor").First(&consultationResult, consultationID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	return consultationResult.ToEntities(), nil
}

func (repository *ConsultationRepo) GetAllConsultation(metadata *entities.Metadata, userID int) (*[]consultationEntities.Consultation, error) {
	var consultationResult []Consultation

	if err := repository.db.Limit(metadata.Limit).Offset((metadata.Page-1)*metadata.Limit).Preload("Doctor").Find(&consultationResult, "user_id LIKE ?", userID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	var consultations []consultationEntities.Consultation
	for _, consultation := range consultationResult {
		consultations = append(consultations, *consultation.ToEntities())
	}

	return &consultations, nil
}

func (repository *ConsultationRepo) UpdateStatusConsultation(consultation *consultationEntities.Consultation) (*consultationEntities.Consultation, error) {
	consultationDB := ToConsultationModel(consultation)
	if err := repository.db.First(&consultationDB, consultationDB.ID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	if err := repository.db.Model(consultationDB).Update("status", consultationDB.Status).Error; err != nil {
		return nil, err
	}

	return consultationDB.ToEntities(), nil
}
