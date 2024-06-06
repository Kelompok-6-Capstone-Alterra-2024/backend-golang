package user

import (
	"capstone/constants"
	userEntitites "capstone/entities/user"
	"capstone/middlewares"
	"context"

	"golang.org/x/crypto/bcrypt"
	myoauth "golang.org/x/oauth2"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type UserUseCase struct {
	repository userEntitites.RepositoryInterface
	oauthConfig *myoauth.Config
}

func NewUserUseCase(repository userEntitites.RepositoryInterface, oauthConfig *myoauth.Config) *UserUseCase {
	return &UserUseCase{
		repository: repository,
		oauthConfig: oauthConfig,
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

func (u *UserUseCase) HandleGoogleLogin() string {
    return u.oauthConfig.AuthCodeURL("state-token", myoauth.AccessTypeOffline)
}

func (u *UserUseCase) HandleGoogleCallback(ctx context.Context, code string) (userEntitites.User, error) {
    token, err := u.oauthConfig.Exchange(ctx, code)
    if err != nil {
        return userEntitites.User{}, constants.ErrExcange
    }

    // Membuat layanan OAuth2
    oauth2Service, err := oauth2.NewService(ctx, option.WithTokenSource(u.oauthConfig.TokenSource(ctx, token)))
    if err != nil {
        return userEntitites.User{}, constants.ErrNewServiceGoogle
    }

    userInfoService := oauth2.NewUserinfoV2MeService(oauth2Service)
    userInfo, err := userInfoService.Get().Do()
    if err != nil {
        return userEntitites.User{}, constants.ErrNewUserInfo
    }

    // Cek apakah pengguna sudah ada di database
	result, myCode, err := u.repository.OauthFindByEmail(userInfo.Email)
    if err != nil && myCode == 0 {
        newUser, err := u.repository.Create(userInfo.Email, userInfo.Picture, userInfo.Name)
		if err != nil {
            return userEntitites.User{}, constants.ErrInsertOAuth
        }

		tokenJWT, _ := middlewares.CreateToken(newUser.Id)
		newUser.Token = tokenJWT

		return newUser, nil
    }

	if err != nil && myCode == 1 {
		return userEntitites.User{}, err
	}

	tokenJWT, _ := middlewares.CreateToken(result.Id)
	result.Token = tokenJWT

    return result, nil
}