package request

type ForumCreateRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	ImageUrl    string `json:"image_url" form:"image_url"`
}