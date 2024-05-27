package request

type ConsultationRequest struct {
	DoctorID uint `json:"doctor_id" form:"doctor_id" binding:"required"`
	UserID   int
}
