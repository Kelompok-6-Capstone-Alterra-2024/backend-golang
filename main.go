package main

import (
	"capstone/configs"
	consultationController "capstone/controllers/consultation"
	doctorController "capstone/controllers/doctor"
	storyController "capstone/controllers/story"
	userController "capstone/controllers/user"
	"capstone/repositories/mysql"
	consultationRepositories "capstone/repositories/mysql/consultation"
	doctorRepositories "capstone/repositories/mysql/doctor"
	storyRepositories "capstone/repositories/mysql/story"
	userRepositories "capstone/repositories/mysql/user"
	"capstone/routes"
	consultationUseCase "capstone/usecases/consultation"
	doctorUseCase "capstone/usecases/doctor"
	storyUseCase "capstone/usecases/story"
	userUseCase "capstone/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())

	userRepo := userRepositories.NewUserRepo(db)
	doctorRepo := doctorRepositories.NewDoctorRepo(db)
	consultationRepo := consultationRepositories.NewConsultationRepo(db)
	StoryRepo := storyRepositories.NewStoryRepo(db)

	userUC := userUseCase.NewUserUseCase(userRepo)
	doctorUC := doctorUseCase.NewDoctorUseCase(doctorRepo)
	consultationUC := consultationUseCase.NewConsultationUseCase(consultationRepo)
	storyUC := storyUseCase.NewStoryUseCase(StoryRepo)

	userCont := userController.NewUserController(userUC)
	doctorCont := doctorController.NewDoctorController(doctorUC)
	consultationCont := consultationController.NewConsultationController(consultationUC)
	storyCont := storyController.NewStoryController(storyUC)

	route := routes.NewRoute(userCont, doctorCont, consultationCont, storyCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
