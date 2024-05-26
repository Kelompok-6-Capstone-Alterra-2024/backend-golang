package response

type DoctorResponse struct {
	ID               uint
	Username         string
	Email            string
	Name             string
	Address          string
	PhoneNumber      string
	Gender           string
	IsAvailable      bool
	ProfilePicture   string
	Balance          int
	Experience       int
	Almamater        string
	GraduationYear   int
	PracticeLocation string
	PracticeCity     string
	PracticeProvince string
	StrNumber        string
	Fee              int
	Specialist       string
}
