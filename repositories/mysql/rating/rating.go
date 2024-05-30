package rating

import (
	"capstone/constants"
	ratingEntities "capstone/entities/rating"

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