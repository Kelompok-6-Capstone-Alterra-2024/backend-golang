package rating

import (
	"capstone/constants"
	ratingEntities "capstone/entities/rating"
)

type RatingUseCase struct {
	ratingRepository ratingEntities.RepositoryInterface
}

func NewRatingUseCase(ratingRepository ratingEntities.RepositoryInterface) *RatingUseCase {
	return &RatingUseCase{
		ratingRepository: ratingRepository,
	}
}

func (ratingUseCase *RatingUseCase) SendFeedback(rating ratingEntities.Rating) (ratingEntities.Rating, error) {
	if rating.Rate < 1 || rating.Rate > 5 {
		return ratingEntities.Rating{}, constants.ErrInvalidRate
	}
	
	result, err := ratingUseCase.ratingRepository.SendFeedback(rating)
	if err != nil {
		return ratingEntities.Rating{}, err
	}
	return result, nil
}