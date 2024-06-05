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
}

type RepositoryInterface interface {
	Register(user *User) (User, int64, error)
	Login(user *User) (User, error)
	Create(user User) (User ,error)
	FindByEmail(email string) (User, error)
}

type UseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
	HandleGoogleLogin() string
	HandleGoogleCallback(ctx context.Context, code string) (User, error)
}