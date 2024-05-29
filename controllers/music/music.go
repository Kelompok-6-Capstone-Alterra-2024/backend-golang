package music

import (
	"capstone/controllers/music/response"
	musicEntities "capstone/entities/music"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"

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