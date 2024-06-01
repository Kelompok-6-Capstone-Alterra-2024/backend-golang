package forum

import (
	"capstone/constants"
	"capstone/entities"
	"capstone/entities/forum"
	userEntities "capstone/entities/user"
	"capstone/repositories/mysql/user"

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

func (f *ForumRepo) GetJoinedForum(userId uint, metadata entities.Metadata) ([]forum.Forum, error) {
	var forumMemberDB []ForumMember

	err := f.db.Preload("Forum").Where("user_id = ?", userId).Limit(metadata.Limit).Offset((metadata.Page - 1) * metadata.Limit).Find(&forumMemberDB).Error
	if err != nil {
		return nil, constants.ErrServer
	}

	counter := make([]int64, len(forumMemberDB))
	for i, forumMemberDBTemp := range forumMemberDB {
		err = f.db.Model(&ForumMember{}).Where("forum_id = ?", forumMemberDBTemp.Forum.ID).Count(&counter[i]).Error
		if err != nil {
			return nil, constants.ErrServer
		}
	}

	forumEnts := make([]forum.Forum, len(forumMemberDB))
	for i, forumMemberDBTemp := range forumMemberDB {
		forumEnts[i].ID = forumMemberDBTemp.Forum.ID
		forumEnts[i].Name = forumMemberDBTemp.Forum.Name
		forumEnts[i].ImageUrl = forumMemberDBTemp.Forum.ImageUrl
		forumEnts[i].NumberOfMembers = int(counter[i])
	}

	var forumMemberDBTemps []ForumMember
	for i:=0; i < len(forumEnts); i++ {
		err = f.db.Where("forum_id = ? AND user_id != ?", forumEnts[i].ID, userId).Limit(2).Find(&forumMemberDBTemps).Error
		if err != nil {
			return nil, constants.ErrServer
		}

		for _, forumMemberDBTemp := range forumMemberDBTemps {
			var userDB user.User
			err = f.db.Where("id = ?", forumMemberDBTemp.UserID).First(&userDB).Error
			if err != nil {
				return nil, constants.ErrServer
			}

			forumEnts[i].User = append(forumEnts[i].User, userEntities.User{
				Id:             userDB.Id,
				ProfilePicture: userDB.ProfilePicture,
			})
		}
	}

	return forumEnts, nil
}