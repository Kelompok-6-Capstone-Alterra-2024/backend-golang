package routes

import (
	"capstone/controllers/complaint"
	"capstone/controllers/consultation"
	"capstone/controllers/doctor"
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
	complaintController    *complaint.ComplaintController
	transactionController  *transaction.TransactionController
}

func NewRoute(
	userController *user.UserController,
	doctorController *doctor.DoctorController,
	consultationController *consultation.ConsultationController,
	complaintController *complaint.ComplaintController,
	transactionController *transaction.TransactionController,
) *RouteController {
	return &RouteController{
		userController:         userController,
		doctorController:       doctorController,
		consultationController: consultationController,
		complaintController:    complaintController,
		transactionController:  transactionController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)

	e.HTTPErrorHandler = base.ErrorHandler

	userAuth := e.Group("/v1/users")
	userAuth.POST("/register", r.userController.Register) //Register User
	userAuth.POST("/login", r.userController.Login)       //Login User

	userRoute := userAuth.Group("/")
	userRoute.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	// Doctor
	userRoute.GET("doctor/:id", r.doctorController.GetByID)         //Get Doctor By ID
	userRoute.GET("doctor", r.doctorController.GetAll)              //Get All Doctor
	userRoute.GET("doctor/available", r.doctorController.GetActive) //Get All Active Doctor

	// Consultation
	userRoute.POST("consultations", r.consultationController.CreateConsultation)     //Get All Consultation
	userRoute.GET("consultations/:id", r.consultationController.GetConsultationByID) //Get Consultation By ID
	userRoute.GET("consultations", r.consultationController.GetAllConsultation)      //Get All Consultation

	// Complaint
	userRoute.POST("complaint", r.complaintController.Create) // Create Complaint

	// Transaction
	userRoute.POST("transaction", r.transactionController.Insert)                               // Create Transaction
	userRoute.GET("transaction/:id", r.transactionController.FindByID)                          // Get Transaction By ID
	userRoute.GET("transaction/consultation/:id", r.transactionController.FindByConsultationID) // Get Transaction By Consultation ID
	userRoute.GET("transactions", r.transactionController.FindAll)                              // Get All Transaction

	doctorAuth := e.Group("/v1/doctors")

	doctorAuth.POST("/register", r.doctorController.Register) //Register Doctor
	doctorAuth.POST("/login", r.doctorController.Login)       //Login Doctor

}
