package response

type PostResponse struct {
	ID       uint             `json:"id"`
	Content  string           `json:"content"`
	ImageUrl string           `json:"image_url"`
	User     UserPostResponse `json:"user"`
}

type UserPostResponse struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
}