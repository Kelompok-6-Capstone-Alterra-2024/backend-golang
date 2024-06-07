package consultation

import (
	"capstone/entities"
	consultationEntities "capstone/entities/consultation"
)

type ConsultationUseCase struct {
	consultationRepo consultationEntities.ConsultationRepository
}

func NewConsultationUseCase(consultationRepo consultationEntities.ConsultationRepository) consultationEntities.ConsultationUseCase {
	return &ConsultationUseCase{
		consultationRepo: consultationRepo,
	}
}

func (usecase *ConsultationUseCase) CreateConsultation(consultation *consultationEntities.Consultation) (*consultationEntities.Consultation, error) {
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
