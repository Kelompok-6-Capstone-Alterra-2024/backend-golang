package consultation

import (
	"capstone/constants"
	"capstone/entities"
	consultationEntities "capstone/entities/consultation"
	doctorEntities "capstone/entities/doctor"
	forumEntities "capstone/entities/forum"
	musicEntities "capstone/entities/music"

	"gorm.io/gorm"
)

type ConsultationRepo struct {
	db *gorm.DB
}

func NewConsultationRepo(db *gorm.DB) consultationEntities.ConsultationRepository {
	return &ConsultationRepo{
		db: db,
	}
}

func (repository *ConsultationRepo) CreateConsultation(consultation *consultationEntities.Consultation) (*consultationEntities.Consultation, error) {
	consultationRequest := ToConsultationModel(consultation)

	if err := repository.db.Create(&consultationRequest).Error; err != nil {
		return nil, constants.ErrInsertDatabase
	}

	consultationResult := consultationRequest.ToEntities()
	return consultationResult, nil
}

func (repository *ConsultationRepo) GetConsultationByID(consultationID int) (consultation *consultationEntities.Consultation, err error) {
	var consultationResult Consultation
	if err = repository.db.Preload("User").Preload("Doctor").First(&consultationResult, consultationID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	return consultationResult.ToEntities(), nil
}

func (repository *ConsultationRepo) GetAllUserConsultation(metadata *entities.Metadata, userID int) (*[]consultationEntities.Consultation, error) {
	var consultationResult []Consultation

	if err := repository.db.Limit(metadata.Limit).Offset(metadata.Offset()).Preload("Doctor").Find(&consultationResult, "user_id LIKE ?", userID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	var consultations []consultationEntities.Consultation
	for _, consultation := range consultationResult {
		consultations = append(consultations, *consultation.ToEntities())
	}

	return &consultations, nil
}

func (repository *ConsultationRepo) UpdateStatusConsultation(consultation *consultationEntities.Consultation) (*consultationEntities.Consultation, error) {
	var consultationDB Consultation
	if err := repository.db.First(&consultationDB, "id LIKE ?", consultation.ID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	consultationDB.Status = consultation.Status

	if err := repository.db.Model(&consultationDB).Save(&consultationDB).Error; err != nil {
		return nil, err
	}

	return consultationDB.ToEntities(), nil
}

func (repository *ConsultationRepo) GetAllDoctorConsultation(metadata *entities.Metadata, doctorID int) (*[]consultationEntities.Consultation, error) {
	var consultationResult []Consultation

	if err := repository.db.Limit(metadata.Limit).Offset(metadata.Offset()).Preload("Complaint").Find(&consultationResult, "doctor_id LIKE ?", doctorID).Error; err != nil {
		return nil, constants.ErrDataNotFound
	}

	var consultations []consultationEntities.Consultation
	for _, consultation := range consultationResult {
		consultations = append(consultations, *consultation.ToEntities())
	}

	return &consultations, nil
}

func (repository *ConsultationRepo) CreateConsultationNotes(consultationNotes consultationEntities.ConsultationNotes) (consultationEntities.ConsultationNotes, error) {
	var notesDB ConstultationNotes
	notesDB.ID = consultationNotes.ID
	notesDB.ConsultationID = consultationNotes.ConsultationID
	notesDB.MusicID = consultationNotes.MusicID
	notesDB.ForumID = consultationNotes.ForumID
	notesDB.MainPoint = consultationNotes.MainPoint
	notesDB.NextStep = consultationNotes.NextStep
	notesDB.AdditionalNote = consultationNotes.AdditionalNote
	notesDB.MoodTrackerNote = consultationNotes.MoodTrackerNote

	if err := repository.db.Create(&notesDB).Error; err != nil {
		return consultationEntities.ConsultationNotes{}, constants.ErrInsertDatabase
	}

	var notesEnt consultationEntities.ConsultationNotes
	notesEnt.ID = notesDB.ID
	notesEnt.ConsultationID = notesDB.ConsultationID
	notesEnt.MusicID = notesDB.MusicID
	notesEnt.ForumID = notesDB.ForumID
	notesEnt.MainPoint = notesDB.MainPoint
	notesEnt.NextStep = notesDB.NextStep
	notesEnt.AdditionalNote = notesDB.AdditionalNote
	notesEnt.MoodTrackerNote = notesDB.MoodTrackerNote

	return notesEnt, nil
}

func (repository *ConsultationRepo) GetConsultationNotesByID(consultationID int) (consultationEntities.ConsultationNotes, error) {
	var notesDB ConstultationNotes
	err := repository.db.Preload("Music").Preload("Forum").Preload("Consultation").Preload("Consultation.Doctor").Where("consultation_id = ?", consultationID).First(&notesDB).Error

	if err != nil {
		return consultationEntities.ConsultationNotes{}, constants.ErrDataNotFound
	}

	var notesEnt consultationEntities.ConsultationNotes
	notesEnt.ID = notesDB.ID
	
	notesEnt.Consultation = consultationEntities.Consultation{
		ID:   notesDB.Consultation.ID,
		Doctor: &doctorEntities.Doctor{
			ID:   notesDB.Consultation.Doctor.ID,
			Name: notesDB.Consultation.Doctor.Name,
		},
	}

	notesEnt.Forum = forumEntities.Forum{
		ID:   notesDB.Forum.ID,
		Name: notesDB.Forum.Name,
		ImageUrl: notesDB.Forum.ImageUrl,
	}

	notesEnt.Music = musicEntities.Music{
		Id:   notesDB.Music.ID,
		Title: notesDB.Music.Title,
		ImageUrl: notesDB.Music.ImageUrl,
	}

	notesEnt.MainPoint = notesDB.MainPoint
	notesEnt.NextStep = notesDB.NextStep
	notesEnt.AdditionalNote = notesDB.AdditionalNote
	notesEnt.MoodTrackerNote = notesDB.MoodTrackerNote

	return notesEnt, nil
}