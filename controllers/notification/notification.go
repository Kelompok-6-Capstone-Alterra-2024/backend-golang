package notification

import (
	"capstone/controllers/notification/response"
	notificationEntities "capstone/entities/notification"
	"capstone/utilities"
	"capstone/utilities/base"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type NotificationController struct {
	notificationUseCase notificationEntities.NotificationUseCase
}

func NewNotificationController(notificationUseCase notificationEntities.NotificationUseCase) *NotificationController {
	return &NotificationController{notificationUseCase}
}

func (controller *NotificationController) GetAllDoctorNotification(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	doctorID, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse("unauthorized"))
	}

	notifications, err := controller.notificationUseCase.GetNotificationByDoctorID(doctorID, metadata)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}
	var notificationResponses []response.NotificationDoctorResponse
	for _, notification := range *notifications {
		notificationResponses = append(notificationResponses, *notification.ToDoctorResponse())
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("success get notifications", metadata, notificationResponses))
}

func (controller *NotificationController) UpdateToReadConsultation(c echo.Context) error {
	strNotificationID := c.Param("notificationID")
	notificationID, err := strconv.Atoi(strNotificationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("notification id must be a number"))
	}
	err = controller.notificationUseCase.UpdateStatusDoctorNotification(notificationID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("success update to read consultation", nil))
}

func (controller *NotificationController) DeleteToReadConsultation(c echo.Context) error {
	strNotificationID := c.Param("notificationID")
	notificationID, err := strconv.Atoi(strNotificationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("notification id must be a number"))
	}
	err = controller.notificationUseCase.DeleteDoctorNotification(notificationID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("success delete notification", nil))
}
