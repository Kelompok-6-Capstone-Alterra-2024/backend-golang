package article

import (
	"capstone/controllers/article/request"
	"capstone/controllers/article/response"
	articleUseCase "capstone/entities/article"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	articleUseCase articleUseCase.ArticleUseCaseInterface
}

func NewArticleController(articleUseCase articleUseCase.ArticleUseCaseInterface) *ArticleController {
	return &ArticleController{
		articleUseCase: articleUseCase,
	}
}

func (controller *ArticleController) CreateArticle(c echo.Context) error {
	newArticle := new(request.CreateArticleRequest)
	if err := c.Bind(newArticle); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid request body"))
	}

	token := c.Request().Header.Get("Authorization")
	userId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse("Invalid token"))
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid image format"))
	}

	imageURL, err := utilities.UploadImage(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse("Failed to upload image"))
	}

	articleEntity := articleUseCase.Article{
		Title:     newArticle.Title,
		Content:   newArticle.Content,
		ImageUrl:  imageURL,
		DoctorID:  uint(userId),
		ViewCount: 0,
	}

	createdArticle, err := controller.articleUseCase.CreateArticle(&articleEntity, userId)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleResponse := response.ArticleCreatedResponse{
		ID:        createdArticle.ID,
		DoctorID:  createdArticle.DoctorID,
		Title:     createdArticle.Title,
		Content:   createdArticle.Content,
		Date:      createdArticle.Date,
		ImageUrl:  createdArticle.ImageUrl,
		ViewCount: createdArticle.ViewCount,
		Doctor: response.DoctorInfoResponse{
			ID:   createdArticle.Doctor.ID,
			Name: createdArticle.Doctor.Name,
		},
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Article created successfully", articleResponse))
}
func (controller *ArticleController) GetAllArticle(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	articles, err := controller.articleUseCase.GetAllArticle(*metadata, userId)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleResponse := make([]response.ArticleListResponse, 0, len(articles))
	for _, article := range articles {
		articleResponse = append(articleResponse, article.ToResponse())
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Articles", articleResponse))
}

func (controller *ArticleController) GetArticleById(c echo.Context) error {
	strId := c.Param("id")
	articleId, _ := strconv.Atoi(strId)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	article, err := controller.articleUseCase.GetArticleById(articleId, userId)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleResp := response.ArticleCreatedResponse{
		ID:        article.ID,
		DoctorID:  article.DoctorID,
		Title:     article.Title,
		Content:   article.Content,
		Date:      article.Date,
		ImageUrl:  article.ImageUrl,
		ViewCount: article.ViewCount,
		IsLiked:   article.IsLiked,
		Doctor: response.DoctorInfoResponse{
			ID:   article.Doctor.ID,
			Name: article.Doctor.Name,
		},
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Article By Id", articleResp))
}

func (controller *ArticleController) GetLikedArticle(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	userId, _ := utilities.GetUserIdFromToken(token)

	articles, err := controller.articleUseCase.GetLikedArticle(*metadata, userId)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleResponse := make([]response.ArticleCreatedResponse, len(articles))
	for i, article := range articles {
		articleResponse[i] = response.ArticleCreatedResponse{
			ID:        article.ID,
			Title:     article.Title,
			Content:   article.Content,
			Date:      article.Date,
			ImageUrl:  article.ImageUrl,
			ViewCount: article.ViewCount,
			IsLiked:   article.IsLiked,
			Doctor: response.DoctorInfoResponse{
				ID:   article.Doctor.ID,
				Name: article.Doctor.Name,
			},
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get Liked Articles", metadata, articleResponse))
}