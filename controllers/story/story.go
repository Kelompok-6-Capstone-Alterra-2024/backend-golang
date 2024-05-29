package story

import (
	"capstone/controllers/story/response"
	storyEntities "capstone/entities/story"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StoryController struct {
	storyUseCase storyEntities.UseCaseInterface
}

func NewStoryController(storyUseCase storyEntities.UseCaseInterface) *StoryController {
	return &StoryController{
		storyUseCase: storyUseCase,
	}
}

func (storyController *StoryController) GetAllStories(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	stories, err := storyController.storyUseCase.GetAllStories(*metadata, userId)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	storiesResp := make([]response.StoriesGetAllResponse, len(stories)) 

	for i, story := range stories {
		storiesResp[i] = response.StoriesGetAllResponse{
			ID:       story.Id,
			Title:    story.Title,
			Content:  story.Content,
			Date:     story.Date,
			ImageUrl: story.ImageUrl,
			ViewCount: story.ViewCount,
			IsLiked:  story.IsLiked,
			Doctor: response.DoctorGetAllResponse{
				ID:   story.Doctor.ID,
				Name: story.Doctor.Name,
			},
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Stories", metadata, storiesResp))
}

func (storyController *StoryController) GetStoryById(c echo.Context) error {
	strId := c.Param("id")
	storyId, _ := strconv.Atoi(strId)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	story, err := storyController.storyUseCase.GetStoryById(storyId, userId)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	storyResp := response.StoriesGetAllResponse{
		ID:       story.Id,
		Title:    story.Title,
		Content:  story.Content,
		Date:     story.Date,
		ImageUrl: story.ImageUrl,
		ViewCount: story.ViewCount,
		IsLiked:  story.IsLiked,
		Doctor: response.DoctorGetAllResponse{
			ID:   story.Doctor.ID,
			Name: story.Doctor.Name,
		},
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Story By Id", storyResp))
}