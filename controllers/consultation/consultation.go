package consultation

import (
	"capstone/controllers/consultation/request"
	"capstone/controllers/consultation/response"
	"capstone/entities/consultation"
	"capstone/utilities"
	"capstone/utilities/base"
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
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	consultationRequest.UserID = userId
	response, err := controller.consultationUseCase.CreateConsultation(consultationRequest.ToEntities(date))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Add Consultation", response))
}

func (controller *ConsultationController) GetConsultationByID(c echo.Context) error {
	consultationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID"))
	}
	response, err := controller.consultationUseCase.GetConsultationByID(consultationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", response.ToResponse()))
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
	consultations, err := controller.consultationUseCase.GetAllConsultation(metadata, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	var responses []response.ConsultationResponse
	for _, value := range *consultations {
		responses = append(responses, *value.ToResponse())
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", responses))
}
