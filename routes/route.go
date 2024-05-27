package routes

import (
	"capstone/controllers/article"
	"capstone/controllers/doctor"
	"capstone/controllers/user"
	myMiddleware "capstone/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	userController    *user.UserController
	doctorController  *doctor.DoctorController
	articleController *article.ArticleController
}

func NewRoute(userController *user.UserController, doctorController *doctor.DoctorController, articleController *article.ArticleController) *RouteController {
	return &RouteController{
		userController:    userController,
		doctorController:  doctorController,
		articleController: articleController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)

	userAuth := e.Group("/v1/user")
	userAuth.POST("/register", r.userController.Register) //Register User
	userAuth.POST("/login", r.userController.Login)       //Login User

	userRoute := userAuth.Group("/")
	userRoute.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	// Doctor
	userRoute.GET("doctor/:id", r.doctorController.GetByID)         //Get Doctor By ID
	userRoute.GET("doctor", r.doctorController.GetAll)              //Get All Doctor
	userRoute.GET("doctor/available", r.doctorController.GetActive) //Get All Active Doctor

	doctorAuth := e.Group("/v1/doctor")
	doctorAuth.POST("/register", r.doctorController.Register) //Register Doctor
	doctorAuth.POST("/login", r.doctorController.Login)       //Login Doctor

	articleAuth := e.Group("/v1/article")
	articleAuth.POST("/create", r.articleController.CreateArticle)
	articleAuth.GET("/list", r.articleController.GetAllArticle)
}
