package user

import (
	userEntities "capstone/entities/user"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (userRepo *UserRepo) Register(user *userEntities.User) (userEntities.User, int64, error) {
	userDb := User {
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
	}

	var counterUsername, counterEmail int64
	err := userRepo.DB.Model(&userDb).Where("username = ?", userDb.Username).Count(&counterUsername).Error
	if err != nil {
		return userEntities.User{}, 0, err
	}

	if counterUsername > 0 {
		return userEntities.User{}, 1, nil
	}

	err = userRepo.DB.Model(&userDb).Where("email = ?", userDb.Email).Count(&counterEmail).Error
	if err != nil {
		return userEntities.User{}, 0, err
	}

	if counterEmail > 0 {
		return userEntities.User{}, 2, nil
	}

	err = userRepo.DB.Create(&userDb).Error
	if err != nil {
		fmt.Println(err)
		return userEntities.User{}, 0, err
	}

	userResult := userEntities.User {
		Id: userDb.Id,
		Username: userDb.Username,
		Email: userDb.Email,
		Password: userDb.Password,
	}

	return userResult, 0, nil
}

func (userRepo *UserRepo) Login(user *userEntities.User) (userEntities.User, error) {
	userDb := User {
		Username: user.Username,
		Password: user.Password,
	}

	password := userDb.Password

	err := userRepo.DB.Where("Username = ?", userDb.Username).First(&userDb).Error
	if err != nil {
		return userEntities.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(password))
	if err != nil {
		return userEntities.User{}, err
	}

	userResult := userEntities.User {
		Id: userDb.Id,
		Name: userDb.Name,
		Username: userDb.Username,
		Email: userDb.Email,
		Password: userDb.Password,
		Address: userDb.Address,
		Bio: userDb.Bio,
		PhoneNumber: userDb.PhoneNumber,
		Gender: userDb.Gender,
		Age: userDb.Age,
		ProfilePicture: userDb.ProfilePicture,
	}

	return userResult, nil
}

func (r *UserRepo) Create(email string, picture string, name string) (userEntities.User ,error) {
	var userDB User
	userDB.Email = email
	userDB.ProfilePicture = picture
	userDB.Name = name
	userDB.IsOauth = true
    
	err := r.DB.Create(&userDB).Error
	if err != nil {
		return userEntities.User{}, err
	}

	var userEnt userEntities.User
	userEnt.Id = userDB.Id
	userEnt.Name = userDB.Name
	userEnt.Email = userDB.Email
	userEnt.ProfilePicture = userDB.ProfilePicture
	userEnt.IsOauth = userDB.IsOauth

	return userEnt, nil
}

func (r *UserRepo) FindByEmail(email string) (userEntities.User, error) {
    var userDB User
    if err := r.DB.Where("email = ?", email).First(&userDB).Error; err != nil {
        return userEntities.User{}, err
    }

	var userEnt userEntities.User
	userEnt.Id = userDB.Id
	userEnt.Name = userDB.Name
	userEnt.Email = userDB.Email
	userEnt.ProfilePicture = userDB.ProfilePicture
	userEnt.IsOauth = userDB.IsOauth

    return userEnt, nil
}