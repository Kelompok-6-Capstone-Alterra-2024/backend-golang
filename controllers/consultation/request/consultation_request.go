package request

import (
	"capstone/entities/consultation"
	"time"
)

type ConsultationRequest struct {
	DoctorID uint `json:"doctor_id" form:"doctor_id" binding:"required"`
	UserID   int
	Date     string `json:"date" form:"date" binding:"required"`
	Time     string `json:"time" form:"time" binding:"required"`
}

func (r ConsultationRequest) ToEntities(date, time time.Time) *consultation.Consultation {
	return &consultation.Consultation{
		DoctorID: r.DoctorID,
		UserID:   uint(r.UserID),
		Date:     date,
		Time:     time,
	}
}
