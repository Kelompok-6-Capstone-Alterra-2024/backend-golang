package doctor

import (
	"capstone/controllers/doctor/request"
	doctorUseCase "capstone/entities/doctor"
	"capstone/utilities/base"
	"github.com/labstack/echo/v4"
)

type DoctorController struct {
	doctorUseCase doctorUseCase.DoctorUseCaseInterface
}

func NewDoctorController(doctorUseCase doctorUseCase.DoctorUseCaseInterface) *DoctorController {
	return &DoctorController{
		doctorUseCase: doctorUseCase,
	}
}

func (controller *DoctorController) Register(c echo.Context) error {
	var doctorFromRequest request.DoctorRegisterRequest
	c.Bind(&doctorFromRequest)

	doctorRequest := doctorFromRequest.ToDoctorEntities()
	doctorResult, err := controller.doctorUseCase.Register(doctorRequest)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	doctorResponse := doctorResult.ToResponse()
	return c.JSON(base.ConvertResponseCode(err), base.NewSuccessResponse("Success Register", doctorResponse))
}

func (controller *DoctorController) Login(c echo.Context) error {
	var doctorFromRequest request.DoctorLoginRequest
	c.Bind(&doctorFromRequest)

	doctorRequest := doctorFromRequest.ToDoctorLoginEntities()
	doctorResult, err := controller.doctorUseCase.Login(doctorRequest)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	doctorResponse := doctorResult.ToResponse()
	return c.JSON(base.ConvertResponseCode(err), base.NewSuccessResponse("Success Login", doctorResponse))
}
