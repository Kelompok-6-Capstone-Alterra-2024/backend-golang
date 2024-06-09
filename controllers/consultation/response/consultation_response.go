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
	ID            int `json:"id"`
	Status        string
	PaymentStatus string
	IsAccepted    bool
	IsActive      bool
	Date          string
	Time          string
	Complaint     *complaintResponse.ComplaintResponse
}
