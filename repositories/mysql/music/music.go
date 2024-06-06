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

func (m *MusicRepo) GetAllMusicsByDoctorId(metadata entities.MetadataFull, userId int) ([]musicEntities.Music, error) {
	var musics []Music
	
	query := m.db.Where("doctor_id = ?", userId).Limit(metadata.Limit).Offset((metadata.Page-1)*metadata.Limit).Order(metadata.Sort + " " + metadata.Order)

	if metadata.Search != "" {
		query = query.Where("title LIKE ?", "%"+metadata.Search+"%")
	}

	err := query.Find(&musics).Error
	if err != nil {
		return []musicEntities.Music{}, constants.ErrServer
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

	err = m.db.Model(&music).Where("id = ?", musicId).Update("view_count", music.ViewCount+1).Error
	if err != nil {
		return musicEntities.Music{}, constants.ErrServer
	}

	return musicEnt, nil
}

func (m *MusicRepo) GetLikedMusics(metadata entities.Metadata, userId int) ([]musicEntities.Music, error) {
	var musicLikesDB []MusicLikes
	err := m.db.Limit(metadata.Limit).Offset((metadata.Page-1)*metadata.Limit).Where("user_id = ?", userId).Find(&musicLikesDB).Error
	if err != nil {
		return []musicEntities.Music{}, constants.ErrDataNotFound
	}

	var musicIds []int
	for _, musicLike := range musicLikesDB {
		musicIds = append(musicIds, int(musicLike.MusicId))
	}

	var musics []Music
	err = m.db.Where("id IN ?", musicIds).Find(&musics).Error
	if err != nil {
		return []musicEntities.Music{}, constants.ErrDataNotFound
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
			IsLiked:   true,
		}
	}

	return musicsEnt, nil
}

func (m *MusicRepo) LikeMusic(musicId int, userId int) error {
	var musicLikesDB MusicLikes

	err := m.db.Where("music_id = ? AND user_id = ?", musicId, userId).First(&musicLikesDB).Error
	if err == nil {
		return constants.ErrAlreadyLiked
	}

	musicLikesDB.MusicId = uint(musicId)
	musicLikesDB.UserId = uint(userId)

	err = m.db.Create(&musicLikesDB).Error
	if err != nil {
		return constants.ErrServer
	}

	return nil
}

func (m *MusicRepo) CountMusicByDoctorId(doctorId int) (int, error) {
	var counter int64
	err := m.db.Model(&Music{}).Where("doctor_id = ?", doctorId).Count(&counter).Error
	if err != nil {
		return 0, constants.ErrServer
	}

	return int(counter), nil
}

func (m *MusicRepo) CountMusicLikesByDoctorId(doctorId int) (int, error) {
	var counter int64
	err := m.db.Table("music_likes").
		Joins("JOIN musics ON music_likes.music_id = musics.id").
		Where("musics.doctor_id = ?", doctorId).
		Count(&counter).Error
	if err != nil {
		return 0, constants.ErrServer
	}

	return int(counter), nil
}

func (m *MusicRepo) CountMusicViewCountByDoctorId(doctorId int) (int, error) {
	var totalViews int64
	err := m.db.Model(&Music{}).
		Where("doctor_id = ?", doctorId).
		Select("SUM(view_count)").
		Scan(&totalViews).Error
	if err != nil {
		return 0, constants.ErrServer
	}

	return int(totalViews), nil
}

func (m *MusicRepo) PostMusic(music musicEntities.Music) (musicEntities.Music, error) {
	var musicDB Music

	musicDB.DoctorId = music.DoctorId
	musicDB.Title = music.Title
	musicDB.Singer = music.Singer
	musicDB.MusicUrl = music.MusicUrl
	musicDB.ImageUrl = music.ImageUrl

	err := m.db.Create(&musicDB).Error
	if err != nil {
		return musicEntities.Music{}, constants.ErrServer
	}

	return musicEntities.Music{
		Id:        musicDB.ID,
		Title:     musicDB.Title,
		Singer:    musicDB.Singer,
		MusicUrl:  musicDB.MusicUrl,
		ImageUrl:  musicDB.ImageUrl,
		ViewCount: musicDB.ViewCount,
	}, nil
}

func (m *MusicRepo) GetMusicByIdForDoctor(musicId int) (musicEntities.Music, error) {
	var musicDB Music
	err := m.db.Where("id = ?", musicId).First(&musicDB).Error
	if err != nil {
		return musicEntities.Music{}, constants.ErrDataNotFound
	}
	return musicEntities.Music{
		Id:        musicDB.ID,
		Title:     musicDB.Title,
		Singer:    musicDB.Singer,
		MusicUrl:  musicDB.MusicUrl,
		ImageUrl:  musicDB.ImageUrl,
		ViewCount: musicDB.ViewCount,
	}, nil
}