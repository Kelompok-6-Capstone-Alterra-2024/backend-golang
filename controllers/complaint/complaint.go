package complaint

import (
	"capstone/controllers/complaint/request"
	complaintUseCase "capstone/entities/complaint"
	"capstone/utilities/base"
	"github.com/labstack/echo/v4"
	"net/http"
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
