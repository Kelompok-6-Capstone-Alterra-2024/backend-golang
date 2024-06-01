package forum

import (
	"capstone/constants"
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