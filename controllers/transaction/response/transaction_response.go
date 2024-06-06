package response

import "capstone/controllers/consultation/response"

type TransactionResponse struct {
	ID           string                            `json:"id"`
	Consultation response.ConsultationUserResponse `json:"consultation"`
	Price        int                               `json:"price"`
	SnapURL      string                            `json:"snap_url"`
	Status       string                            `json:"status"`
}
