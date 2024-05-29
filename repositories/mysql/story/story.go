package story

import (
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
		return nil, err
	}

	storyLikes := make([]StoryLikes, len(storiesDb))
	var counter int64
	var isLiked []bool

	for i := 0; i < len(storiesDb); i++ {
		storyLikes[i].UserId = uint(userId)
		storyLikes[i].StoryId = storiesDb[i].ID
		err = repository.DB.Model(&storyLikes[i]).Where("user_id = ? AND story_id = ?", storyLikes[i].UserId, storyLikes[i].StoryId).Count(&counter).Error

		if err != nil {
			return nil, err
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