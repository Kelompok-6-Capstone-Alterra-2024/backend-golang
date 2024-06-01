package forum

import (
	"capstone/constants"

	"gorm.io/gorm"
)

type ForumRepo struct {
	db *gorm.DB
}

func NewForumRepo(db *gorm.DB) *ForumRepo {
	return &ForumRepo{
		db: db,
	}
}

func (f *ForumRepo) JoinForum(forumId uint, userId uint) error {
	var forumMemberDB ForumMember
	forumMemberDB.ForumID = forumId
	forumMemberDB.UserID = userId

	err := f.db.Create(&forumMemberDB).Error
	if err != nil {
		return constants.ErrServer
	}
	return nil
}