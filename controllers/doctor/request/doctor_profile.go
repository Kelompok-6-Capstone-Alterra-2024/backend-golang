package request

type UpdateDoctorProfileRequest struct {
	Username         string `json:"username" form:"username"`
	Name             string `json:"name" form:"name"`
	Address          string `json:"address" form:"address"`
	PhoneNumber      string `json:"phone_number" form:"phone_number"`
	Gender           string `json:"gender" form:"gender"`
	ProfilePicture   string `json:"profile_picture" form:"profile_picture"`
	Experience       int    `json:"experience" form:"experience"`
	Almamater        string `json:"almamater" form:"almamater"`
	GraduationYear   int    `json:"graduation_year" form:"graduation_year"`
	PracticeLocation string `json:"practice_location" form:"practice_location"`
	PracticeCity     string `json:"practice_city" form:"practice_city"`
	PracticeProvince string `json:"practice_province" form:"practice_province"`
	StrNumber        string `json:"str_number" form:"str_number"`
	Fee              int    `json:"fee" form:"fee"`
	Specialist       string `json:"specialist" form:"specialist"`
}
