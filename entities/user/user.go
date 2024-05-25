package user

type User struct {
	Id             int
	Name           string
	Username       string
	Email          string
	Password       string
	Address        string
	Bio            string
	PhoneNumber    string
	Gender         string
	Age            int
	ProfilePicture string
	Token          string
}

type RepositoryInterface interface {
	Register(user *User) (User, int64, error)
	Login(user *User) (User, error)
}

type UseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
}