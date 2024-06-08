package consultation

import (
	"capstone/constants"
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

func (usecase *ConsultationUseCase) GetAllConsultation(metadata *entities.Metadata, userID int) (*[]consultationEntities.Consultation, error) {
	result, err := usecase.consultationRepo.GetAllConsultation(metadata, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *ConsultationUseCase) CreateConsultationNotes(consultationNotes consultationEntities.ConsultationNotes) (consultationEntities.ConsultationNotes, error) {
	if consultationNotes.ConsultationID == 0 {
		return consultationEntities.ConsultationNotes{}, constants.ErrInvalidConsultationID
	}
	
	result, err := usecase.consultationRepo.CreateConsultationNotes(consultationNotes)
	if err != nil {
		return result, err
	}
	
	return result, nil
}

func (usecase *ConsultationUseCase) GetConsultationNotesByID(consultationID int) (consultationEntities.ConsultationNotes, error) {
	result, err := usecase.consultationRepo.GetConsultationNotesByID(consultationID)
	if err != nil {
		return result, err
	}
	return result, nil
}