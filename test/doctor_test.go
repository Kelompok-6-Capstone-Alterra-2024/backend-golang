package test

import (
	"capstone/configs"
	"capstone/controllers/doctor/request"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	loginReq = request.DoctorLoginRequest{
		Username: "udin",
		Password: "udin",
	}

	loginReqFailed = request.DoctorLoginRequest{
		Username: "udin",
		Password: "udin1",
	}
)

func TestRegisterAlreadyExist(t *testing.T) {
	e := echo.New()

	doctorReq := request.DoctorRegisterRequest{
		Email:    "wahyu@gmail.com",
		Username: "wahyu",
		Password: "123456",
	}
	output, err := json.Marshal(doctorReq)
	require.NoError(t, err)

	reqBody := strings.NewReader(string(output))
	req := httptest.NewRequest("POST", "/v1/doctors/register", reqBody)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Add("X-API-KEY", configs.InitConfigJWT())
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err = doctorCont.Register(c)
	require.NoError(t, err)

	fmt.Println(rec.Body.String())

	assert.Equal(t, 409, rec.Code)
}

func TestRegisterSuccess(t *testing.T) {
	e := echo.New()

	doctorReq := request.DoctorRegisterRequest{
		Email:    "udin281@gmail.com",
		Username: "udin281",
		Password: "123456",
	}
	output, err := json.Marshal(doctorReq)
	require.NoError(t, err)

	reqBody := strings.NewReader(string(output))
	req := httptest.NewRequest("POST", "/v1/doctors/register", reqBody)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Add("X-API-KEY", configs.InitConfigJWT())
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err = doctorCont.Register(c)
	require.NoError(t, err)

	fmt.Println(rec.Body.String())

	assert.Equal(t, 200, rec.Code)
}

func TestRegisterEmptyEmail(t *testing.T) {
	e := echo.New()

	doctorReq := request.DoctorRegisterRequest{
		Email:    "",
		Username: "dasdksa",
		Password: "123456",
	}
	output, err := json.Marshal(doctorReq)
	require.NoError(t, err)

	reqBody := strings.NewReader(string(output))
	req := httptest.NewRequest("POST", "/v1/doctors/register", reqBody)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Add("X-API-KEY", configs.InitConfigJWT())
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err = doctorCont.Register(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestLoginSuccess(t *testing.T) {
	e := echo.New()

	output, err := json.Marshal(loginReq)
	require.NoError(t, err)

	reqBody := strings.NewReader(string(output))
	req := httptest.NewRequest("POST", "/v1/doctors/login", reqBody)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err = doctorCont.Login(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestLoginFailed(t *testing.T) {
	e := echo.New()

	output, err := json.Marshal(loginReqFailed)
	require.NoError(t, err)

	reqBody := strings.NewReader(string(output))
	req := httptest.NewRequest("POST", "/v1/doctors/login", reqBody)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err = doctorCont.Login(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestLoginEmptyUsername(t *testing.T) {
	e := echo.New()

	loginEmptyUser := request.DoctorLoginRequest{
		Username: "",
		Password: "udin",
	}
	output, err := json.Marshal(loginEmptyUser)
	require.NoError(t, err)

	reqBody := strings.NewReader(string(output))
	req := httptest.NewRequest("POST", "/v1/doctors/login", reqBody)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err = doctorCont.Login(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetAllDoctor(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest("GET", "/v1/doctors/1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/v1/doctors/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := doctorCont.GetByID(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestGetDoctor(t *testing.T) {
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v1/doctors/1", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/v1/doctors/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := doctorCont.GetByID(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Id Not Found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v1/doctors/100", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/v1/doctors/:id")
		c.SetParamNames("id")
		c.SetParamValues("100")

		err := doctorCont.GetByID(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Get All Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v1/doctors", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		err := doctorCont.GetAll(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Get Active Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v1/doctors/active", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		err := doctorCont.GetActive(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Search Doctor Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v1/doctors/search", nil)
		req.RequestURI = "/v1/doctors/search?name=udin"
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		err := doctorCont.SearchDoctor(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)
	})
}
