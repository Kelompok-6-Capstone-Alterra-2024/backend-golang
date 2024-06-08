package consultation

import (
	"capstone/controllers/consultation/request"
	"capstone/controllers/consultation/response"
	"capstone/entities/consultation"
	"capstone/utilities"
	"capstone/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ConsultationController struct {
	consultationUseCase consultation.ConsultationUseCase
}

func NewConsultationController(consultationUseCase consultation.ConsultationUseCase) *ConsultationController {
	return &ConsultationController{
		consultationUseCase,
	}
}

func (controller *ConsultationController) CreateConsultation(c echo.Context) error {
	var consultationRequest request.ConsultationRequest
	c.Bind(&consultationRequest)
	token := c.Request().Header.Get("Authorization")
	userId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	date, err := utilities.StringToDate(consultationRequest.Date)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	consultationRequest.UserID = userId
	consultationResponse, err := controller.consultationUseCase.CreateConsultation(consultationRequest.ToEntities(date))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Add Consultation", consultationResponse.ToUserResponse()))
}

func (controller *ConsultationController) GetConsultationByID(c echo.Context) error {
	consultationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID"))
	}
	consultationResponse, err := controller.consultationUseCase.GetConsultationByID(consultationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", consultationResponse.ToUserResponse()))
}

func (controller *ConsultationController) GetAllConsultation(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	userId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	consultations, err := controller.consultationUseCase.GetAllUserConsultation(metadata, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	var responses []response.ConsultationUserResponse
	for _, value := range *consultations {
		responses = append(responses, *value.ToUserResponse())
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", responses))
}

func (controller *ConsultationController) UpdateStatusConsultation(c echo.Context) error {
	var consultationRequest request.ConsultationStatusUpdateRequest

	consultationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID"))
	}
	c.Bind(&consultationRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	consultationRequest.ID = uint(consultationID)
	consultationResponse, err := controller.consultationUseCase.UpdateStatusConsultation(consultationRequest.ToEntities())
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Consultation", consultationResponse.ToUserResponse()))
}

func (controller *ConsultationController) GetAllDoctorConsultation(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	doctorId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	consultations, err := controller.consultationUseCase.GetAllDoctorConsultation(metadata, doctorId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	var responses []response.ConsultationDoctorResponse
	for _, value := range *consultations {
		responses = append(responses, *value.ToDoctorResponse())
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation", responses))
}

func (controller *ConsultationController) CreateConsultationNotes(c echo.Context) error {
	var consultationNotesRequest request.ConsultationNotesRequest
	c.Bind(&consultationNotesRequest)

	var notesEnt consultation.ConsultationNotes

	notesEnt.ConsultationID = consultationNotesRequest.ConsultationID
	notesEnt.MusicID = consultationNotesRequest.MusicID
	notesEnt.ForumID = consultationNotesRequest.ForumID
	notesEnt.MainPoint = consultationNotesRequest.MainPoint
	notesEnt.NextStep = consultationNotesRequest.NextStep
	notesEnt.AdditionalNote = consultationNotesRequest.AdditionalNote
	notesEnt.MoodTrackerNote = consultationNotesRequest.MoodTrackerNote

	res, err := controller.consultationUseCase.CreateConsultationNotes(notesEnt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	var resp response.ConsultationNotesCreateResponse
	resp.ID = res.ID
	resp.ConsultationID = res.ConsultationID
	resp.MusicID = res.MusicID
	resp.ForumID = res.ForumID
	resp.MainPoint = res.MainPoint
	resp.NextStep = res.NextStep
	resp.AdditionalNote = res.AdditionalNote
	resp.MoodTrackerNote = res.MoodTrackerNote

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Add Consultation Notes", resp))
}

func (controller *ConsultationController) GetConsultationNotesByID(c echo.Context) error {
	consultationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID"))
	}
	res, err := controller.consultationUseCase.GetConsultationNotesByID(consultationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	var resp response.ConsultationNotesDetailResponse
	resp.ID = res.ID
	resp.ConsultationID = res.Consultation.ID

	resp.Doctor = response.NotesDoctorDetailResponse{
		ID: res.Consultation.Doctor.ID,
		Name: res.Consultation.Doctor.Name,
		ImageUrl: res.Consultation.Doctor.ProfilePicture,
	}

	resp.Music = response.NotesMusicDetailResponse{
		ID: res.Music.Id,
		Title: res.Music.Title,
		ImageUrl: res.Music.ImageUrl,
	}
	
	resp.Forum = response.NotesForumDetailResponse{
		ID: res.Forum.ID,
		Name: res.Forum.Name,
		ImageUrl: res.Forum.ImageUrl,
	}
	resp.MainPoint = res.MainPoint
	resp.NextStep = res.NextStep
	resp.AdditionalNote = res.AdditionalNote
	resp.MoodTrackerNote = res.MoodTrackerNote

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Consultation Notes", resp))
}