package post

import (
	"capstone/repositories/mysql/forum"
	"capstone/repositories/mysql/user"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content  string    		`gorm:"type:text"`
	UserID   uint      		`gorm:"type:int;index"`
	User     user.User 		`gorm:"foreignKey:user_id;references:id"`
	ForumID  uint      		`gorm:"type:int;index"`
	Forum    forum.Forum    `gorm:"foreignKey:forum_id;references:id"`
	ImageUrl string    		`gorm:"type:varchar(255)"`
}