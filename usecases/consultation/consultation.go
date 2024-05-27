package consultation

import consultationEntities "capstone/entities/consultation"

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

func (usecase *ConsultationUseCase) GetAllConsultation(userID int) (*[]consultationEntities.Consultation, error) {
	//TODO implement me
	panic("implement me")
}
