package doctor

import (
	"capstone/constants"
	"capstone/entities"
	doctorEntities "capstone/entities/doctor"
	"capstone/middlewares"
	"context"

	"golang.org/x/crypto/bcrypt"
	myoauth "golang.org/x/oauth2"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type DoctorUseCase struct {
	doctorRepository doctorEntities.DoctorRepositoryInterface
	oauthConfig *myoauth.Config
}

func NewDoctorUseCase(doctorRepository doctorEntities.DoctorRepositoryInterface, oauthConfig *myoauth.Config) doctorEntities.DoctorUseCaseInterface {
	return &DoctorUseCase{
		doctorRepository: doctorRepository,
		oauthConfig: oauthConfig,
	}
}

func (usecase *DoctorUseCase) Register(doctor *doctorEntities.Doctor) (*doctorEntities.Doctor, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(doctor.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, constants.ErrHashedPassword
	}

	doctor.Password = string(hashedPassword)

	doctorResult, err := usecase.doctorRepository.Register(doctor)
	if err != nil {
		return nil, err
	}
	token, err := middlewares.CreateToken(int(doctorResult.ID))
	if err != nil {
		return nil, err
	}
	doctorResult.Token = token
	return doctorResult, nil

}

func (usecase *DoctorUseCase) Login(doctor *doctorEntities.Doctor) (*doctorEntities.Doctor, error) {
	if (doctor.Email == "" && doctor.Password == "") || (doctor.Username == "" && doctor.Password == "") {
		return nil, constants.ErrEmptyInputLogin
	}
	userResult, err := usecase.doctorRepository.Login(doctor)
	if err != nil {
		return nil, err
	}
	token, err := middlewares.CreateToken(int(userResult.ID))
	if err != nil {
		return nil, err
	}
	userResult.Token = token
	return userResult, nil
}

func (usecase *DoctorUseCase) GetDoctorByID(doctorID int) (*doctorEntities.Doctor, error) {
	result, err := usecase.doctorRepository.GetDoctorByID(doctorID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *DoctorUseCase) GetAllDoctor(metadata *entities.Metadata) (*[]doctorEntities.Doctor, error) {
	result, err := usecase.doctorRepository.GetAllDoctor(metadata)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (usecase *DoctorUseCase) GetActiveDoctor(metadata *entities.Metadata) (*[]doctorEntities.Doctor, error) {
	result, err := usecase.doctorRepository.GetActiveDoctor(metadata)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *DoctorUseCase) HandleGoogleLogin() string {
    return u.oauthConfig.AuthCodeURL("state-token", myoauth.AccessTypeOffline)
}

func (u *DoctorUseCase) HandleGoogleCallback(ctx context.Context, code string) (doctorEntities.Doctor, error) {
    token, err := u.oauthConfig.Exchange(ctx, code)
    if err != nil {
        return doctorEntities.Doctor{}, constants.ErrExcange
    }

    // Membuat layanan OAuth2
    oauth2Service, err := oauth2.NewService(ctx, option.WithTokenSource(u.oauthConfig.TokenSource(ctx, token)))
    if err != nil {
        return doctorEntities.Doctor{}, constants.ErrNewServiceGoogle
    }

    userInfoService := oauth2.NewUserinfoV2MeService(oauth2Service)
    userInfo, err := userInfoService.Get().Do()
    if err != nil {
        return doctorEntities.Doctor{}, constants.ErrNewUserInfo
    }

    // Cek apakah pengguna sudah ada di database
	result, myCode, err := u.doctorRepository.OauthFindByEmail(userInfo.Email)
    if err != nil && myCode == 0 {
        newUser, err := u.doctorRepository.Create(userInfo.Email, userInfo.Picture, userInfo.Name)
		if err != nil {
            return doctorEntities.Doctor{}, constants.ErrInsertOAuth
        }

		tokenJWT, _ := middlewares.CreateToken(int(newUser.ID))
		newUser.Token = tokenJWT

		return newUser, nil
    }

	if err != nil && myCode == 1 {
		return doctorEntities.Doctor{}, err
	}

	tokenJWT, _ := middlewares.CreateToken(int(result.ID))
	result.Token = tokenJWT

    return result, nil
}