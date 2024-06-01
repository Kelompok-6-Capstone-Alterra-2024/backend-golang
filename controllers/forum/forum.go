package forum

import (
	"capstone/controllers/forum/request"
	forumEntities "capstone/entities/forum"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ForumController struct {
	forumUseCase forumEntities.UseCaseInterface
}

func NewForumController(forumUseCase forumEntities.UseCaseInterface) *ForumController {
	return &ForumController{
		forumUseCase: forumUseCase,
	}
}

func (forumController *ForumController) JoinForum(c echo.Context) error {
	var req request.ForumJoinRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	err = forumController.forumUseCase.JoinForum(req.ForumID, uint(userId))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Join Forum", nil))
}