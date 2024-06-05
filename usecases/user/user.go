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
    result, err := u.repository.FindByEmail(userInfo.Email)
    if err != nil {
        // Jika pengguna belum ada, buat pengguna baru
        result = userEntitites.User{
            Email:          userInfo.Email,
            Name:           userInfo.Name,
            ProfilePicture: userInfo.Picture,
            IsOauth:        true,
        }
        res, err := u.repository.Create(result)
		if  err != nil {
            return result, constants.ErrServer
        }

		tokenJWT, _ := middlewares.CreateToken(res.Id)
		res.Token = tokenJWT

		return res, nil
    }

	tokenJWT, _ := middlewares.CreateToken(result.Id)
	result.Token = tokenJWT

    return result, nil
}