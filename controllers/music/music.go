package music

import (
	"capstone/controllers/music/response"
	musicEntities "capstone/entities/music"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MusicController struct {
	musicUseCase musicEntities.UseCaseInterface
}

func NewMusicController(musicUseCase musicEntities.UseCaseInterface) *MusicController {
	return &MusicController{
		musicUseCase: musicUseCase,
	}
}

func (musicController *MusicController) GetAllMusics(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	musics, err := musicController.musicUseCase.GetAllMusics(*metadata, userId)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	musicResp := make([]response.MusicResponse, len(musics))

	for i, music := range musics {
		musicResp[i] = response.MusicResponse{
			Id:        music.Id,
			Title:     music.Title,
			Singer:    music.Singer,
			MusicUrl:  music.MusicUrl,
			ImageUrl:  music.ImageUrl,
			ViewCount: music.ViewCount,
			IsLiked:   music.IsLiked,
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Musics", metadata, musicResp))
}

func (musicController *MusicController) GetMusicByID(c echo.Context) error {
	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	music, err := musicController.musicUseCase.GetMusicById(id, userId)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	musicResp := response.MusicResponse{
		Id:        music.Id,
		Title:     music.Title,
		Singer:    music.Singer,
		MusicUrl:  music.MusicUrl,
		ImageUrl:  music.ImageUrl,
		ViewCount: music.ViewCount,
		IsLiked:   music.IsLiked,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Music By Id", musicResp))
}

func (musicController *MusicController) GetLikedMusics(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	musics, err := musicController.musicUseCase.GetLikedMusics(*metadata, userId)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	musicResp := make([]response.MusicResponse, len(musics))

	for i, music := range musics {
		musicResp[i] = response.MusicResponse{
			Id:        music.Id,
			Title:     music.Title,
			Singer:    music.Singer,
			MusicUrl:  music.MusicUrl,
			ImageUrl:  music.ImageUrl,
			ViewCount: music.ViewCount,
			IsLiked:   music.IsLiked,
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get Liked Musics", metadata, musicResp))
}