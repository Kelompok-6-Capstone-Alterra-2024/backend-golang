package response

import (
	doctorResponse "capstone/controllers/doctor/response"
	"time"
)

type ConsultationResponse struct {
	ID            int `json:"id"`
	Doctor        *doctorResponse.DoctorResponse
	Status        string
	PaymentStatus string
	IsAccepted    bool
	IsActive      bool
	Date          time.Time
}
