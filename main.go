package main

import (
	"capstone/configs"
	userController "capstone/controllers/user"
	"capstone/repositories/mysql"
	userRepositories "capstone/repositories/mysql/user"
	"capstone/routes"
	userUseCase "capstone/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())

	userRepo := userRepositories.NewUserRepo(db)
	userUC := userUseCase.NewUserUseCase(userRepo)
	userCont := userController.NewUserController(userUC)

	route := routes.NewRoute(userCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}