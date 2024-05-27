package consultation

import (
	"capstone/controllers/consultation/request"
	"capstone/entities/consultation"
	"capstone/utilities"
	"capstone/utilities/base"
	"github.com/labstack/echo/v4"
	"net/http"
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
