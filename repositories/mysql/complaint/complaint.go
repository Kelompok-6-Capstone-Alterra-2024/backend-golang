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
	if err := repository.db.Table("Consultations").Where("id LIKE ?", complaint.ConsultationID).Update("complaint_id", complaintModel.ID).Error; err != nil {
		return nil, err
	}
	return complaintModel.ToEntities(), nil
}

func (repository *ComplaintRepo) GetAllByUserID(metadata *entities.Metadata, userID int) (*[]complaint.Complaint, error) {
	var complaintDB []Complaint
	if err := repository.db.Limit(metadata.Limit).Offset(metadata.Offset()).Where("user_id = ?", userID).Preload("Consultation").Find(&complaintDB).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}
	var complaints []complaint.Complaint
	for _, value := range complaintDB {
		complaints = append(complaints, *value.ToEntities())
	}
	return &complaints, nil
}

func (repository *ComplaintRepo) GetByID(complaintID int) (*complaint.Complaint, error) {
	var complaintDB Complaint
	if err := repository.db.Preload("Consultation").First(&complaintDB, "complaint_id LIKE ?", complaintID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}
	return complaintDB.ToEntities(), nil
}
