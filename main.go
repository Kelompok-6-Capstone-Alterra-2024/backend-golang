package main

import (
	"capstone/configs"
	articleController "capstone/controllers/article"
	doctorController "capstone/controllers/doctor"
	userController "capstone/controllers/user"
	"capstone/repositories/mysql"
	articleRepositories "capstone/repositories/mysql/article"
	doctorRepositories "capstone/repositories/mysql/doctor"
	userRepositories "capstone/repositories/mysql/user"
	"capstone/routes"
	articleUseCase "capstone/usecases/article"
	doctorUseCase "capstone/usecases/doctor"
	userUseCase "capstone/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())

	userRepo := userRepositories.NewUserRepo(db)
	doctorRepo := doctorRepositories.NewDoctorRepo(db)
	articleRepo := articleRepositories.NewArticleRepo(db)

	userUC := userUseCase.NewUserUseCase(userRepo)
	doctorUC := doctorUseCase.NewDoctorUseCase(doctorRepo)
	articleUC := articleUseCase.NewArticleUseCase(articleRepo)

	userCont := userController.NewUserController(userUC)
	doctorCont := doctorController.NewDoctorController(doctorUC)
	articleCont := articleController.NewArticleController(articleUC)

	route := routes.NewRoute(userCont, doctorCont, articleCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
