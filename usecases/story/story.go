package stories

import (
	"capstone/entities"
	storyEntities "capstone/entities/story"
)

type StoryUseCase struct {
	storyRepository storyEntities.RepositoryInterface
}

func NewStoryUseCase(storyRepository storyEntities.RepositoryInterface) *StoryUseCase {
	return &StoryUseCase{
		storyRepository: storyRepository,
	}
}

func (storiesUseCase *StoryUseCase) GetAllStories(metadata entities.Metadata, userId int) ([]storyEntities.Story, error) {
	stories, err := storiesUseCase.storyRepository.GetAllStories(metadata, userId)
	if err != nil {
		return []storyEntities.Story{}, err
	}
	return stories, nil
}

func (storiesUseCase *StoryUseCase) GetStoryById(storyId int, userId int) (storyEntities.Story, error) {
	story, err := storiesUseCase.storyRepository.GetStoryById(storyId, userId)
	if err != nil {
		return storyEntities.Story{}, err
	}
	return story, nil
}

func (storiesUseCase *StoryUseCase) GetLikedStories(metadata entities.Metadata, userId int) ([]storyEntities.Story, error) {
	stories, err := storiesUseCase.storyRepository.GetLikedStories(metadata, userId)
	if err != nil {
		return []storyEntities.Story{}, err
	}
	return stories, nil
}

func (storiesUseCase *StoryUseCase) LikeStory(storyId int, userId int) error {
	err := storiesUseCase.storyRepository.LikeStory(storyId, userId)
	if err != nil {
		return err
	}
	return nil
}