package post

import (
	"capstone/entities"
	postEntities "capstone/entities/post"
	"capstone/entities/user"

	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (postRepo *PostRepo) GetAllPostsByForumId(forumId uint, metadata entities.Metadata) ([]postEntities.Post, error) {
	var posts []Post
	err := postRepo.db.Limit(metadata.Limit).Offset((metadata.Page-1)*metadata.Limit).Where("forum_id = ?", forumId).Preload("User").Find(&posts).Error
	if err != nil {
		return []postEntities.Post{}, err
	}

	var postEnts []postEntities.Post
	for _, post := range posts {
		postEnts = append(postEnts, postEntities.Post{
			ID:       post.ID,
			ForumId:  post.ForumID,
			Content:  post.Content,
			ImageUrl: post.ImageUrl,
			User:     user.User{
				Id:             post.User.Id,
				Username:       post.User.Username,
				ProfilePicture: post.User.ProfilePicture,
			},
		})
	}

	return postEnts, nil
}

func (postRepo *PostRepo) GetPostById(postId uint) (postEntities.Post, error) {
	var post Post
	err := postRepo.db.Where("id = ?", postId).Preload("User").Find(&post).Error
	if err != nil {
		return postEntities.Post{}, err
	}
	
	var postEnt postEntities.Post
	postEnt.ID = post.ID
	postEnt.ForumId = post.ForumID
	postEnt.Content = post.Content
	postEnt.ImageUrl = post.ImageUrl
	postEnt.User = user.User{
		Id:             post.User.Id,
		Username:       post.User.Username,
		ProfilePicture: post.User.ProfilePicture,
	}
	return postEnt, nil
}