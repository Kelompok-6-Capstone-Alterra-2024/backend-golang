package routes

import (
	"capstone/controllers/complaint"
	"capstone/controllers/consultation"
	"capstone/controllers/doctor"
	"capstone/controllers/mood"
	"capstone/controllers/music"
	"capstone/controllers/rating"
	"capstone/controllers/story"
	"capstone/controllers/transaction"
	"capstone/controllers/user"
	myMiddleware "capstone/middlewares"
	"capstone/utilities/base"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
)

type RouteController struct {
	userController         *user.UserController
	doctorController       *doctor.DoctorController
	consultationController *consultation.ConsultationController
	storyController        *story.StoryController
	complaintController    *complaint.ComplaintController
	transactionController  *transaction.TransactionController
	musicController        *music.MusicController
	ratingController       *rating.RatingController
	moodController         *mood.MoodController
}

func NewRoute(
	userController *user.UserController,
	doctorController *doctor.DoctorController,
	consultationController *consultation.ConsultationController,
	storyContoller *story.StoryController,
	complaintController *complaint.ComplaintController,
	transactionController *transaction.TransactionController,
	musicController *music.MusicController,
	ratingController *rating.RatingController,
	moodController *mood.MoodController,
) *RouteController {
	return &RouteController{
		userController:         userController,
		doctorController:       doctorController,
		consultationController: consultationController,
		storyController:        storyContoller,
		complaintController:    complaintController,
		transactionController:  transactionController,
		musicController:        musicController,
		ratingController:       ratingController,
		moodController:         moodController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)

	e.HTTPErrorHandler = base.ErrorHandler
	e.Use(myMiddleware.CORSMiddleware())

	userAuth := e.Group("/v1/users")
	userAuth.POST("/register", r.userController.Register) //Register User
	userAuth.POST("/login", r.userController.Login)       //Login User

	userRoute := userAuth.Group("/")
	userRoute.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	// Doctor
	userRoute.GET("doctors/:id", r.doctorController.GetByID)         //Get Doctor By ID
	userRoute.GET("doctors", r.doctorController.GetAll)              //Get All Doctor
	userRoute.GET("doctors/available", r.doctorController.GetActive) //Get All Active Doctor

	// Consultation
	userRoute.POST("consultations", r.consultationController.CreateConsultation)     //Create Consultation
	userRoute.GET("consultations/:id", r.consultationController.GetConsultationByID) //Get Consultation By ID
	userRoute.GET("consultations", r.consultationController.GetAllConsultation)      //Get All Consultation

	// Inspirational Stories
	userRoute.GET("stories", r.storyController.GetAllStories)         //Get All Stories
	userRoute.GET("stories/:id", r.storyController.GetStoryById)      //Get Story By ID
	userRoute.GET("stories/liked", r.storyController.GetLikedStories) //Get Liked Stories
	userRoute.POST("stories/like", r.storyController.LikeStory)       //Like Story

	// Music
	userRoute.GET("musics", r.musicController.GetAllMusics)         //Get All Music
	userRoute.GET("musics/:id", r.musicController.GetMusicByID)     //Get Music By ID
	userRoute.GET("musics/liked", r.musicController.GetLikedMusics) //Get Liked Music
	userRoute.POST("musics/like", r.musicController.LikeMusic)      //Like Music

	// Complaint
	userRoute.POST("complaint", r.complaintController.Create) // Create Complaint

	// Transaction
	userRoute.POST("transaction", r.transactionController.InsertWithBuiltIn)                    // Create Transaction
	userRoute.GET("transaction/:id", r.transactionController.FindByID)                          // Get Transaction By ID
	userRoute.GET("transaction/consultation/:id", r.transactionController.FindByConsultationID) // Get Transaction By Consultation ID
	userRoute.GET("transactions", r.transactionController.FindAll)                              // Get All Transaction
	userRoute.POST("transaction/bank-transfer", r.transactionController.BankTransfer)           // Bank Transfer

	// Rating
	userRoute.POST("feedbacks", r.ratingController.SendFeedback) // Create Rating

	// Mood
	userRoute.POST("moods", r.moodController.CreateMood)     // Create Mood
	userRoute.GET("moods", r.moodController.GetAllMoods)     // Get All Moods
	userRoute.GET("moods/:id", r.moodController.GetMoodById) // Get Mood By ID

	doctorAuth := e.Group("/v1/doctors")

	doctorAuth.POST("/register", r.doctorController.Register) //Register Doctor
	doctorAuth.POST("/login", r.doctorController.Login)       //Login Doctor

}
