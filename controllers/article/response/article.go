package response

type ArticleCreatedResponse struct {
	ID       uint   `json:"id"`
	DoctorID uint   `json:"doctor_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
}
