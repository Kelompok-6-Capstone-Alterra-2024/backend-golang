package forum

import (
	"capstone/controllers/forum/request"
	"capstone/controllers/forum/response"
	forumEntities "capstone/entities/forum"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"
	"strconv"

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

func (forumController *ForumController) GetJoinedForum(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	forums, err := forumController.forumUseCase.GetJoinedForum(uint(userId), *metadata)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var resp []response.ForumJoinedResponse
	for _, forum := range forums {
		resp = append(resp, response.ForumJoinedResponse{
			ForumID:         forum.ID,
			Name:            forum.Name,
			ImageUrl:        forum.ImageUrl,
			NumberOfMembers: forum.NumberOfMembers,
		})

		for _, user := range forum.User {
			resp[len(resp)-1].User = append(resp[len(resp)-1].User, response.UserJoined{
				UserID:   uint(user.Id),
				ProfilePicture: user.ProfilePicture,
			})
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get Joined Forum", metadata, resp))
}

func (forumController *ForumController) GetRecommendationForum(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	forums, err := forumController.forumUseCase.GetRecommendationForum(uint(userId), *metadata)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var resp []response.ForumRecommendationResponse
	for _, forum := range forums {
		resp = append(resp, response.ForumRecommendationResponse{
			ForumID:         forum.ID,
			Name:            forum.Name,
			ImageUrl:        forum.ImageUrl,
			NumberOfMembers: forum.NumberOfMembers,
		})
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get Recommendation Forum", metadata, resp))
}

func (forumController *ForumController) GetForumById(c echo.Context) error {
	forumId := c.Param("id")
	forumIdInt, _ := strconv.Atoi(forumId)

	forum, err := forumController.forumUseCase.GetForumById(uint(forumIdInt))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var resp []response.ForumDetailResponse

	respForum := response.ForumDetailResponse{
		ForumID:     forum.ID,
		Name:        forum.Name,
		Description: forum.Description,
		ImageUrl:    forum.ImageUrl,
		Post:        []response.PostResponse{}, // Initialize Post slice
	}

	for _, post := range forum.Post {
		respPost := response.PostResponse{
			PostID:   post.ID,
			Content:  post.Content,
			ImageUrl: post.ImageUrl,
			User: response.UserPostResponse{
				UserID:   uint(post.User.Id),
				Username: post.User.Username,
				ImageUrl: post.User.ProfilePicture,
			},
		}
		respForum.Post = append(respForum.Post, respPost)
	}

	resp = append(resp, respForum)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Forum By Id", resp))
}