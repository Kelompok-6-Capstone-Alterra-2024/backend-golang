package request

import "capstone/entities/consultation"

type ConsultationStatusUpdateRequest struct {
	ID     uint   `validate:"required"`
	Status string `json:"status" validate:"required"`
}

func (r *ConsultationStatusUpdateRequest) ToEntities() *consultation.Consultation {
	return &consultation.Consultation{
		ID:     r.ID,
		Status: r.Status,
	}
}
