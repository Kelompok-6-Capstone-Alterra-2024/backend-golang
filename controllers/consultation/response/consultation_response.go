package response

import (
	complaintResponse "capstone/controllers/complaint/response"
	doctorResponse "capstone/controllers/doctor/response"
)

type ConsultationUserResponse struct {
	ID            int `json:"id"`
	Doctor        *doctorResponse.DoctorResponse
	Status        string
	PaymentStatus string
	IsAccepted    bool
	IsActive      bool
	Date          string
	Time          string
}

type ConsultationDoctorResponse struct {
	ID            int                                  `json:"id"`
	Status        string                               `json:"status"`
	PaymentStatus string                               `json:"payment_status"`
	IsAccepted    bool                                 `json:"is_accepted"`
	IsActive      bool                                 `json:"is_active"`
	Date          string                               `json:"date"`
	Time          string                               `json:"time"`
	Complaint     *complaintResponse.ComplaintResponse `json:"complaint"`
}
