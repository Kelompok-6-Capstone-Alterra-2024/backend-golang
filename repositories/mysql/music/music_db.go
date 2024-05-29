package music

import (
	"capstone/repositories/mysql/user"

	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100)"`
	Singer      string `gorm:"type:varchar(100)"`
	MusicUrl    string `gorm:"type:varchar(255)"`
	ImageUrl    string `gorm:"type:varchar(255)"`
	ViewCount   int    `gorm:"type:int;default:0"`
}

type MusicLikes struct {
	gorm.Model
	MusicId uint `gorm:"type:int;index"`
	Music   Music `gorm:"foreignKey:music_id;references:id"`
	UserId  uint `gorm:"type:int;index"`
	User    user.User `gorm:"foreignKey:user_id;references:id"`
}