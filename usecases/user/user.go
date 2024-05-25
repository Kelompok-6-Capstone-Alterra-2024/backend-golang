package user

import (
	"capstone/constants"
	userEntitites "capstone/entities/user"
	"capstone/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repository userEntitites.RepositoryInterface
}

func NewUserUseCase(repository userEntitites.RepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (userUseCase *UserUseCase) Register(user *userEntitites.User) (userEntitites.User, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return userEntitites.User{}, constants.ErrEmptyInputUser
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return userEntitites.User{}, constants.ErrHashedPassword
	}

	user.Password = string(hashedPassword)
	
	var kode int64
	userResult, kode, err := userUseCase.repository.Register(user)
	if err != nil {
		return userEntitites.User{}, constants.ErrInsertDatabase
	}

	if kode == 1 {
		return userEntitites.User{}, constants.ErrUsernameAlreadyExist
	}

	if kode == 2 {
		return userEntitites.User{}, constants.ErrEmailAlreadyExist
	}

	token, _ := middlewares.CreateToken(userResult.Id)
	userResult.Token = token

	return userResult, nil
}

func (userUseCase *UserUseCase) Login(user *userEntitites.User) (userEntitites.User, error) {
	if user.Username == "" || user.Password == "" {
		return userEntitites.User{}, constants.ErrEmptyInputLogin
	}

	userResult, err := userUseCase.repository.Login(user)
	if err != nil {
		return userEntitites.User{}, constants.ErrUserNotFound
	}

	token, _ := middlewares.CreateToken(userResult.Id)
	userResult.Token = token

	return userResult, nil
}