package complaint

import (
	"capstone/constants"
	"capstone/entities"
	"capstone/entities/complaint"
	"gorm.io/gorm"
)

type ComplaintRepo struct {
	db *gorm.DB
}

func NewComplaintRepo(db *gorm.DB) complaint.ComplaintRepository {
	return &ComplaintRepo{db}
}

func (repository *ComplaintRepo) Create(complaint *complaint.Complaint) (*complaint.Complaint, error) {
	complaintModel := ToComplaintModel(complaint)
	if err := repository.db.Create(&complaintModel).Error; err != nil {
		return nil, constants.ErrInsertDatabase
	}
	if err := repository.db.Table("consultations").Where("id LIKE ?", complaintModel.ID).Update("complaint_id", complaintModel.ID).Error; err != nil {
		return nil, err
	}
	return complaintModel.ToEntities(), nil
}

func (repository *ComplaintRepo) GetAllByUserID(metadata *entities.Metadata, userID int) (*[]complaint.Complaint, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ComplaintRepo) GetByID(complaintID int) (*complaint.Complaint, error) {
	//TODO implement me
	panic("implement me")
}
