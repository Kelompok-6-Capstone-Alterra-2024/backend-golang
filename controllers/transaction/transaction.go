package transaction

import (
	"capstone/controllers/transaction/request"
	transactionEntities "capstone/entities/transaction"
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

	response, err := controller.transactionUseCase.Insert(transactionRequest.ToEntities())
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, response.ToResponse())
}
