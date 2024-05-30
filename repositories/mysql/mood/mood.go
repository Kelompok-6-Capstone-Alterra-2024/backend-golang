package mood

import (
	"capstone/constants"
	moodEntities "capstone/entities/mood"

	"gorm.io/gorm"
)

type MoodRepo struct {
	db *gorm.DB
}

func NewMoodRepo(db *gorm.DB) *MoodRepo {
	return &MoodRepo{
		db: db,
	}
}

func (moodRepo *MoodRepo) SendMood(mood moodEntities.Mood) (moodEntities.Mood, error) {
	moodDB := Mood{
		UserId:     mood.UserId,
		MoodTypeId: mood.MoodTypeId,
		Message:    mood.Message,
		Date:       mood.Date,
		ImageUrl:   mood.ImageUrl,
	}

	err := moodRepo.db.Create(&moodDB).Error
	if err != nil {
		return moodEntities.Mood{}, constants.ErrServer
	}

	err = moodRepo.db.Model(&moodDB).Preload("MoodType").Find(&moodDB).Error
	if err != nil {
		return moodEntities.Mood{}, constants.ErrServer
	}

	result := moodEntities.Mood{
		ID:         moodDB.ID,
		UserId:     moodDB.UserId,
		MoodTypeId: moodDB.MoodTypeId,
		MoodType:   moodEntities.MoodType{
			ID:   moodDB.MoodType.ID,
			Name: moodDB.MoodType.Name,
		},
		Message:    moodDB.Message,
		Date:       moodDB.Date,
		ImageUrl:   moodDB.ImageUrl,
	}

	return result, nil
}