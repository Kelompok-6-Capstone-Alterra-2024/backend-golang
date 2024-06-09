package consultation

import (
	"capstone/constants"
	"capstone/entities"
	chatEntities "capstone/entities/chat"
	consultationEntities "capstone/entities/consultation"
)

type ConsultationUseCase struct {
	consultationRepo consultationEntities.ConsultationRepository
	chatRepo chatEntities.RepositoryInterface
}

func NewConsultationUseCase(consultationRepo consultationEntities.ConsultationRepository, chatRepo chatEntities.RepositoryInterface) consultationEntities.ConsultationUseCase {
	return &ConsultationUseCase{
		consultationRepo: consultationRepo,
		chatRepo:         chatRepo,
	}
}

func (usecase *ConsultationUseCase) CreateConsultation(consultation *consultationEntities.Consultation) (*consultationEntities.Consultation, error) {
	result, err := usecase.consultationRepo.CreateConsultation(consultation)
	if err != nil {
		return nil, err
	}

	err = usecase.chatRepo.CreateChatRoom(result.ID)
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