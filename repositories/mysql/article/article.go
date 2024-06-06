package article

import (
	"capstone/constants"
	"capstone/entities"
	articleEntities "capstone/entities/article"
	doctorEntities "capstone/entities/doctor"
	"time"

	"gorm.io/gorm"
)

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{
		db: db,
	}
}

func (repository *ArticleRepo) CreateArticle(article *articleEntities.Article, userId int) (*articleEntities.Article, error) {
	articleDB := Article{
		Title:     article.Title,
		Content:   article.Content,
		ImageUrl:  article.ImageUrl,
		DoctorID:  uint(userId), // Assuming userId is the ID of the doctor creating the article
		ViewCount: 0,            // Initialize view count to 0
		Date:      time.Now(),
	}

	err := repository.db.Create(&articleDB).Error
	if err != nil {
		return nil, err
	}

	// Retrieve the created article with doctor details
	err = repository.db.Where("id = ?", articleDB.ID).Preload("Doctor").First(&articleDB).Error
	if err != nil {
		return nil, constants.ErrDataNotFound
	}

	articleResp := articleEntities.Article{
		ID:        articleDB.ID,
		Title:     articleDB.Title,
		Content:   articleDB.Content,
		Date:      articleDB.Date,
		ImageUrl:  articleDB.ImageUrl,
		ViewCount: articleDB.ViewCount,
		DoctorID:  articleDB.DoctorID,
		Doctor: doctorEntities.Doctor{
			ID:   articleDB.Doctor.ID,
			Name: articleDB.Doctor.Name,
		},
	}

	return &articleResp, nil
}

func (repository *ArticleRepo) GetAllArticle(metadata entities.Metadata, userId int) ([]articleEntities.Article, error) {
	var articlesDb []Article

	// Pagination
	err := repository.db.Limit(metadata.Limit).
		Offset((metadata.Page - 1) * metadata.Limit).
		Preload("Doctor"). // Preload doctor data
		Find(&articlesDb).Error
	if err != nil {
		return nil, err
	}

	articleLikes := make([]ArticleLikes, len(articlesDb))
	var counter int64
	var isLiked []bool

	// Check if articles are liked by the user
	for i := 0; i < len(articlesDb); i++ {
		articleLikes[i].UserId = uint(userId)
		articleLikes[i].ArticleID = articlesDb[i].ID
		err = repository.db.Model(&articleLikes[i]).Where("user_id = ? AND article_id = ?", articleLikes[i].UserId, articleLikes[i].ArticleID).Count(&counter).Error

		if err != nil {
			return nil, err
		}

		if counter > 0 {
			isLiked = append(isLiked, true)
		} else {
			isLiked = append(isLiked, false)
		}

		counter = 0
	}

	articlesEnt := make([]articleEntities.Article, len(articlesDb))
	for i := 0; i < len(articlesDb); i++ {
		articlesEnt[i] = articleEntities.Article{
			ID:        articlesDb[i].ID,
			Title:     articlesDb[i].Title,
			Content:   articlesDb[i].Content,
			Date:      articlesDb[i].Date,
			ImageUrl:  articlesDb[i].ImageUrl,
			ViewCount: articlesDb[i].ViewCount,
			DoctorID:  articlesDb[i].DoctorID,
			Doctor: doctorEntities.Doctor{
				ID:   articlesDb[i].Doctor.ID,
				Name: articlesDb[i].Doctor.Name,
			},
			IsLiked: isLiked[i],
		}
	}

	return articlesEnt, nil
}

func (repository *ArticleRepo) GetArticleById(articleId int, userId int) (articleEntities.Article, error) {
	var articleDb Article
	err := repository.db.Where("id = ?", articleId).Preload("Doctor").First(&articleDb).Error
	if err != nil {
		return articleEntities.Article{}, constants.ErrDataNotFound
	}

	var articleLikes ArticleLikes
	var isLiked bool
	var counter int64

	err = repository.db.Model(&articleLikes).Where("user_id = ? AND article_id = ?", userId, articleId).Count(&counter).Error

	if err != nil {
		return articleEntities.Article{}, constants.ErrServer
	}

	if counter > 0 {
		isLiked = true
	} else {
		isLiked = false
	}

	articleResp := articleEntities.Article{
		ID:        articleDb.ID,
		Title:     articleDb.Title,
		Content:   articleDb.Content,
		Date:      articleDb.Date,
		ImageUrl:  articleDb.ImageUrl,
		ViewCount: articleDb.ViewCount,
		DoctorID:  articleDb.DoctorID,
		Doctor: doctorEntities.Doctor{
			ID:   articleDb.Doctor.ID,
			Name: articleDb.Doctor.Name,
		},
		IsLiked: isLiked,
	}

	return articleResp, nil
}

