package base

import (
	"capstone/constants"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
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

	case constants.ErrEmptyInputForum:
		return http.StatusBadRequest

	case constants.ErrEmptyInputPost:
		return http.StatusBadRequest

	case constants.ErrEmptyInputLike:
		return http.StatusBadRequest

	case constants.ErrEmptyInputComment:
		return http.StatusBadRequest

	case constants.ErrExcange:
		return http.StatusInternalServerError

	case constants.ErrNewServiceGoogle:
		return http.StatusInternalServerError

	case constants.ErrNewUserInfo:
		return http.StatusInternalServerError

	case constants.ErrInsertOAuth:
		return http.StatusInternalServerError

	case constants.ErrEmptyInputMusic:
		return http.StatusBadRequest

	case constants.ErrEmptyInputStory:
		return http.StatusBadRequest

	case constants.ErrInvalidConsultationID:
		return http.StatusBadRequest

	case constants.ErrEmptyCreateForum:
		return http.StatusBadRequest

	case constants.ErrEmptyChat:
		return http.StatusBadRequest

	case constants.ErrEmptyInputEmailOTP:
		return http.StatusBadRequest

	case constants.ErrEmptyInputVerifyOTP:
		return http.StatusBadRequest

	case constants.ErrInvalidOTP:
		return http.StatusUnauthorized

	case constants.ErrExpiredOTP:
		return http.StatusGone
		
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
