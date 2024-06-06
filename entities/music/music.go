package music

import (
	"capstone/entities"
	"capstone/entities/doctor"
	"mime/multipart"
)

type Music struct {
	Id        uint
	Title     string
	Singer    string
	MusicUrl  string
	ImageUrl  string
	ViewCount int
	IsLiked   bool
	DoctorId  uint
	Doctor    doctor.Doctor
}

type RepositoryInterface interface {
	GetAllMusics(metadata entities.Metadata, userId int) ([]Music, error)
	GetAllMusicsByDoctorId(metadata entities.MetadataFull, userId int) ([]Music, error)
	GetMusicById(musicId int, userId int) (Music, error)
	GetLikedMusics(metadata entities.Metadata, userId int) ([]Music, error)
	LikeMusic(musicId int, userId int) error
	CountMusicByDoctorId(doctorId int) (int, error)
	CountMusicLikesByDoctorId(doctorId int) (int, error)
	CountMusicViewCountByDoctorId(doctorId int) (int, error)
	PostMusic(music Music) (Music, error)
	GetMusicByIdForDoctor(musicId int) (Music, error)
}

type UseCaseInterface interface {
	GetAllMusics(metadata entities.Metadata, userId int) ([]Music, error)
	GetAllMusicsByDoctorId(metadata entities.MetadataFull, userId int) ([]Music, error)
	GetMusicById(musicId int, userId int) (Music, error)
	GetLikedMusics(metadata entities.Metadata, userId int) ([]Music, error)
	LikeMusic(musicId int, userId int) error
	CountMusicByDoctorId(doctorId int) (int, error)
	CountMusicLikesByDoctorId(doctorId int) (int, error)
	CountMusicViewCountByDoctorId(doctorId int) (int, error)
	PostMusic(music Music, fileImage *multipart.FileHeader, fileMusic *multipart.FileHeader) (Music, error)
	GetMusicByIdForDoctor(musicId int) (Music, error)
}