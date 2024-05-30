package response

import "time"

type ArticleCreatedResponse struct {
	ID        uint                 `json:"id"`
	DoctorID  uint                 `json:"doctor_id"`
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	ImageURL  string               `json:"image_url"`
	Date      time.Time            `json:"date"`
	ViewCount int                  `json:"view_count"`
	IsLiked   bool                 `json:"is_liked"`
	Doctor    DoctorGetAllResponse `json:"doctor"`
}

type DoctorGetAllResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
