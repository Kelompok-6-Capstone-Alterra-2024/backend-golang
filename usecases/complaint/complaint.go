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

func (usecase *ComplaintUseCase) GetAllByUserID(metadata *entities.Metadata, userID int) (*[]complaintEntities.Complaint, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *ComplaintUseCase) GetByID(complaintID int) (*complaintEntities.Complaint, error) {
	//TODO implement me
	panic("implement me")
}