func (repository *ArticleRepo) GetLikedArticle(metadata entities.Metadata, userId int) ([]articleEntities.Article, error) {
	var articleLikesDb []ArticleLikes
	err := repository.db.Limit(metadata.Limit).Offset((metadata.Page-1)*metadata.Limit).Where("user_id = ?", userId).Find(&articleLikesDb).Error
	if err != nil {
		return nil, constants.ErrDataNotFound
	}

	var likedArticleIDs []int
	for i := 0; i < len(articleLikesDb); i++ {
		likedArticleIDs = append(likedArticleIDs, int(articleLikesDb[i].ArticleID))
	}

	var articlesDb []Article
	err = repository.db.Where("id IN ?", likedArticleIDs).Preload("Doctor").Find(&articlesDb).Error
	if err != nil {
		return nil, constants.ErrServer
	}

	articlesEnt := make([]articleEntities.Article, len(articlesDb))
	for i := 0; i < len(articlesDb); i++ {
		articlesEnt[i] = articleEntities.Article{
			ID:        articlesDb[i].ID,
			Title:     articlesDb[i].Title,
			Content:   articlesDb[i].Content,
			Date:      articlesDb[i].Date,
			ImageUrl:  articlesDb[i].ImageUrl,
			ViewCount: articlesDb[i].ViewCount,
			DoctorID:  articlesDb[i].DoctorID,
			Doctor: doctorEntities.Doctor{
				ID:   articlesDb[i].Doctor.ID,
				Name: articlesDb[i].Doctor.Name,
			},
			IsLiked: true,
		}
	}

	return articlesEnt, nil
}

func (repository *ArticleRepo) LikeArticle(articleId int, userId int) error {
	var articleLikes ArticleLikes

	err := repository.db.Where("user_id = ? AND article_id = ?", userId, articleId).First(&articleLikes).Error
	if err == nil {
		return constants.ErrAlreadyLiked
	}

	var article Article
	err = repository.db.Where("id = ?", articleId).First(&article).Error
	if err != nil {
		return constants.ErrServer
	}

	articleLikes = ArticleLikes{
		UserId:    uint(userId),
		ArticleID: uint(articleId),
	}

	err = repository.db.Create(&articleLikes).Error
	if err != nil {
		return constants.ErrServer
	}

	return nil
}

func (repository *ArticleRepo) GetArticleByIdForDoctor(articleId int) (articleEntities.Article, error) {
	var articleDb Article
	err := repository.db.Where("id = ?", articleId).First(&articleDb).Error
	if err != nil {
		return articleEntities.Article{}, constants.ErrDataNotFound
	}

	return articleEntities.Article{
		ID:        articleDb.ID,
		Title:     articleDb.Title,
		Content:   articleDb.Content,
		Date:      articleDb.Date,
		ImageUrl:  articleDb.ImageUrl,
		ViewCount: articleDb.ViewCount,
		DoctorID:  articleDb.DoctorID,
	}, nil
}

func (repository *ArticleRepo) GetAllArticleByDoctorId(metadata entities.MetadataFull, doctorId int) ([]articleEntities.Article, error) {
	var articleDb []Article

	query := repository.db.Where("doctor_id = ?", doctorId).Limit(metadata.Limit).Offset((metadata.Page - 1) * metadata.Limit).Order(metadata.Sort + " " + metadata.Order)

	if metadata.Search != "" {
		query = query.Where("title LIKE ?", "%"+metadata.Search+"%")
	}

	err := query.Find(&articleDb).Error
	if err != nil {
		return []articleEntities.Article{}, constants.ErrServer
	}

	articlesEnt := make([]articleEntities.Article, len(articleDb))
	for i := 0; i < len(articleDb); i++ {
		articlesEnt[i] = articleEntities.Article{
			ID:        articleDb[i].ID,
			Title:     articleDb[i].Title,
			Content:   articleDb[i].Content,
			Date:      articleDb[i].Date,
			ImageUrl:  articleDb[i].ImageUrl,
			ViewCount: articleDb[i].ViewCount,
			DoctorID:  articleDb[i].DoctorID,
		}
	}

	return articlesEnt, nil
}
