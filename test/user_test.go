package test

import (
	"capstone/controllers/user/request"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserLoginSuccess(t *testing.T) {
	e := echo.New()

	userLoginReq := request.UserLoginRequest{
		Username: "wahyu1",
		Password: "wahyu1",
	}
	output, err := json.Marshal(userLoginReq)
	require.NoError(t, err)

	reqBody := strings.NewReader(string(output))
	req := httptest.NewRequest("POST", "/v1/users/login", reqBody)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err = userCont.Login(c)
	fmt.Println(rec.Body.String())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestUserLoginFailed(t *testing.T) {
	e := echo.New()

	userLoginReq := request.UserLoginRequest{
		Username: "wahyu1",
		Password: "wahyu",
	}
	output, err := json.Marshal(userLoginReq)
	require.NoError(t, err)

	reqBody := strings.NewReader(string(output))
	req := httptest.NewRequest("POST", "/v1/users/login", reqBody)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err = userCont.Login(c)
	fmt.Println(rec.Body.String())
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUserLoginEmptyPassword(t *testing.T) {
	e := echo.New()

	userLoginReq := request.UserLoginRequest{
		Username: "wahyu1",
		Password: "",
	}
	output, err := json.Marshal(userLoginReq)
	require.NoError(t, err)

	reqBody := strings.NewReader(string(output))
	req := httptest.NewRequest("POST", "/v1/users/login", reqBody)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err = userCont.Login(c)
	fmt.Println(rec.Body.String())
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}
