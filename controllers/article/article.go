package article

import (
	"capstone/controllers/article/request"
	"capstone/controllers/article/response"
	articleUseCase "capstone/entities/article"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"
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
	doctorId, err := utilities.GetUserIdFromToken(token)

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

	articleRequest := createRequest.ToArticleEntities()
	articleRequest.ImageUrl = imageUpload
	articleRequest.Date = time.Now()
	articleRequest.DoctorID = uint(doctorId)

	// Create article in repository
	createdArticle, err := controller.articleUseCase.CreateArticle(articleRequest)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleResponse := createdArticle.ToResponse()
	return c.JSON(base.ConvertResponseCode(err), base.NewSuccessResponse("Success Create Article", articleResponse))
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
