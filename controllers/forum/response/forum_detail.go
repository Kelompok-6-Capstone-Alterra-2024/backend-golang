package response

type ForumDetailResponse struct {
	ForumID     uint           `json:"forum_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ImageUrl    string         `json:"image_url"`
	Post        []PostResponse `json:"post"`
}

type PostResponse struct {
	PostID   uint             `json:"post_id"`
	Content  string           `json:"content"`
	ImageUrl string           `json:"image_url"`
	User     UserPostResponse `json:"user"`
}

type UserPostResponse struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	ImageUrl string `json:"image_url"`
}