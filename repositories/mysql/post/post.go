package post

import (
	"capstone/entities"
	postEntities "capstone/entities/post"
	userEntities "capstone/entities/user"
	"capstone/repositories/mysql/user"
	"time"

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
			User:     userEntities.User{
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
	postEnt.User = userEntities.User{
		Id:             post.User.Id,
		Username:       post.User.Username,
		ProfilePicture: post.User.ProfilePicture,
	}
	return postEnt, nil
}

func (postRepo *PostRepo) SendPost(post postEntities.Post) (postEntities.Post, error) {
	var postDB Post
	postDB.ForumID = post.ForumId
	postDB.Content = post.Content
	postDB.ImageUrl = post.ImageUrl
	postDB.UserID = post.UserId
	err := postRepo.db.Create(&postDB).Error
	if err != nil {
		return postEntities.Post{}, err
	}

	var userDB user.User
	err = postRepo.db.Where("id = ?", postDB.UserID).Find(&userDB).Error
	if err != nil {
		return postEntities.Post{}, err
	}
	
	var postEnt postEntities.Post
	postEnt.ID = postDB.ID
	postEnt.ForumId = postDB.ForumID
	postEnt.Content = postDB.Content
	postEnt.ImageUrl = postDB.ImageUrl
	postEnt.UserId = postDB.UserID
	postEnt.User = userEntities.User{
		Id:             userDB.Id,
		Username:       userDB.Username,
		ProfilePicture: userDB.ProfilePicture,
	}

	return postEnt, nil
}

func (postRepo *PostRepo) LikePost(postId uint, userId uint) error {
	var postLikeDB PostLike
	postLikeDB.PostID = postId
	postLikeDB.UserID = userId

	err := postRepo.db.Create(&postLikeDB).Error
	if err != nil {
		return err
	}
	return nil
}

func (postRepo *PostRepo) SendComment(comment postEntities.PostComment) (postEntities.PostComment, error) {
	var commentDB PostComment
	commentDB.Content = comment.Content
	commentDB.PostID = comment.PostID
	commentDB.UserID = comment.UserID
	err := postRepo.db.Create(&commentDB).Error
	if err != nil {
		return postEntities.PostComment{}, err
	}

	var userDB user.User
	err = postRepo.db.Where("id = ?", commentDB.UserID).Find(&userDB).Error
	if err != nil {
		return postEntities.PostComment{}, err
	}
	
	var commentEnt postEntities.PostComment
	commentEnt.ID = commentDB.ID
	commentEnt.Content = commentDB.Content
	commentEnt.PostID = commentDB.PostID
	commentEnt.UserID = commentDB.UserID
	commentEnt.CreatedAt = commentDB.CreatedAt.Format(time.RFC3339)
	commentEnt.User = userEntities.User{
		Id:             userDB.Id,
		Username:       userDB.Username,
		ProfilePicture: userDB.ProfilePicture,
	}

	return commentEnt, nil
}