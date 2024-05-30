package mood

import (
	"capstone/constants"
	moodEntities "capstone/entities/mood"
	"capstone/utilities"
	"mime/multipart"
)

type MoodUseCase struct {
	moodInterface moodEntities.RepositoryInterface
}

func NewMoodUseCase(moodInterface moodEntities.RepositoryInterface) *MoodUseCase {
	return &MoodUseCase{
		moodInterface: moodInterface,
	}
}

func (moodUseCase *MoodUseCase) SendMood(file *multipart.FileHeader, mood moodEntities.Mood) (moodEntities.Mood, error) {
	if mood.MoodTypeId == 0 || mood.Date == "" {
		return moodEntities.Mood{}, constants.ErrEmptyInputMood
	}

	if file != nil {
		secureUrl, err := utilities.UploadImage(file)
		if err != nil {
			return moodEntities.Mood{}, constants.ErrUploadImage
		}
		mood.ImageUrl = secureUrl
	}

	mood, err := moodUseCase.moodInterface.SendMood(mood)
	if err != nil {
		return moodEntities.Mood{}, err
	}
	
	return mood, nil
}