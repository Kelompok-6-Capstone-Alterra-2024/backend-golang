package user

import (
	"capstone/controllers/user/request"
	"capstone/controllers/user/response"
	userEntities "capstone/entities/user"
	"capstone/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase userEntities.UseCaseInterface
}

func NewUserController(userUseCase userEntities.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (userController *UserController) Register(c echo.Context) error {
	var userFromRequest request.UserRegisterRequest
	c.Bind(&userFromRequest)

	userEntities := userEntities.User{
		Username: userFromRequest.Username,
		Email:    userFromRequest.Email,
		Password: userFromRequest.Password,
	}

	newUser, err := userController.userUseCase.Register(&userEntities)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserLoginRegisterResponse{
		Id:    newUser.Id,
		Token: newUser.Token,
	}
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Register", userResponse))
}

func (userController *UserController) Login(c echo.Context) error {
	var userFromRequest request.UserLoginRequest
	c.Bind(&userFromRequest)

	userEntities := userEntities.User{
		Username: userFromRequest.Username,
		Password: userFromRequest.Password,
	}

	userFromDb, err := userController.userUseCase.Login(&userEntities)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserLoginRegisterResponse{
		Id:    userFromDb.Id,
		Token: userFromDb.Token,
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", userResponse))
}

func (c *UserController) GoogleLogin(ctx echo.Context) error {
    url := c.userUseCase.HandleGoogleLogin()
    return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *UserController) GoogleCallback(ctx echo.Context) error {
    code := ctx.QueryParam("code")
    result, err := c.userUseCase.HandleGoogleCallback(ctx.Request().Context(), code)
    if err != nil {
        return ctx.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
    }

	var res response.UserLoginRegisterResponse
	res.Id = result.Id
	res.Token = result.Token

    return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Success Login Oauth", res))
}