package forum

import (
	"capstone/constants"
	"capstone/entities"
	forumEntities "capstone/entities/forum"
)

type ForumUseCase struct {
	forumInterface forumEntities.RepositoryInterface
}

func NewForumUseCase(forumInterface forumEntities.RepositoryInterface) *ForumUseCase {
	return &ForumUseCase{
		forumInterface: forumInterface,
	}
}

func (forumUseCase *ForumUseCase) JoinForum(forumId uint, userId uint) error {
	if forumId == 0 {
		return constants.ErrEmptyInputForum
	}

	err := forumUseCase.forumInterface.JoinForum(forumId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (forumUseCase *ForumUseCase) GetJoinedForum(userId uint, metadata entities.Metadata) ([]forumEntities.Forum, error) {
	forums, err := forumUseCase.forumInterface.GetJoinedForum(userId, metadata)
	if err != nil {
		return nil, err
	}
	return forums, nil
}