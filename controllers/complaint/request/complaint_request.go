package request

import "capstone/entities/complaint"

type ComplaintRequest struct {
	ConsultationID uint   `json:"consultation_id" form:"consultation_id"`
	Name           string `json:"name" form:"name"`
	Age            int    `json:"age" form:"age"`
	Gender         string `json:"gender" form:"gender"`
	Message        string `json:"message" form:"message"`
	MedicalHistory string `json:"medical_history" form:"medical_history"`
}

func (r *ComplaintRequest) ToEntities() *complaint.Complaint {
	return &complaint.Complaint{
		Name:           r.Name,
		Age:            r.Age,
		Gender:         r.Gender,
		Message:        r.Message,
		MedicalHistory: r.MedicalHistory,
		ConsultationID: r.ConsultationID,
	}
}
