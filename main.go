package main

import (
	"capstone/configs"
	articleController "capstone/controllers/article"
	complaintController "capstone/controllers/complaint"
	consultationController "capstone/controllers/consultation"
	doctorController "capstone/controllers/doctor"
	moodController "capstone/controllers/mood"
	musicController "capstone/controllers/music"
	ratingController "capstone/controllers/rating"
	storyController "capstone/controllers/story"
	transactionController "capstone/controllers/transaction"
	userController "capstone/controllers/user"
	"capstone/repositories/mysql"
	articleRepositories "capstone/repositories/mysql/article"
	complaintRepositories "capstone/repositories/mysql/complaint"
	consultationRepositories "capstone/repositories/mysql/consultation"
	doctorRepositories "capstone/repositories/mysql/doctor"
	moodRepositories "capstone/repositories/mysql/mood"
	musicRepositories "capstone/repositories/mysql/music"
	ratingRepositories "capstone/repositories/mysql/rating"
	storyRepositories "capstone/repositories/mysql/story"
	transactionRepositories "capstone/repositories/mysql/transaction"
	userRepositories "capstone/repositories/mysql/user"
	"capstone/routes"
	articleUseCase "capstone/usecases/article"
	complaintUseCase "capstone/usecases/complaint"
	consultationUseCase "capstone/usecases/consultation"
	doctorUseCase "capstone/usecases/doctor"
	midtransUseCase "capstone/usecases/midtrans"
	moodUseCase "capstone/usecases/mood"
	musicUseCase "capstone/usecases/music"
	ratingUseCase "capstone/usecases/rating"
	storyUseCase "capstone/usecases/story"
	transactionUseCase "capstone/usecases/transaction"
	userUseCase "capstone/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())
	midtransConfig := configs.MidtransConfig()

	userRepo := userRepositories.NewUserRepo(db)
	doctorRepo := doctorRepositories.NewDoctorRepo(db)
	consultationRepo := consultationRepositories.NewConsultationRepo(db)
	storyRepo := storyRepositories.NewStoryRepo(db)
	complaintRepo := complaintRepositories.NewComplaintRepo(db)
	transactionRepo := transactionRepositories.NewTransactionRepo(db)
	musicRepo := musicRepositories.NewMusicRepo(db)
	ratingRepo := ratingRepositories.NewRatingRepo(db)
	moodRepo := moodRepositories.NewMoodRepo(db)
	articleRepo := articleRepositories.NewArticleRepo(db)

	userUC := userUseCase.NewUserUseCase(userRepo)
	doctorUC := doctorUseCase.NewDoctorUseCase(doctorRepo)
	consultationUC := consultationUseCase.NewConsultationUseCase(consultationRepo)
	storyUC := storyUseCase.NewStoryUseCase(storyRepo)
	complaintUC := complaintUseCase.NewComplaintUseCase(complaintRepo)
	midtransUC := midtransUseCase.NewMidtransUseCase(midtransConfig)
	transactionUC := transactionUseCase.NewTransactionUseCase(transactionRepo, midtransUC)
	musicUC := musicUseCase.NewMusicUseCase(musicRepo)
	ratingUC := ratingUseCase.NewRatingUseCase(ratingRepo)
	moodUC := moodUseCase.NewMoodUseCase(moodRepo)
	articleUC := articleUseCase.NewArticleUseCase(articleRepo)

	userCont := userController.NewUserController(userUC)
	doctorCont := doctorController.NewDoctorController(doctorUC)
	consultationCont := consultationController.NewConsultationController(consultationUC)
	storyCont := storyController.NewStoryController(storyUC)
	complaintCont := complaintController.NewComplaintController(complaintUC)
	transactionCont := transactionController.NewTransactionController(transactionUC)
	musicCont := musicController.NewMusicController(musicUC)
	ratingCont := ratingController.NewRatingController(ratingUC)
	moodCont := moodController.NewMoodController(moodUC)
	articleCont := articleController.NewArticleController(articleUC)

	route := routes.NewRoute(userCont, doctorCont, consultationCont, storyCont, complaintCont, transactionCont, musicCont, ratingCont, moodCont, articleCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
