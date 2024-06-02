package post

import (
	"capstone/entities"
	postEntities "capstone/entities/post"
)

type PostUseCase struct {
	postRepository postEntities.RepositoryInterface
}

func NewPostUseCase(postRepository postEntities.RepositoryInterface) *PostUseCase {
	return &PostUseCase{
		postRepository: postRepository,
	}
}

func (postUseCase *PostUseCase) GetAllPostsByForumId(forumId uint, metadata entities.Metadata) ([]postEntities.Post, error) {
	posts, err := postUseCase.postRepository.GetAllPostsByForumId(forumId, metadata)
	if err != nil {
		return []postEntities.Post{}, err
	}
	return posts, nil
}

func (postUseCase *PostUseCase) GetPostById(postId uint) (postEntities.Post, error) {
	post, err := postUseCase.postRepository.GetPostById(postId)
	if err != nil {
		return postEntities.Post{}, err
	}
	return post, nil
}