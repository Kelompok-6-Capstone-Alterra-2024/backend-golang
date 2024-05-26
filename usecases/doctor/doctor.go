package doctor

import (
	"capstone/constants"
	"capstone/entities"
	doctorEntities "capstone/entities/doctor"
	"capstone/middlewares"
	"golang.org/x/crypto/bcrypt"
)

type DoctorUseCase struct {
	doctorRepository doctorEntities.DoctorRepositoryInterface
}

func NewDoctorUseCase(doctorRepository doctorEntities.DoctorRepositoryInterface) doctorEntities.DoctorUseCaseInterface {
	return &DoctorUseCase{
		doctorRepository: doctorRepository,
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
