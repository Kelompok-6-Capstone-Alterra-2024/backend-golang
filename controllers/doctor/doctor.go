package doctor

import (
	"capstone/controllers/doctor/request"
	"capstone/controllers/doctor/response"
	doctorUseCase "capstone/entities/doctor"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DoctorController struct {
	doctorUseCase doctorUseCase.DoctorUseCaseInterface
}

func NewDoctorController(doctorUseCase doctorUseCase.DoctorUseCaseInterface) *DoctorController {
	return &DoctorController{
		doctorUseCase: doctorUseCase,
	}
}

func (controller *DoctorController) Register(c echo.Context) error {
	var doctorFromRequest request.DoctorRegisterRequest
	c.Bind(&doctorFromRequest)

	imageFromRequest, err := c.FormFile("profile_picture")
	doctorFromRequest.ProfilePicture = imageFromRequest

	doctorRequest, err := doctorFromRequest.ToDoctorEntities()
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	doctorResult, err := controller.doctorUseCase.Register(doctorRequest)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	doctorResponse := doctorResult.ToResponse()
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register", doctorResponse))
}

func (controller *DoctorController) Login(c echo.Context) error {
	var doctorFromRequest request.DoctorLoginRequest
	c.Bind(&doctorFromRequest)

	doctorRequest := doctorFromRequest.ToDoctorLoginEntities()
	doctorResult, err := controller.doctorUseCase.Login(doctorRequest)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	doctorResponse := doctorResult.ToResponse()
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", doctorResponse))
}

func (controller *DoctorController) GetByID(c echo.Context) error {
	strDoctorID := c.Param("id")
	doctorID, err := strconv.Atoi(strDoctorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	doctorResult, err := controller.doctorUseCase.GetDoctorByID(doctorID)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	doctorResponse := doctorResult.ToDoctorResponse()
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Doctor By ID", doctorResponse))
}

func (controller *DoctorController) GetAll(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	doctorResult, err := controller.doctorUseCase.GetAllDoctor(metadata)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var doctorResponse []response.DoctorResponse
	for _, doctor := range *doctorResult {
		doctorResponse = append(doctorResponse, *doctor.ToDoctorResponse())
	}
	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Doctor", metadata, doctorResponse))
}

func (controller *DoctorController) GetActive(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	doctorResult, err := controller.doctorUseCase.GetActiveDoctor(metadata)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var doctorResponse []response.DoctorResponse
	for _, doctor := range *doctorResult {
		doctorResponse = append(doctorResponse, *doctor.ToDoctorResponse())
	}
	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get Active Doctor", metadata, doctorResponse))
}

func (c *DoctorController) GoogleLogin(ctx echo.Context) error {
	url := c.doctorUseCase.HandleGoogleLogin()
	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *DoctorController) GoogleCallback(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	result, err := c.doctorUseCase.HandleGoogleCallback(ctx.Request().Context(), code)
	if err != nil {
		return ctx.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var res response.DoctorLoginAndRegisterResponse
	res.ID = result.ID
	res.Token = result.Token

	return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Success Login Oauth", res))
}

func (c *DoctorController) FacebookLogin(ctx echo.Context) error {
	url := c.doctorUseCase.HandleFacebookLogin()
	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *DoctorController) FacebookCallback(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	result, err := c.doctorUseCase.HandleFacebookCallback(ctx.Request().Context(), code)
	if err != nil {
		return ctx.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var res response.DoctorLoginAndRegisterResponse
	res.ID = result.ID
	res.Token = result.Token

	return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Success Login Oauth", res))
}

func (c *DoctorController) UpdateDoctorProfile(ctx echo.Context) error {
	var doctorFromRequest request.UpdateDoctorProfileRequest
	if err := ctx.Bind(&doctorFromRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	token := ctx.Request().Header.Get("Authorization")
	doctorID, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return ctx.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	imageURL, err := utilities.UploadImage(file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	doctorEntities := doctorUseCase.Doctor{
		ID:               uint(doctorID),
		Username:         doctorFromRequest.Username,
		Name:             doctorFromRequest.Name,
		Address:          doctorFromRequest.Address,
		PhoneNumber:      doctorFromRequest.PhoneNumber,
		Gender:           doctorFromRequest.Gender,
		ProfilePicture:   imageURL,
		Experience:       doctorFromRequest.Experience,
		Almamater:        doctorFromRequest.Almamater,
		GraduationYear:   doctorFromRequest.GraduationYear,
		PracticeLocation: doctorFromRequest.PracticeLocation,
		PracticeCity:     doctorFromRequest.PracticeCity,
		PracticeProvince: doctorFromRequest.PracticeProvince,
		StrNumber:        doctorFromRequest.StrNumber,
		Fee:              doctorFromRequest.Fee,
		Specialist:       doctorFromRequest.Specialist,
	}

	updatedDoctor, err := c.doctorUseCase.UpdateDoctorProfile(&doctorEntities)
	if err != nil {
		return ctx.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	doctorResponse := response.DoctorUpdateProfileResponse{
		ID:               updatedDoctor.ID,
		Username:         updatedDoctor.Username,
		Email:            updatedDoctor.Email,
		Name:             updatedDoctor.Name,
		Address:          updatedDoctor.Address,
		PhoneNumber:      updatedDoctor.PhoneNumber,
		Gender:           updatedDoctor.Gender,
		ProfilePicture:   updatedDoctor.ProfilePicture,
		Experience:       updatedDoctor.Experience,
		Almamater:        updatedDoctor.Almamater,
		GraduationYear:   updatedDoctor.GraduationYear,
		PracticeLocation: updatedDoctor.PracticeLocation,
		PracticeCity:     updatedDoctor.PracticeCity,
		PracticeProvince: updatedDoctor.PracticeProvince,
		StrNumber:        updatedDoctor.StrNumber,
		Fee:              updatedDoctor.Fee,
		Specialist:       updatedDoctor.Specialist,
	}

	return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Profile", doctorResponse))
}
