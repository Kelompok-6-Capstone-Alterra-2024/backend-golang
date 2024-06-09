package consultation

import (
	"capstone/constants"
	"capstone/entities"
	consultationEntities "capstone/entities/consultation"
	"fmt"
	"gorm.io/gorm"
	"time"
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
	fmt.Println(consultationRequest.Time)

	if err := repository.db.Create(&consultationRequest).Error; err != nil {
		return nil, constants.ErrInsertDatabase
	}

	consultationResult, err := consultationRequest.ToEntities()
	if err != nil {
		return nil, constants.ErrInputTime
	}
	return consultationResult, nil
}

func (repository *ConsultationRepo) GetConsultationByID(consultationID int) (consultation *consultationEntities.Consultation, err error) {
	var consultationDB Consultation
	if err = repository.db.Preload("User").Preload("Doctor").First(&consultationDB, consultationID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}
	consultationResult, err := consultationDB.ToEntities()
	if err != nil {
		return nil, constants.ErrInputTime
	}

	return consultationResult, nil
}

func (repository *ConsultationRepo) GetAllUserConsultation(metadata *entities.Metadata, userID int) (*[]consultationEntities.Consultation, error) {
	var consultationDB []Consultation

	if err := repository.db.Limit(metadata.Limit).Offset(metadata.Offset()).Preload("Doctor").Preload("Complaint").Find(&consultationDB, "user_id LIKE ?", userID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	var consultations []consultationEntities.Consultation
	for _, consultation := range consultationDB {
		consultationResult, err := consultation.ToEntities()
		if err != nil {
			return nil, constants.ErrInputTime
		}
		consultations = append(consultations, *consultationResult)
	}

	return &consultations, nil
}

func (repository *ConsultationRepo) UpdateStatusConsultation(consultation *consultationEntities.Consultation) (*consultationEntities.Consultation, error) {
	var consultationDB Consultation
	if err := repository.db.First(&consultationDB, "id LIKE ?", consultation.ID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	consultationDB.Status = consultation.Status

	if err := repository.db.Model(&consultationDB).Save(&consultationDB).Error; err != nil {
		return nil, err
	}

	consultationResult, err := consultationDB.ToEntities()
	if err != nil {
		return nil, constants.ErrInputTime
	}
	return consultationResult, nil
}

func (repository *ConsultationRepo) GetAllDoctorConsultation(metadata *entities.Metadata, doctorID int) (*[]consultationEntities.Consultation, error) {
	var consultationbDB []Consultation

	if err := repository.db.Limit(metadata.Limit).Offset(metadata.Offset()).Preload("Complaint").Find(&consultationbDB, "doctor_id LIKE ?", doctorID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	var consultations []consultationEntities.Consultation
	for _, consultation := range consultationbDB {
		consultationResult, err := consultation.ToEntities()
		if err != nil {
			return nil, constants.ErrInputTime
		}
		consultations = append(consultations, *consultationResult)
	}

	return &consultations, nil
}

func (repository *ConsultationRepo) CountConsultationToday(doctorID int) (int64, error) {
	var count int64
	if err := repository.db.Model(&Consultation{}).Where("doctor_id LIKE ? AND date LIKE ?", doctorID, time.Now().Format("2006-11-22")).Count(&count).Error; err != nil {
		return 0, constants.ErrDataNotFound
	}
	return count, nil
}

func (repository *ConsultationRepo) CountConsultationByDoctorID(doctorID int) (int64, error) {
	var count int64
	if err := repository.db.Model(&Consultation{}).Where("doctor_id LIKE ? AND status NOT LIKE ?", doctorID, constants.REJECTED).Count(&count).Error; err != nil {
		return 0, constants.ErrDataNotFound
	}
	return count, nil
}

func (repository *ConsultationRepo) CountConsultationByStatus(doctorID int, status string) (int64, error) {
	var count int64
	if err := repository.db.Model(&Consultation{}).Where("doctor_id LIKE ? AND status LIKE ?", doctorID, status).Count(&count).Error; err != nil {
		return 0, constants.ErrDataNotFound
	}
	return count, nil
}
