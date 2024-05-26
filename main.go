package main

import (
	"capstone/configs"
	doctorController "capstone/controllers/doctor"
	userController "capstone/controllers/user"
	"capstone/repositories/mysql"
	doctorRepositories "capstone/repositories/mysql/doctor"
	userRepositories "capstone/repositories/mysql/user"
	"capstone/routes"
	doctorUseCase "capstone/usecases/doctor"
	userUseCase "capstone/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())

	userRepo := userRepositories.NewUserRepo(db)
	doctorRepo := doctorRepositories.NewDoctorRepo(db)

	userUC := userUseCase.NewUserUseCase(userRepo)
	doctorUC := doctorUseCase.NewDoctorUseCase(doctorRepo)

	userCont := userController.NewUserController(userUC)
	doctorCont := doctorController.NewDoctorController(doctorUC)

	route := routes.NewRoute(userCont, doctorCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
