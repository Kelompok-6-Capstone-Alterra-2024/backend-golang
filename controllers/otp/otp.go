package otp

import (
	"capstone/controllers/otp/request"
	otpEntities "capstone/entities/otp"
	"capstone/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OtpController struct {
	otpUseCase otpEntities.UseCaseInterface
}

func NewOtpController(otpUseCase otpEntities.UseCaseInterface) *OtpController {
	return &OtpController{
		otpUseCase: otpUseCase,
	}
}

func (otpController *OtpController) SendOtp(c echo.Context) error {
	var req request.OTPRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	otpEnt := otpEntities.Otp{
		Email: req.Email,
	}

	err = otpController.otpUseCase.SendOTP(otpEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success send otp", nil))
}

func (otpController *OtpController) VerifyOtp(c echo.Context) error {
	var req request.OTPVerifyRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	otpEnt := otpEntities.Otp{
		Email: req.Email,
		Code:  req.Code,
	}

	err = otpController.otpUseCase.VerifyOTP(otpEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success verify otp", nil))
}