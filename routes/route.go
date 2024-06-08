package routes

import (
	"capstone/controllers/article"
	"capstone/controllers/chatbot"
	"capstone/controllers/complaint"
	"capstone/controllers/consultation"
	"capstone/controllers/doctor"
	"capstone/controllers/forum"
	"capstone/controllers/mood"
	"capstone/controllers/music"
	"capstone/controllers/post"
	"capstone/controllers/rating"
	"capstone/controllers/story"
	"capstone/controllers/transaction"
	"capstone/controllers/user"
	myMiddleware "capstone/middlewares"
	"capstone/utilities/base"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
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
	forumController        *forum.ForumController
	postController         *post.PostController
	chatbotController      *chatbot.ChatbotController
	articleController      *article.ArticleController
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
	forumController *forum.ForumController,
	postController *post.PostController,
	chatbotController *chatbot.ChatbotController,
	articleController *article.ArticleController) *RouteController {
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
		forumController:        forumController,
		postController:         postController,
		chatbotController:      chatbotController,
		articleController:      articleController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)

	e.HTTPErrorHandler = base.ErrorHandler
	e.Use(myMiddleware.CORSMiddleware())

	e.POST("/v1/payment-callback", r.transactionController.CallbackTransaction)

	// chatbot
	e.GET("/v1/users/chatbots/customer-service", r.chatbotController.ChatbotCS)        //customer service chatbot
	e.GET("/v1/users/chatbots/mental-health", r.chatbotController.ChatbotMentalHealth) //mental health chatbot
	e.GET("/v1/doctors/chatbots/treatment", r.chatbotController.ChatbotTreatment) //Chatbot Treatment

	userAuth := e.Group("/v1/users")
	userAuth.POST("/register", r.userController.Register) //Register User
	userAuth.POST("/login", r.userController.Login)       //Login User

	userAuth.GET("/auth/google/login", r.userController.GoogleLogin)
	userAuth.GET("/auth/google/callback", r.userController.GoogleCallback)

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
	userRoute.POST("payments/gateway", r.transactionController.InsertWithBuiltIn)               // Create Transaction
	userRoute.GET("transaction/:id", r.transactionController.FindByID)                          // Get Transaction By ID
	userRoute.GET("transaction/consultation/:id", r.transactionController.FindByConsultationID) // Get Transaction By Consultation ID
	userRoute.GET("transactions", r.transactionController.FindAll)                              // Get All Transaction
	userRoute.POST("payments/bank-transfer", r.transactionController.BankTransfer)              // Bank Transfer
	userRoute.POST("payments/e-wallet", r.transactionController.EWallet)                        // E-Wallet

	// Rating
	userRoute.POST("feedbacks", r.ratingController.SendFeedback) // Create Rating

	// Mood
	userRoute.POST("moods", r.moodController.CreateMood)     // Create Mood
	userRoute.GET("moods", r.moodController.GetAllMoods)     // Get All Moods
	userRoute.GET("moods/:id", r.moodController.GetMoodById) // Get Mood By ID

	// Forum
	userRoute.POST("forums/join", r.forumController.JoinForum)                       // Join Forum
	userRoute.DELETE("forums/:id", r.forumController.LeaveForum)                     // Leave Forum
	userRoute.GET("forums", r.forumController.GetJoinedForum)                        // Get All Forum
	userRoute.GET("forums/recommendation", r.forumController.GetRecommendationForum) // Get Recommendation Forum
	userRoute.GET("forums/:id", r.forumController.GetForumById)                      // Get Forum By ID

	// Posts
	userRoute.GET("forums/:forumId/posts", r.postController.GetAllPostsByForumId)   // Get All Posts By Forum ID
	userRoute.GET("posts/:id", r.postController.GetPostById)                        // Get Post By ID
	userRoute.POST("posts", r.postController.SendPost)                              // Create Post
	userRoute.POST("posts/like", r.postController.LikePost)                         // Like Post
	userRoute.POST("comments", r.postController.SendComment)                        // Create Comment
	userRoute.GET("posts/:postId/comments", r.postController.GetAllCommentByPostId) // Get All Comment By Post ID

	// Article
	userRoute.GET("articles", r.articleController.GetAllArticle) // Get All Article
	userRoute.GET("articles/:id", r.articleController.GetArticleById)
	userRoute.GET("articles/liked", r.articleController.GetLikedArticle)

	// consultation notes
	userRoute.GET("consultation-notes/consultation/:id", r.consultationController.GetConsultationNotesByID) // Get Consultation Note By ID

	doctorAuth := e.Group("/v1/doctors")

	doctorAuth.POST("/register", r.doctorController.Register)                  //Register Doctor
	doctorAuth.POST("/login", r.doctorController.Login)                        //Login Doctor
	doctorAuth.GET("/auth/google/login", r.doctorController.GoogleLogin)       // Google Login
	doctorAuth.GET("/auth/google/callback", r.doctorController.GoogleCallback) // Google Callback

	doctorRoute := doctorAuth.Group("/")
	doctorRoute.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	// articles
	doctorRoute.POST("articles", r.articleController.CreateArticle) // Create Article
	doctorRoute.GET("articles", r.articleController.GetAllArticle)  // Get All Article

	// musics
	doctorRoute.POST("musics", r.musicController.PostMusic)                    // Post Music
	doctorRoute.GET("musics", r.musicController.GetAllMusicsByDoctorId)         // Get All Music By Doctor ID
	doctorRoute.GET("musics/:id", r.musicController.GetMusicByIdForDoctor)              // Get Music By ID
	doctorRoute.GET("musics/count", r.musicController.CountMusicByDoctorId) // Count Music By Doctor ID
	doctorRoute.GET("musics/like/count", r.musicController.CountMusicLikesByDoctorId) // Count Music Likes By Doctor ID
	doctorRoute.GET("musics/view/count", r.musicController.CountMusicViewCountByDoctorId) // Count Music View Count By Doctor ID
	doctorRoute.PUT("musics/:id", r.musicController.EditMusic)              // Update Music
	doctorRoute.DELETE("musics/:id", r.musicController.DeleteMusic)          // Delete Music

	doctorRoute.POST("stories", r.storyController.PostStory)                    // Post Story
	doctorRoute.GET("stories", r.storyController.GetAllStoriesByDoctorId)         // Get All Story By Doctor ID
	doctorRoute.GET("stories/:id", r.storyController.GetStoryByIdForDoctor)         // Get Story By ID
	doctorRoute.GET("stories/count", r.storyController.CountStoriesByDoctorId) // Count Stories By Doctor ID
	doctorRoute.GET("stories/like/count", r.storyController.CountStoryLikesByDoctorId) // Count Stories Likes By Doctor ID
	doctorRoute.GET("stories/view/count", r.storyController.CountStoryViewByDoctorId) // Count Stories View Count By Doctor ID
	doctorRoute.PUT("stories/:id", r.storyController.EditStory)              // Update Story
	doctorRoute.DELETE("stories/:id", r.storyController.DeleteStory)          // Delete Story

	// consultation notes
	doctorRoute.POST("consultation-notes", r.consultationController.CreateConsultationNotes) // Post Consultation Note

	// Rating
	doctorRoute.GET("feedbacks", r.ratingController.GetAllFeedbacks) // Get All Feedbacks
}