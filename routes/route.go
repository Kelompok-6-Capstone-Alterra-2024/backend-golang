package routes

import (
	"capstone/controllers/doctor"
	"capstone/controllers/user"
	myMiddleware "capstone/middlewares"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	userController   *user.UserController
	doctorController *doctor.DoctorController
}

func NewRoute(userController *user.UserController, doctorController *doctor.DoctorController) *RouteController {
	return &RouteController{
		userController:   userController,
		doctorController: doctorController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)

	userAuth := e.Group("/v1/user")
	userAuth.POST("register", r.userController.Register) //Register User
	userAuth.POST("login", r.userController.Login)       //Login User

	//user := userAuth.Group("", echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	doctorAuth := e.Group("/v1/doctor")
	doctorAuth.POST("/register", r.doctorController.Register) //Register Doctor
	doctorAuth.POST("/login", r.doctorController.Login)       //Login Doctor

}
