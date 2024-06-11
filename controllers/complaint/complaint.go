package complaint

import (
	"capstone/controllers/complaint/request"
	complaintUseCase "capstone/entities/complaint"
	"capstone/utilities"
	"capstone/utilities/base"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ComplaintController struct {
	complaintUseCase complaintUseCase.ComplaintUseCase
}

func NewComplaintController(complaint complaintUseCase.ComplaintUseCase) *ComplaintController {
	return &ComplaintController{complaintUseCase: complaint}
}

func (controller *ComplaintController) Create(c echo.Context) error {
	var complaintRequest request.ComplaintRequest
	c.Bind(&complaintRequest)

	complaint, err := controller.complaintUseCase.Create(complaintRequest.ToEntities())
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Complaint Created", complaint.ToResponse()))
}

func (controller *ComplaintController) GetAllByDoctorID(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	doctorID, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	complaints, err := controller.complaintUseCase.GetAllByUserID(metadata, doctorID)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Complaints Retrieved", complaints))
}

func (controller *ComplaintController) GetByComplaintID(c echo.Context) error {
	strComplaintID := c.Param("id")
	complaintID, err := strconv.Atoi(strComplaintID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	complaint, err := controller.complaintUseCase.GetByID(complaintID)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Complaint Retrieved", complaint))
}
