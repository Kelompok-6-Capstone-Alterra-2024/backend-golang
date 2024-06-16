package request

import "capstone/entities/consultation"

type ConsultationStatusUpdateRequest struct {
	ID       uint
	Status   string `json:"status"`
	DoctorID uint
}

func (r *ConsultationStatusUpdateRequest) ToEntities() *consultation.Consultation {
	return &consultation.Consultation{
		ID:       r.ID,
		Status:   r.Status,
		DoctorID: r.DoctorID,
	}
}
