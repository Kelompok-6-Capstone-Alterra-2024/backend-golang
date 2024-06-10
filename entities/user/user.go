package user

import "context"

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
	IsOauth        bool
	Points         int
}

type RepositoryInterface interface {
	Register(user *User) (User, int64, error)
	Login(user *User) (User, error)
	Create(email string, picture string, name string, username string) (User, error)
	OauthFindByEmail(email string) (User, int, error)
	GetPointsByUserId(id int) (int, error)
	ResetPassword(email string, password string) error
}

type UseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
	HandleGoogleLogin() string
	HandleGoogleCallback(ctx context.Context, code string) (User, error)
	GetPointsByUserId(id int) (int, error)
	ResetPassword(email string, password string) error
	HandleFacebookLogin() string
	HandleFacebookCallback(ctx context.Context, code string) (User, error)
}
