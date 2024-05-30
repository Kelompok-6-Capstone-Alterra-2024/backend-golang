package music

import (
	"capstone/entities"
	musicEntities "capstone/entities/music"
)

type MusicUseCase struct {
	musicInterface musicEntities.RepositoryInterface
}

func NewMusicUseCase(musicInterface musicEntities.RepositoryInterface) *MusicUseCase {
	return &MusicUseCase{
		musicInterface: musicInterface,
	}
}

func (musicUseCase *MusicUseCase) GetAllMusics(metadata entities.Metadata, userId int) ([]musicEntities.Music, error) {
	musics, err := musicUseCase.musicInterface.GetAllMusics(metadata, userId)
	if err != nil {
		return []musicEntities.Music{}, err
	}
	return musics, nil
}

func (musicUseCase *MusicUseCase) GetMusicById(musicId int, userId int) (musicEntities.Music, error) {
	music, err := musicUseCase.musicInterface.GetMusicById(musicId, userId)
	if err != nil {
		return musicEntities.Music{}, err
	}
	return music, nil
}

func (musicUseCase *MusicUseCase) GetLikedMusics(metadata entities.Metadata, userId int) ([]musicEntities.Music, error) {
	musics, err := musicUseCase.musicInterface.GetLikedMusics(metadata, userId)
	if err != nil {
		return []musicEntities.Music{}, err
	}
	return musics, nil
}