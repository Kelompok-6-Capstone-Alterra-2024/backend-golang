package request

type ResetPasswordRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}