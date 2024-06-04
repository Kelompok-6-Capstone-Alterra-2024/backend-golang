package transaction

import (
	"capstone/controllers/transaction/request"
	"capstone/controllers/transaction/response"
	transactionEntities "capstone/entities/transaction"
	"capstone/utilities"
	"capstone/utilities/base"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransactionController struct {
	transactionUseCase transactionEntities.TransactionUseCase
}

func NewTransactionController(transactionUseCase transactionEntities.TransactionUseCase) *TransactionController {
	return &TransactionController{
		transactionUseCase: transactionUseCase,
	}
}

func (controller *TransactionController) Insert(c echo.Context) error {
	var transactionRequest request.TransactionRequest
	if err := c.Bind(&transactionRequest); err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	transactionResponse, err := controller.transactionUseCase.Insert(transactionRequest.ToEntities())
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, transactionResponse.ToResponse())
}

func (controller *TransactionController) FindByID(c echo.Context) error {
	id := c.Param("id")

	transactionResponse, err := controller.transactionUseCase.FindByID(id)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, transactionResponse.ToResponse())
}

func (controller *TransactionController) FindByConsultationID(c echo.Context) error {
	c.Param("id")
	transactionResponse, err := controller.transactionUseCase.FindByConsultationID(1)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, transactionResponse.ToResponse())
}

func (controller *TransactionController) FindAll(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)
	token := c.Request().Header.Get("Authorization")
	userId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	transactions, err := controller.transactionUseCase.FindAll(metadata, uint(userId))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var transactionResponses []response.TransactionResponse
	for _, transaction := range *transactions {
		transactionResponses = append(transactionResponses, *transaction.ToResponse())
	}
	return c.JSON(http.StatusOK, transactionResponses)
}
