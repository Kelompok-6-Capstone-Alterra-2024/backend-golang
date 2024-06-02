package post

import (
	"capstone/controllers/post/request"
	"capstone/controllers/post/response"
	postEntities "capstone/entities/post"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	postUseCase postEntities.UseCaseInterface
}

func NewPostController(postUseCase postEntities.UseCaseInterface) *PostController {
	return &PostController{
		postUseCase: postUseCase,
	}
}

func (postController *PostController) GetAllPostsByForumId(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	forumId := c.Param("forumId")
	forumIdInt, _ := strconv.Atoi(forumId)

	posts, err := postController.postUseCase.GetAllPostsByForumId(uint(forumIdInt), *metadata)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var resp []response.PostResponse

	for _, post := range posts {
		resp = append(resp, response.PostResponse{
			ID:       post.ID,
			Content:  post.Content,
			ImageUrl: post.ImageUrl,
			User: response.UserPostResponse{
				ID:             uint(post.User.Id),
				Username:       post.User.Username,
				ProfilePicture: post.User.ProfilePicture,
			},
		})
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Posts By Forum Id", metadata, resp))
}

func (postController *PostController) GetPostById(c echo.Context) error {
	postId := c.Param("id")
	postIdInt, _ := strconv.Atoi(postId)

	post, err := postController.postUseCase.GetPostById(uint(postIdInt))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var resp response.PostResponse
	resp.ID = post.ID
	resp.Content = post.Content
	resp.ImageUrl = post.ImageUrl
	resp.User = response.UserPostResponse{
		ID:             uint(post.User.Id),
		Username:       post.User.Username,
		ProfilePicture: post.User.ProfilePicture,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Post By Id", resp))
}

func (postController *PostController) SendPost(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	var postReq request.PostSendRequest
	c.Bind(&postReq)

	file, _ := c.FormFile("image")

	postEnt := postEntities.Post{
		ForumId:  postReq.ForumId,
		UserId:   uint(userId),
		Content:  postReq.Content,
		ImageUrl: postReq.ImageUrl,
	}

	result, err := postController.postUseCase.SendPost(postEnt, file)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var resp response.PostCreateResponse
	resp.ID = result.ID
	resp.ForumId = result.ForumId
	resp.Content = result.Content
	resp.ImageUrl = result.ImageUrl
	resp.User = response.UserPostResponse{
		ID:             uint(result.User.Id),
		Username:       result.User.Username,
		ProfilePicture: result.User.ProfilePicture,
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Send Post", resp))
}

func (postController *PostController) LikePost(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	var postLikeReq request.PostLikeRequest
	c.Bind(&postLikeReq)

	err := postController.postUseCase.LikePost(uint(postLikeReq.PostId), uint(userId))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Like Post", nil))
}