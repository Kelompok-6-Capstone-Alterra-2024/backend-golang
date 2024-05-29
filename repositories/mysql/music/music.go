package music

import (
	"capstone/constants"
	"capstone/entities"
	musicEntities "capstone/entities/music"

	"gorm.io/gorm"
)

type MusicRepo struct {
	db *gorm.DB
}

func NewMusicRepo(db *gorm.DB) *MusicRepo {
	return &MusicRepo{
		db: db,
	}
}

func (m *MusicRepo) GetAllMusics(metadata entities.Metadata, userId int) ([]musicEntities.Music, error) {
	var musics []Music
	
	err := m.db.Limit(metadata.Limit).Offset((metadata.Page-1)*metadata.Limit).Find(&musics).Error
	if err != nil {
		return []musicEntities.Music{}, constants.ErrDataNotFound
	}

	musicLikes := make([]MusicLikes, len(musics))
	var counter int64
	var isLiked []bool

	for i := 0; i < len(musics); i++ {
		musicLikes[i].UserId = uint(userId)
		musicLikes[i].MusicId = musics[i].ID
		err = m.db.Model(&musicLikes[i]).Where("user_id = ? AND music_id = ?", musicLikes[i].UserId, musicLikes[i].MusicId).Count(&counter).Error
		if err != nil {
			return []musicEntities.Music{}, constants.ErrServer
		}

		if counter > 0 {
			isLiked = append(isLiked, true)
		} else {
			isLiked = append(isLiked, false)
		}

		counter = 0
	}

	musicsEnt := make([]musicEntities.Music, len(musics))
	for i := 0; i < len(musics); i++ {
		musicsEnt[i] = musicEntities.Music{
			Id:        musics[i].ID,
			Title:     musics[i].Title,
			Singer:    musics[i].Singer,
			MusicUrl:  musics[i].MusicUrl,
			ImageUrl:  musics[i].ImageUrl,
			ViewCount: musics[i].ViewCount,
			IsLiked:   isLiked[i],
		}
	}

	return musicsEnt, nil
}

func (m *MusicRepo) GetMusicById(musicId int, userId int) (musicEntities.Music, error) {
	var music Music
	err := m.db.Where("id = ?", musicId).First(&music).Error
	if err != nil {
		return musicEntities.Music{}, constants.ErrDataNotFound
	}

	var musicLikes MusicLikes
	var isLiked bool
	var counter int64

	err = m.db.Model(&musicLikes).Where("user_id = ? AND music_id = ?", userId, musicId).Count(&counter).Error
	if err != nil {
		return musicEntities.Music{}, constants.ErrServer
	}

	if counter > 0 {
		isLiked = true
	} else {
		isLiked = false
	}

	musicEnt := musicEntities.Music{
		Id:        music.ID,
		Title:     music.Title,
		Singer:    music.Singer,
		MusicUrl:  music.MusicUrl,
		ImageUrl:  music.ImageUrl,
		ViewCount: music.ViewCount,
		IsLiked:   isLiked,
	}

	return musicEnt, nil
}