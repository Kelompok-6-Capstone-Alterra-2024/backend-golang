package consultation

import (
	"capstone/entities"
	consultationEntities "capstone/entities/consultation"
	"github.com/go-playground/validator/v10"
)

type ConsultationUseCase struct {
	consultationRepo consultationEntities.ConsultationRepository
	validate         *validator.Validate
}

func NewConsultationUseCase(consultationRepo consultationEntities.ConsultationRepository, validate *validator.Validate) consultationEntities.ConsultationUseCase {
	return &ConsultationUseCase{
		consultationRepo: consultationRepo,
	}
}

func (usecase *ConsultationUseCase) CreateConsultation(consultation *consultationEntities.Consultation) (*consultationEntities.Consultation, error) {
	if err := usecase.validate.Struct(consultation); err != nil {
		return nil, err
	}
	result, err := usecase.consultationRepo.CreateConsultation(consultation)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *ConsultationUseCase) GetConsultationByID(consultationID int) (*consultationEntities.Consultation, error) {
	result, err := usecase.consultationRepo.GetConsultationByID(consultationID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *ConsultationUseCase) GetAllUserConsultation(metadata *entities.Metadata, userID int) (*[]consultationEntities.Consultation, error) {
	result, err := usecase.consultationRepo.GetAllUserConsultation(metadata, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *ConsultationUseCase) UpdateStatusConsultation(consultation *consultationEntities.Consultation) (*consultationEntities.Consultation, error) {
	result, err := usecase.consultationRepo.UpdateStatusConsultation(consultation)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *ConsultationUseCase) GetAllDoctorConsultation(metadata *entities.Metadata, doctorID int) (*[]consultationEntities.Consultation, error) {
	result, err := usecase.consultationRepo.GetAllDoctorConsultation(metadata, doctorID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *ConsultationUseCase) CountConsultationByDoctorID(doctorID int) (int64, error) {
	result, err := usecase.consultationRepo.CountConsultationByDoctorID(doctorID)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (usecase *ConsultationUseCase) CountConsultationToday(doctorID int) (int64, error) {
	result, err := usecase.consultationRepo.CountConsultationToday(doctorID)
	if err != nil {
		return 0, nil
	}
	return result, nil
}

func (usecase *ConsultationUseCase) CountConsultationByStatus(doctorID int, status string) (int64, error) {
	result, err := usecase.consultationRepo.CountConsultationByStatus(doctorID, status)
	if err != nil {
		return 0, err
	}
	return result, nil
}
