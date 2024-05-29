package response

type MusicResponse struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Singer    string `json:"singer"`
	MusicUrl  string `json:"music_url"`
	ImageUrl  string `json:"image_url"`
	ViewCount int    `json:"view_count"`
	IsLiked   bool   `json:"is_liked"`
}