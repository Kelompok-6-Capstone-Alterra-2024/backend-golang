package main

import (
	"capstone/configs"
	complaintController "capstone/controllers/complaint"
	consultationController "capstone/controllers/consultation"
	doctorController "capstone/controllers/doctor"
	userController "capstone/controllers/user"
	"capstone/repositories/mysql"
	complaintRepositories "capstone/repositories/mysql/complaint"
	consultationRepositories "capstone/repositories/mysql/consultation"
	doctorRepositories "capstone/repositories/mysql/doctor"
	userRepositories "capstone/repositories/mysql/user"
	"capstone/routes"
	complaintUseCase "capstone/usecases/complaint"
	consultationUseCase "capstone/usecases/consultation"
	doctorUseCase "capstone/usecases/doctor"
	userUseCase "capstone/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())

	userRepo := userRepositories.NewUserRepo(db)
	doctorRepo := doctorRepositories.NewDoctorRepo(db)
	consultationRepo := consultationRepositories.NewConsultationRepo(db)
	complaintRepo := complaintRepositories.NewComplaintRepo(db)

	userUC := userUseCase.NewUserUseCase(userRepo)
	doctorUC := doctorUseCase.NewDoctorUseCase(doctorRepo)
	consultationUC := consultationUseCase.NewConsultationUseCase(consultationRepo)
	complaintUC := complaintUseCase.NewComplaintUseCase(complaintRepo)

	userCont := userController.NewUserController(userUC)
	doctorCont := doctorController.NewDoctorController(doctorUC)
	consultationCont := consultationController.NewConsultationController(consultationUC)
	complaintCont := complaintController.NewComplaintController(complaintUC)

	route := routes.NewRoute(userCont, doctorCont, consultationCont, complaintCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
