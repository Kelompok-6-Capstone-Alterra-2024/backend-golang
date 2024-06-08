package rating

import (
	"capstone/constants"
	"capstone/entities"
	ratingEntities "capstone/entities/rating"
	"capstone/entities/user"

	"gorm.io/gorm"
)

type RatingRepo struct {
	DB *gorm.DB
}

func NewRatingRepo(db *gorm.DB) *RatingRepo {
	return &RatingRepo{
		DB: db,
	}
}

func (repository *RatingRepo) SendFeedback(rating ratingEntities.Rating) (ratingEntities.Rating, error) {
	ratingDB := Rating{
		UserId:   rating.UserId,
		DoctorId: rating.DoctorId,
		Rate:     rating.Rate,
		Message:  rating.Message,
	}

	err := repository.DB.Create(&ratingDB).Error
	if err != nil {
		return ratingEntities.Rating{}, constants.ErrServer
	}

	result := ratingEntities.Rating{
		Id:       ratingDB.ID,
		UserId:   ratingDB.UserId,
		DoctorId: ratingDB.DoctorId,
		Rate:     ratingDB.Rate,
		Message:  ratingDB.Message,
	}

	return result, nil
}

func (repository *RatingRepo) GetAllFeedbacks(metadata entities.Metadata, doctorId uint) ([]ratingEntities.Rating, error) {
	var ratingDB []Rating

	err := repository.DB.Limit(metadata.Limit).Offset((metadata.Page-1)*metadata.Limit).Preload("User").Where("doctor_id = ?", doctorId).Find(&ratingDB).Error
	if err != nil {
		return []ratingEntities.Rating{}, constants.ErrDataNotFound
	}

	result := make([]ratingEntities.Rating, len(ratingDB))
	for i := 0; i < len(ratingDB); i++ {
		result[i] = ratingEntities.Rating{
			Id:       ratingDB[i].ID,
			UserId:   ratingDB[i].UserId,
			User:     user.User{
				Id:       ratingDB[i].User.Id,
				Name:     ratingDB[i].User.Name,
				Username: ratingDB[i].User.Username,
				ProfilePicture: ratingDB[i].User.ProfilePicture,
			},
			Rate:     ratingDB[i].Rate,
			Message:  ratingDB[i].Message,
			Date:     ratingDB[i].CreatedAt.String(),
		}
	}

	return result, nil
}