package response

type ForumGetDoctorResponse struct {
	Name            string `json:"name"`
	ImageUrl        string `json:"image_url"`
	NumberOfMembers int    `json:"number_of_members"`
}