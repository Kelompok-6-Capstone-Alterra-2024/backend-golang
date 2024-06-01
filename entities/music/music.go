package music

import "capstone/entities"

type Music struct {
	Id        uint
	Title     string
	Singer    string
	MusicUrl  string
	ImageUrl  string
	ViewCount int
	IsLiked   bool
}

type RepositoryInterface interface {
	GetAllMusics(metadata entities.Metadata, userId int) ([]Music, error)
	GetMusicById(musicId int, userId int) (Music, error)
	GetLikedMusics(metadata entities.Metadata, userId int) ([]Music, error)
	LikeMusic(musicId int, userId int) error
}

type UseCaseInterface interface {
	GetAllMusics(metadata entities.Metadata, userId int) ([]Music, error)
	GetMusicById(musicId int, userId int) (Music, error)
	GetLikedMusics(metadata entities.Metadata, userId int) ([]Music, error)
	LikeMusic(musicId int, userId int) error
}