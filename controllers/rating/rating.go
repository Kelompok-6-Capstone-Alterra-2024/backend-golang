package rating

import (
	"capstone/controllers/rating/request"
	"capstone/controllers/rating/response"
	ratingEntities "capstone/entities/rating"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RatingController struct {
	ratingUseCase ratingEntities.UseCaseInterface
}

func NewRatingController(ratingUseCase ratingEntities.UseCaseInterface) *RatingController {
	return &RatingController{
		ratingUseCase: ratingUseCase,
	}
}

func (ratingController *RatingController) SendFeedback(c echo.Context) error {
	var feedBackReq request.SendFeedbackRequest
	c.Bind(&feedBackReq)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	ratingEnt := ratingEntities.Rating{
		DoctorId: feedBackReq.DoctorId,
		UserId:   uint(userId),
		Rate:     feedBackReq.Rate,
		Message:  feedBackReq.Message,
	}

	result, err := ratingController.ratingUseCase.SendFeedback(ratingEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	responseResult := response.RatingCreateResponse{
		Id:       result.Id,
		UserId:   result.UserId,
		DoctorId: result.DoctorId,
		Rate:     result.Rate,
		Message:  result.Message,
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Send Feedback", responseResult))
}
