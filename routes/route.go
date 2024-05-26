package routes

import (
	"capstone/controllers/user"
	myMiddleware "capstone/middlewares"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	userController *user.UserController
}

func NewRoute(userController *user.UserController) *RouteController {
	return &RouteController{
		userController:     userController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)

	e.POST("/v1/user/register", r.userController.Register) //Register User
	e.POST("/v1/user/login", r.userController.Login) //Login User
}