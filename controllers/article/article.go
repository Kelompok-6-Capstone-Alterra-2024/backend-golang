package article

import (
	"capstone/controllers/article/request"
	"capstone/controllers/article/response"
	articleUseCase "capstone/entities/article"
	"capstone/utilities"
	"capstone/utilities/base"

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
	defer fileContent.Close() // Make sure to close the file after uploading

	// Upload gambar ke Cloudinary
	imageUpload, err := utilities.UploadImage(fileContent, "article_images/"+file.Filename, "article_images")
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse("Failed to upload image"))
	}

	articleRequest := createRequest.ToArticleEntities()
	articleRequest.ImageURL = imageUpload

	// Create article in repository
	createdArticle, err := controller.articleUseCase.CreateArticle(articleRequest)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleResponse := createdArticle.ToResponse()
	return c.JSON(base.ConvertResponseCode(err), base.NewSuccessResponse("Success Create Article", articleResponse))
}

func (controller *ArticleController) GetAllArticle(c echo.Context) error {
	articles, err := controller.articleUseCase.GetAllArticle()
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleResponse := make([]response.ArticleCreatedResponse, 0, len(articles))
	for _, article := range articles {
		articleResponse = append(articleResponse, article.ToResponse())
	}

	return c.JSON(base.ConvertResponseCode(err), base.NewSuccessResponse("Success Get All Articles", articleResponse))
}
