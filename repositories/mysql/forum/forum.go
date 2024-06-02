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

func (f *ForumRepo) GetRecommendationForum(userId uint, metadata entities.Metadata) ([]forum.Forum, error) {
	var forumMemberDB []ForumMember
	err := f.db.Preload("Forum").Where("user_id = ?", userId).Find(&forumMemberDB).Error
	if err != nil {
		return nil, constants.ErrServer
	}

	var forumIDs []uint
	for _, forumMemberDBTemp := range forumMemberDB {
		forumIDs = append(forumIDs, forumMemberDBTemp.Forum.ID)
	}

	var forumDB []Forum
	err = f.db.Model(&Forum{}).Where("id NOT IN ?", forumIDs).Limit(metadata.Limit).Offset((metadata.Page - 1) * metadata.Limit).Find(&forumDB).Error
	if err != nil {
		return nil, constants.ErrServer
	}

	counter := make([]int64, len(forumDB))
	for i, forumDBtemp := range forumDB {
		err = f.db.Model(&ForumMember{}).Where("forum_id = ?", forumDBtemp.ID).Count(&counter[i]).Error
		if err != nil {
			return nil, constants.ErrServer
		}
	}

	var forumEnts []forum.Forum
	for i, forumDBTemp := range forumDB {
		forumEnts = append(forumEnts, forum.Forum{
			ID:               forumDBTemp.ID,
			Name:             forumDBTemp.Name,
			ImageUrl:         forumDBTemp.ImageUrl,
			NumberOfMembers:  int(counter[i]),
		})
	}

	return forumEnts, nil
}

func (f *ForumRepo) GetForumById(forumId uint) (forum.Forum, error) {
	var forumDB Forum

	err := f.db.Where("id = ?", forumId).First(&forumDB).Error
	if err != nil {
		return forum.Forum{}, constants.ErrServer
	}

	var forumEnt forum.Forum
	forumEnt.ID = forumDB.ID
	forumEnt.Name = forumDB.Name
	forumEnt.Description = forumDB.Description
	forumEnt.ImageUrl = forumDB.ImageUrl

	return forumEnt, nil
}