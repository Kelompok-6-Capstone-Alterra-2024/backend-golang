package user

import (
	"capstone/controllers/user/request"
	"capstone/controllers/user/response"
	userEntities "capstone/entities/user"
	"capstone/utilities"
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

func (c *UserController) GetPointsByUserId(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	points, err := c.userUseCase.GetPointsByUserId(userId)
	if err != nil {
		return ctx.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var res response.UserPointsResponse
	res.Points = points
	return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Points", res))
}

func (c *UserController) ResetPassword(ctx echo.Context) error {
	var req request.ResetPasswordRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	err := c.userUseCase.ResetPassword(req.Email, req.Password)
	if err != nil {
		return ctx.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Success Reset Password", nil))
}

func (c *UserController) FacebookLogin(ctx echo.Context) error {
	url := c.userUseCase.HandleFacebookLogin()
	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *UserController) FacebookCallback(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	result, err := c.userUseCase.HandleFacebookCallback(ctx.Request().Context(), code)
	if err != nil {
		return ctx.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var res response.UserLoginRegisterResponse
	res.Id = result.Id
	res.Token = result.Token

	return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Success Login Oauth", res))
}

func (userController *UserController) UpdateProfile(c echo.Context) error {
	var userFromRequest request.UpdateProfileRequest
	if err := c.Bind(&userFromRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	token := c.Request().Header.Get("Authorization")
	userId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse("Invalid token"))
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid image format"))
	}

	imageURL, err := utilities.UploadImage(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse("Failed to upload image"))
	}

	userEntities := userEntities.User{
		Id:             userId,
		Name:           userFromRequest.Name,
		Username:       userFromRequest.Username,
		Address:        userFromRequest.Address,
		Bio:            userFromRequest.Bio,
		PhoneNumber:    userFromRequest.PhoneNumber,
		Gender:         userFromRequest.Gender,
		Age:            userFromRequest.Age,
		ProfilePicture: imageURL,
	}

	updatedUser, err := userController.userUseCase.UpdateUserProfile(&userEntities)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserUpdateProfileResponse{
		Id:             updatedUser.Id,
		Name:           updatedUser.Name,
		Username:       updatedUser.Username,
		Address:        updatedUser.Address,
		Bio:            updatedUser.Bio,
		PhoneNumber:    updatedUser.PhoneNumber,
		Gender:         updatedUser.Gender,
		Age:            updatedUser.Age,
		ProfilePicture: updatedUser.ProfilePicture,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Profile", userResponse))
}
