package complaint

import (
	"capstone/controllers/complaint/request"
	"capstone/controllers/complaint/response"
	complaintUseCase "capstone/entities/complaint"
	consultationEntities "capstone/entities/consultation"
	"capstone/utilities"
	"capstone/utilities/base"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ComplaintController struct {
	complaintUseCase    complaintUseCase.ComplaintUseCase
	consultationUseCase consultationEntities.ConsultationUseCase
}

func NewComplaintController(complaint complaintUseCase.ComplaintUseCase, consultationUseCase consultationEntities.ConsultationUseCase) *ComplaintController {
	return &ComplaintController{
		complaintUseCase:    complaint,
		consultationUseCase: consultationUseCase,
	}
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
	consultations, err := controller.consultationUseCase.GetAllDoctorConsultation(metadata, doctorID)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var complaints []response.ComplaintResponse
	for _, value := range *consultations {
		if value.Complaint.ID != 0 {
			complaints = append(complaints, *value.Complaint.ToResponse())
		}
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Complaints Retrieved", complaints))
}

func (controller *ComplaintController) GetByComplaintID(c echo.Context) error {
	strComplaintID := c.Param("id")
	complaintID, err := strconv.Atoi(strComplaintID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	consultation, err := controller.consultationUseCase.GetConsultationByComplaintID(complaintID)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Complaint Retrieved", consultation.ToDoctorResponse()))
}
