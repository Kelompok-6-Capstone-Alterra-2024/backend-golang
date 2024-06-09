package consultation

import (
	"capstone/controllers/consultation/request"
	"capstone/controllers/consultation/response"
	"capstone/entities/consultation"
	"capstone/utilities"
	"capstone/utilities/base"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ConsultationController struct {
	consultationUseCase consultation.ConsultationUseCase
}

func NewConsultationController(consultationUseCase consultation.ConsultationUseCase) *ConsultationController {
	return &ConsultationController{
		consultationUseCase,
	}
}

func (controller *ConsultationController) CreateConsultation(c echo.Context) error {
	var consultationRequest request.ConsultationRequest
	c.Bind(&consultationRequest)
	token := c.Request().Header.Get("Authorization")
	userId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	date, err := utilities.StringToDate(consultationRequest.Date)
	time, err := utilities.StringToTime(consultationRequest.Time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	consultationRequest.UserID = userId
	consultationResponse, err := controller.consultationUseCase.CreateConsultation(consultationRequest.ToEntities(date, time))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Add Consultation", consultationResponse.ToUserResponse()))
}

func (controller *ConsultationController) GetConsultationByID(c echo.Context) error {
	consultationID, err := strconv.Atoi(c.Param("id"))
	token := c.Request().Header.Get("Authorization")
	userId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID"))
	}
	consultationResponse, err := controller.consultationUseCase.GetConsultationByID(consultationID)
	if userId != consultationResponse.UserID {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse("Unauthorized"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", consultationResponse.ToUserResponse()))
}

func (controller *ConsultationController) GetAllConsultation(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	userId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	consultations, err := controller.consultationUseCase.GetAllUserConsultation(metadata, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	var responses []response.ConsultationUserResponse
	for _, value := range *consultations {
		responses = append(responses, *value.ToUserResponse())
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", responses))
}

func (controller *ConsultationController) UpdateStatusConsultation(c echo.Context) error {
	var consultationRequest request.ConsultationStatusUpdateRequest

	consultationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID"))
	}
	c.Bind(&consultationRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	consultationRequest.ID = uint(consultationID)
	consultationResponse, err := controller.consultationUseCase.UpdateStatusConsultation(consultationRequest.ToEntities())
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Consultation", consultationResponse.ToUserResponse()))
}

func (controller *ConsultationController) GetAllDoctorConsultation(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	doctorId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	consultations, err := controller.consultationUseCase.GetAllDoctorConsultation(metadata, doctorId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	var responses []response.ConsultationDoctorResponse
	for _, value := range *consultations {
		responses = append(responses, *value.ToDoctorResponse())
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", responses))
}

func (controller *ConsultationController) CountConsultationByDoctorID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	doctorId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	var count int64
	status := c.QueryParam("status")
	fmt.Print(status)
	if status != "" {
		count, err = controller.consultationUseCase.CountConsultationByStatus(doctorId, status)
		if err != nil {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", count))
	}

	count, err = controller.consultationUseCase.CountConsultationByDoctorID(doctorId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", count))
}

func (controller *ConsultationController) CountConsultationToday(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	doctorId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	count, err := controller.consultationUseCase.CountConsultationToday(doctorId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", count))
}
