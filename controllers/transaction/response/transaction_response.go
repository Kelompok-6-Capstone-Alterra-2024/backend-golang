package response

import "capstone/controllers/consultation/response"

type TransactionResponse struct {
	ID           string                            `json:"id"`
	Consultation response.ConsultationUserResponse `json:"consultation"`
	Price        int                               `json:"price"`
	PaymentType  string                            `json:"payment_type"`
	PaymentLink  string                            `json:"payment_link"`
	Bank         string                            `json:"bank"`
	Status       string                            `json:"status"`
	CreatedAt    string                            `json:"created_at"`
	UpdatedAt    string                            `json:"updated_at"`
}
