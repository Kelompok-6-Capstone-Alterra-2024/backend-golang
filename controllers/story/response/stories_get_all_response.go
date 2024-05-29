package response

import "time"

type StoriesGetAllResponse struct {
	ID        uint                 `json:"id"`
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	Date      time.Time            `json:"date"`
	ImageUrl  string               `json:"image_url"`
	ViewCount int                  `json:"view_count"`
	IsLiked   bool                 `json:"is_liked"`
	Doctor    DoctorGetAllResponse `json:"doctor"`
}

type DoctorGetAllResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}