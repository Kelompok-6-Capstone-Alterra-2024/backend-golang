package response

type DoctorResponse struct {
	ID               uint   `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phone_number"`
	Gender           string `json:"gender"`
	IsAvailable      bool   `json:"is_available"`
	ProfilePicture   string `json:"profile_picture"`
	Balance          int    `json:"balance"`
	Experience       int    `json:"experience"`
	Almamater        string `json:"almamater"`
	GraduationYear   int    `json:"graduation_year"`
	PracticeLocation string `json:"practice_location"`
	PracticeCity     string `json:"practice_city"`
	PracticeProvince string `json:"practice_province"`
	StrNumber        string `json:"str_number"`
	Fee              int    `json:"fee"`
	Specialist       string `json:"specialist"`
	Amount           int    `json:"amount"`
}
