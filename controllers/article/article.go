package article

import (
	"capstone/controllers/article/request"
	"capstone/controllers/article/response"
	articleUseCase "capstone/entities/article"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"
	"strconv"
	"time"

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
	// Binding request data
	var createRequest request.CreateArticleRequest

	token := c.Request().Header.Get("Authorization")
	doctorId, _ := utilities.GetUserIdFromToken(token)

	if err := c.Bind(&createRequest); err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse("Failed to get image from form"))
	}

	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse("Failed to open image file"))
	}
	defer fileContent.Close()

	// Upload image to Cloudinary
	imageUpload, err := utilities.UploadImage(fileContent, "article_images/"+file.Filename, "article_images")
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse("Failed to upload image"))
	}

	// Set the timezone to Asia/Jakarta
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse("Failed to load timezone"))
	}

	// Get the current time in the specified timezone
	currentTime := time.Now().In(loc)

	articleRequest := createRequest.ToArticleEntities()
	articleRequest.ImageUrl = imageUpload
	articleRequest.Date = currentTime
	articleRequest.DoctorID = uint(doctorId)

	// Create article in repository
	createdArticle, err := controller.articleUseCase.CreateArticle(articleRequest)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleResponse := createdArticle.ToResponse()
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Create Article", articleResponse))
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

	articleResponse := make([]response.ArticleCreatedResponse, 0, len(articles))
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
		Title:     article.Title,
		Content:   article.Content,
		Date:      article.Date,
		ImageURL:  article.ImageUrl,
		ViewCount: article.ViewCount,
		IsLiked:   article.IsLiked,
		Doctor: response.DoctorGetAllResponse{
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
			ImageURL:  article.ImageUrl,
			ViewCount: article.ViewCount,
			IsLiked:   article.IsLiked,
			Doctor: response.DoctorGetAllResponse{
				ID:   article.Doctor.ID,
				Name: article.Doctor.Name,
			},
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get Liked Articles", metadata, articleResponse))
}
