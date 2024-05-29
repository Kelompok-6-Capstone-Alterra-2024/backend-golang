package story

import (
	"capstone/constants"
	"capstone/entities"
	doctorEntities "capstone/entities/doctor"
	storyEntities "capstone/entities/story"

	"gorm.io/gorm"
)

type StoriesRepo struct {
	DB *gorm.DB
}

func NewStoryRepo(db *gorm.DB) *StoriesRepo {
	return &StoriesRepo{
		DB: db,
	}
}

func (repository *StoriesRepo) GetAllStories(metadata entities.Metadata, userId int) ([]storyEntities.Story, error) {
	var storiesDb []Story

	err := repository.DB.Limit(metadata.Limit).Offset((metadata.Page-1)*metadata.Limit).Preload("Doctor").Find(&storiesDb).Error

	if err != nil {
		return nil, constants.ErrDataNotFound
	}

	storyLikes := make([]StoryLikes, len(storiesDb))
	var counter int64
	var isLiked []bool

	for i := 0; i < len(storiesDb); i++ {
		storyLikes[i].UserId = uint(userId)
		storyLikes[i].StoryId = storiesDb[i].ID
		err = repository.DB.Model(&storyLikes[i]).Where("user_id = ? AND story_id = ?", storyLikes[i].UserId, storyLikes[i].StoryId).Count(&counter).Error

		if err != nil {
			return nil, constants.ErrServer
		}

		if counter > 0 {
			isLiked = append(isLiked, true)
		} else {
			isLiked = append(isLiked, false)
		}

		counter = 0
	}

	storiesEnt := make([]storyEntities.Story, len(storiesDb))
	for i := 0; i < len(storiesDb); i++ {
		storiesEnt[i] = storyEntities.Story{
			Id:       storiesDb[i].ID,
			Title:    storiesDb[i].Title,
			Content:  storiesDb[i].Content,
			Date:     storiesDb[i].Date,
			ImageUrl: storiesDb[i].ImageUrl,
			ViewCount: storiesDb[i].ViewCount,
			DoctorId: storiesDb[i].DoctorId,
			Doctor: doctorEntities.Doctor{
				ID:   storiesDb[i].Doctor.ID,
				Name: storiesDb[i].Doctor.Name,
			},
			IsLiked: isLiked[i],
		}
	}

	return storiesEnt, nil
}

func (repository *StoriesRepo) GetStoryById(id int) (storyEntities.Story, error) {
	var storyDb Story
	err := repository.DB.Where("id = ?", id).Preload("Doctor").First(&storyDb).Error
	if err != nil {
		return storyEntities.Story{}, constants.ErrDataNotFound
	}

	storyResp := storyEntities.Story{
		Id:       storyDb.ID,
		Title:    storyDb.Title,
		Content:  storyDb.Content,
		Date:     storyDb.Date,
		ImageUrl: storyDb.ImageUrl,
		ViewCount: storyDb.ViewCount,
		DoctorId: storyDb.DoctorId,
		Doctor: doctorEntities.Doctor{
			ID:   storyDb.Doctor.ID,
			Name: storyDb.Doctor.Name,
		},
	}

	return storyResp, nil
}