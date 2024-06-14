package complaint

import (
	"capstone/entities"
	complaintEntities "capstone/entities/complaint"
)

type ComplaintUseCase struct {
	complaintRepo complaintEntities.ComplaintRepository
}

func NewComplaintUseCase(complaintRepo complaintEntities.ComplaintRepository) complaintEntities.ComplaintUseCase {
	return &ComplaintUseCase{complaintRepo: complaintRepo}
}

func (usecase *ComplaintUseCase) Create(complaint *complaintEntities.Complaint) (*complaintEntities.Complaint, error) {
	result, err := usecase.complaintRepo.Create(complaint)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *ComplaintUseCase) GetAllByDoctorID(metadata *entities.Metadata, doctorID int) (*[]complaintEntities.Complaint, error) {
	result, err := usecase.complaintRepo.GetAllByDoctorID(metadata, doctorID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *ComplaintUseCase) GetByID(complaintID int) (*complaintEntities.Complaint, error) {
	result, err := usecase.complaintRepo.GetByID(complaintID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *ComplaintUseCase) SearchComplaintByPatientName(metadata *entities.Metadata, name string, doctorID uint) (*[]complaintEntities.Complaint, error) {
	result, err := usecase.complaintRepo.SearchComplaintByPatientName(metadata, name, doctorID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
