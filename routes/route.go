package routes

import (
	"capstone/controllers/article"
	"capstone/controllers/doctor"
	"capstone/controllers/user"
	myMiddleware "capstone/middlewares"

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
	userAuth.POST("register", r.userController.Register) //Register User
	userAuth.POST("login", r.userController.Login)       //Login User

	//user := userAuth.Group("", echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	doctorAuth := e.Group("/v1/doctor")
	doctorAuth.POST("/register", r.doctorController.Register) //Register Doctor
	doctorAuth.POST("/login", r.doctorController.Login)       //Login Doctor

	articleAuth := e.Group("/v1/article")
	articleAuth.POST("/create", r.articleController.CreateArticle)
	articleAuth.GET("/list", r.articleController.GetAllArticle)
}
