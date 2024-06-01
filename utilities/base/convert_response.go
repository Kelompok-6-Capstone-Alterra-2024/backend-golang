package base

import (
	"capstone/constants"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConvertResponseCode(err error) int {
	switch err {
	case constants.ErrEmptyInputUser:
		return http.StatusBadRequest

	case constants.ErrHashedPassword:
		return http.StatusInternalServerError

	case constants.ErrInsertDatabase:
		return http.StatusInternalServerError

	case constants.ErrUsernameAlreadyExist:
		return http.StatusConflict

	case constants.ErrEmailAlreadyExist:
		return http.StatusConflict

	case constants.ErrEmptyInputLogin:
		return http.StatusBadRequest

	case constants.ErrUserNotFound:
		return http.StatusNotFound

	case constants.ErrDataNotFound:
		return http.StatusNotFound

	case constants.ErrInvalidToken:
		return http.StatusUnauthorized

	case constants.ErrServer:
		return http.StatusInternalServerError

	case constants.ErrInvalidRate:
		return http.StatusBadRequest

	case constants.ErrCloudinary:
		return http.StatusInternalServerError

	case constants.ErrEmptyInputMood:
		return http.StatusBadRequest

	case constants.ErrUploadImage:
		return http.StatusInternalServerError

	case constants.ErrEmptyRangeDateMood:
		return http.StatusBadRequest

	case constants.ErrInvalidStartDate:
		return http.StatusBadRequest

	case constants.ErrInvalidEndDate:
		return http.StatusBadRequest

	case constants.ErrStartDateGreater:
		return http.StatusBadRequest

	case constants.ErrAlreadyLiked:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}

func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}

	switch code {
	case http.StatusNotFound:
		c.JSON(code, NewErrorResponse("resource not found"))
		return
	case http.StatusBadRequest:
		c.JSON(code, NewErrorResponse(err.Error()))
		return
	case http.StatusUnauthorized:
		c.JSON(code, NewErrorResponse("unauthorized"))
	}
}
