package complaint

import (
	consultationEntities "capstone/entities/consultation"
	userEntities "capstone/entities/user"
)

type Complaint struct {
	ID                 uint
	ComplaintID        uint
	Consultation       consultationEntities.Consultation
	UserID             int
	User               userEntities.User
	Name               string
	Age                int
	Gender             string
	Message            string
	MedicalHistory     string
	DoctorNotification string
}
