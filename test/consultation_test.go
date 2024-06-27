package test

import (
	"capstone/controllers/consultation/request"
	"capstone/middlewares"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConsultation(t *testing.T) {
	e := echo.New()
	var (
		consultationCreateReq = request.ConsultationRequest{
			DoctorID: 1,
			UserID:   1,
			Date:     "2024-10-10",
			Time:     "10:00",
		}
	)

	t.Run("Create Consultation", func(t *testing.T) {
		consultationCR, err := json.Marshal(consultationCreateReq)

		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		reqBody := strings.NewReader(string(consultationCR))
		req := httptest.NewRequest("POST", "/v1/users/consultations", reqBody)
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Add("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		err = consultationCont.CreateConsultation(c)
		fmt.Println(rec.Body.String())
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("Get Consultation", func(t *testing.T) {
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest("GET", "/v1/users/consultations", nil)
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Add("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		err = consultationCont.GetAllConsultation(c)
		fmt.Println(rec.Body.String())
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
