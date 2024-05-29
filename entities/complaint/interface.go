package complaint

import "capstone/entities"

type ComplaintRepository interface {
	Create(complaint *Complaint) (*Complaint, error)
	GetAllByUserID(metadata *entities.Metadata, userID int) (*[]Complaint, error)
	GetByID(complaintID int) (*Complaint, error)
}

type ComplaintUseCase interface {
	Create(complaint *Complaint) (*Complaint, error)
	GetAllByUserID(metadata *entities.Metadata, userID int) (*[]Complaint, error)
	GetByID(complaintID int) (*Complaint, error)
}
