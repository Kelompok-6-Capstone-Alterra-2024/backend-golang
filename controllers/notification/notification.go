package notification

import (
	"capstone/controllers/notification/request"
	"capstone/controllers/notification/response"
	notificationEntities "capstone/entities/notification"
	"capstone/utilities"
	"capstone/utilities/base"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NotificationController struct {
	notificationUseCase notificationEntities.NotificationUseCase
}

func NewNotificationController(notificationUseCase notificationEntities.NotificationUseCase) *NotificationController {
	return &NotificationController{notificationUseCase}
}

func (controller *NotificationController) GetAllDoctorNotification(c echo.Context) error {
	var notificationRequest request.NotificationCreateRequest

	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	if err := c.Bind(&notificationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	notifications, err := controller.notificationUseCase.GetNotificationByDoctorID(notificationRequest.UserID, metadata)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}
	var notificationResponses []response.NotificationDoctorResponse
	for _, notification := range *notifications {
		notificationResponses = append(notificationResponses, *notification.ToDoctorResponse())
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("success get notifications", metadata, notificationResponses))
}
