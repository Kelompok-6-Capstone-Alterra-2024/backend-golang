package transaction

import (
	"capstone/constants"
	"capstone/controllers/transaction/request"
	"capstone/controllers/transaction/response"
	midtransEntities "capstone/entities/midtrans"
	transactionEntities "capstone/entities/transaction"
	"capstone/utilities"
	"capstone/utilities/base"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransactionController struct {
	transactionUseCase transactionEntities.TransactionUseCase
	midtransUseCase    midtransEntities.MidtransUseCase
}

func NewTransactionController(transactionUseCase transactionEntities.TransactionUseCase, midtransUseCase midtransEntities.MidtransUseCase) *TransactionController {
	return &TransactionController{
		transactionUseCase: transactionUseCase,
		midtransUseCase:    midtransUseCase,
	}
}

func (controller *TransactionController) InsertWithBuiltIn(c echo.Context) error {
	var transactionRequest request.TransactionRequest
	if err := c.Bind(&transactionRequest); err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	transactionResponse, err := controller.transactionUseCase.InsertWithBuiltInInterface(transactionRequest.ToEntities())
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

	transactions, err := controller.transactionUseCase.FindAllByUserID(metadata, uint(userId))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var transactionResponses []response.TransactionResponse
	for _, transaction := range *transactions {
		transactionResponses = append(transactionResponses, *transaction.ToResponse())
	}
	return c.JSON(http.StatusOK, transactionResponses)
}

func (controller *TransactionController) BankTransfer(c echo.Context) error {
	var transaction request.TransactionRequest
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	bankName := c.QueryParam("bank")
	transaction.Bank = bankName
	transaction.PaymentType = constants.BankTransfer
	transactionRequest := transaction.ToEntities()
	transactionResponse, err := controller.transactionUseCase.InsertWithCustomInterface(transactionRequest)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return c.JSON(201, base.NewSuccessResponse("Transaction created", transactionResponse.ToResponse()))
}

func (controller *TransactionController) EWallet(c echo.Context) error {
	var transaction request.TransactionRequest
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	transaction.PaymentType = constants.GoPay
	transactionRequest := transaction.ToEntities()
	transactionResponse, err := controller.transactionUseCase.InsertWithCustomInterface(transactionRequest)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return c.JSON(201, base.NewSuccessResponse("Transaction created", transactionResponse.ToResponse()))
}

func (controller *TransactionController) CallbackTransaction(c echo.Context) error {
	var transactionCallback response.TransactionCallback
	if err := c.Bind(&transactionCallback); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	statusCode, err := controller.midtransUseCase.VerifyPayment(transactionCallback.OrderID)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	fmt.Println(transactionCallback.TransactionID)
	transaction, err := controller.transactionUseCase.ConfirmedPayment(transactionCallback.OrderID, statusCode)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Payment confirmed", transaction.ToResponse()))
}

func (controller *TransactionController) CountTransactionByDoctorID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	doctorId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	count, err := controller.transactionUseCase.CountTransactionByDoctorID(uint(doctorId))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Count transaction success", count))
}

func (controller *TransactionController) FindAllByDoctorID(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)
	token := c.Request().Header.Get("Authorization")
	doctorId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	transactions, err := controller.transactionUseCase.FindAllByDoctorID(metadata, uint(doctorId))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var transactionResponses []response.TransactionResponse
	for _, transaction := range *transactions {
		transactionResponses = append(transactionResponses, *transaction.ToResponse())
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Get all transaction success", transactionResponses))
}
