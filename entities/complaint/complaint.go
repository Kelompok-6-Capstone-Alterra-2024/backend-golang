package complaint

import (
	"capstone/controllers/complaint/response"
)

type Complaint struct {
	ID                 uint
	ConsultationID     uint
	Name               string
	Age                int
	Gender             string
	Message            string
	MedicalHistory     string
	DoctorNotification string
}

func (r *Complaint) ToResponse() *response.ComplaintResponse {
	return &response.ComplaintResponse{
		ID:             r.ID,
		Name:           r.Name,
		Age:            r.Age,
		Gender:         r.Gender,
		MedicalHistory: r.MedicalHistory,
	}
}