package base

import (
	"capstone/constants"
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

		default:
			return http.StatusInternalServerError
		}
}