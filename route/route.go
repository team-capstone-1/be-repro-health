package route

import (
	"capstone-project/config"
	"capstone-project/constant"
	"capstone-project/controller"
	"capstone-project/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	m "capstone-project/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Trailing Slash for slashing in endpoint
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	//JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(config.JWT_KEY)))

	// davin
	// user auth route
	e.POST("/users/signup", controller.SignUpUserController)
	e.POST("/users/login", controller.LoginUserController)
	e.PUT("/users/send-otp", controller.SendOTP)
	e.PUT("/users/validate-otp", controller.ValidateOTP)
	r.PUT("/users/change-password", controller.ChangeUserPasswordController, m.CheckRole(constant.ROLE_USER))

	// user appointment route
	e.GET("/specialists", controller.GetSpecialistsController)
	e.GET("/clinics", controller.GetClinicsController)
	e.GET("/clinics/:id/specialists", controller.GetSpecialistsByClinicController)
	e.GET("/clinics/:clinic_id/specialists/:specialist_id/doctors", controller.GetDoctorsBySpecialistAndClinicController)
	e.GET("/specialists/:id/doctors", controller.GetDoctorsBySpecialistController)
	e.GET("/clinics/:id/doctors", controller.GetDoctorsByClinicController)
	e.GET("/doctors", controller.GetDoctorsController)
	e.GET("/doctors/:id", controller.GetDoctorController)
	r.POST("/consultations", controller.CreateConsultationController, m.CheckRole(constant.ROLE_USER))

	// patient route
	r.GET("/patients", controller.GetPatientsController, m.CheckRole(constant.ROLE_USER))
	r.GET("/patients/:id", controller.GetPatientController, m.CheckRole(constant.ROLE_USER))
	r.POST("/patients", controller.CreatePatientController, m.CheckRole(constant.ROLE_USER))
	r.PUT("/patients/:id", controller.UpdatePatientController, m.CheckRole(constant.ROLE_USER))
	r.DELETE("/patients/:id", controller.DeletePatientController, m.CheckRole(constant.ROLE_USER))

	// user forum
	e.GET("/forums", controller.GetForumsController)
	e.GET("/forums/:id", controller.GetForumController)
	r.POST("/forums", controller.CreateForumController, m.CheckRole(constant.ROLE_USER))
	r.DELETE("/forums/:id", controller.DeleteForumController, m.CheckRole(constant.ROLE_USER))

	// user article
	r.GET("/articles", controller.GetArticlesController, m.CheckRole(constant.ROLE_USER))
	r.GET("/articles/bookmarks", controller.GetBookmarkedArticlesController, m.CheckRole(constant.ROLE_USER))
	r.GET("/articles/:id", controller.GetArticleController, m.CheckRole(constant.ROLE_USER))
	r.POST("/articles/:id/comments", controller.CreateCommentController, m.CheckRole(constant.ROLE_USER))
	r.POST("/articles/:id/bookmarks", controller.BookmarkController, m.CheckRole(constant.ROLE_USER))

	// user ai
	aiController := controller.NewUserAIController(repository.NewUserAIRepository())
	r.POST("/chatbot/users-health-recommendation", aiController.UserGetHealthRecommendation, m.CheckRole(constant.ROLE_USER))
	r.GET("/chatbot/users/:id", aiController.GetHealthRecommendationHistory, m.CheckRole(constant.ROLE_USER))

	// transaction
	r.GET("/transactions/:id", controller.GetTransactionController, m.CheckRole(constant.ROLE_USER))
	r.GET("/transactions", controller.GetTransactionsController, m.CheckRole(constant.ROLE_USER))
	r.GET("/transactions/patients/:id", controller.GetPatientTransactionsController, m.CheckRole(constant.ROLE_USER))
	r.POST("/transactions/:id/payments", controller.CreatePaymentController, m.CheckRole(constant.ROLE_USER))
	r.PUT("/transactions/:id/reschedule", controller.RescheduleController, m.CheckRole(constant.ROLE_USER))
	r.POST("/transactions/:id/cancel", controller.CancelTransactionController, m.CheckRole(constant.ROLE_USER))
	r.PUT("/transactions/:id/payment-timeout", controller.PaymentTimeOut, m.CheckRole(constant.ROLE_USER))
	r.PUT("/refund/:id", controller.ValidateRefund, m.CheckRole(constant.ROLE_ADMIN))

	r.GET("/notifications/patients/:id", controller.GetNotificationsController, m.CheckRole(constant.ROLE_USER))
	// davin

	// ADMIN ROUTE
	e.POST("/admins/login", controller.AdminLoginController)
	adm := e.Group("/admins")
	adm.Use(middleware.JWT([]byte(config.JWT_KEY)))
	adm.POST("/doctors/signup", controller.SignUpDoctorController, m.CheckRole("admin"))
	// SPECIALIST
	adm.GET("/specialists", controller.GetSpecialistsController, m.CheckRole(constant.ROLE_ADMIN))
	adm.POST("/specialists", controller.CreateSpecialistController, m.CheckRole(constant.ROLE_ADMIN))
	adm.PUT("/specialists/:id", controller.UpdateSpecialistController, m.CheckRole(constant.ROLE_ADMIN))
	adm.DELETE("/specialists/:id", controller.DeleteSpecialistController, m.CheckRole(constant.ROLE_ADMIN))
	// WORK HISTORY
	adm.POST("/profile/work-history", controller.CreateDoctorWorkHistoryController, m.CheckRole(constant.ROLE_ADMIN))
	adm.PUT("/profile/work-history/:id", controller.UpdateDoctorWorkHistoryController, m.CheckRole(constant.ROLE_ADMIN))
	adm.DELETE("/profile/work-history/:id", controller.DeleteDoctorWorkHistoryController, m.CheckRole(constant.ROLE_ADMIN))
	// EDUCATION
	adm.POST("/profile/education", controller.CreateDoctorEducationController, m.CheckRole(constant.ROLE_ADMIN))
	adm.PUT("/profile/education/:id", controller.UpdateDoctorEducationController, m.CheckRole(constant.ROLE_ADMIN))
	adm.DELETE("/profile/education/:id", controller.DeleteDoctorEducationController, m.CheckRole(constant.ROLE_ADMIN))
	// CERTIFICATION
	adm.POST("/profile/certification", controller.CreateDoctorCertificationController, m.CheckRole(constant.ROLE_ADMIN))
	adm.PUT("/profile/certification/:id", controller.UpdateDoctorCertificationController, m.CheckRole(constant.ROLE_ADMIN))
	adm.DELETE("/profile/certification/:id", controller.DeleteDoctorCertificationController, m.CheckRole(constant.ROLE_ADMIN))

	// DOCTOR ROUTE
	e.POST("/doctors/login", controller.DoctorLoginController)
	e.PUT("/doctors/send-otp", controller.DoctorSendOTPController)
	e.PUT("/doctors/validate-otp", controller.DoctorValidateOTPController)
	doctor := e.Group("/doctors")
	doctor.Use(middleware.JWT([]byte(config.JWT_KEY)))
	doctor.GET("/profile", controller.GetDoctorProfileController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/profile/work-histories", controller.GetDoctorWorkHistoriesController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/profile/educations", controller.GetDoctorEducationController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/profile/certifications", controller.GetDoctorCertificationController, m.CheckRole(constant.ROLE_DOCTOR))
	// DOCTOR FORUM
	doctor.GET("/forums", controller.GetDoctorAllForumsController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/forums/details/:id", controller.GetDoctorForumDetails, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.POST("/forum-replies", controller.CreateDoctorReplyForum, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.PUT("/forum-replies/:id", controller.UpdateDoctorReplyForum, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/forum-replies/:id", controller.GetDoctorForumReplyID, m.CheckRole(constant.ROLE_ADMIN))
	doctor.DELETE("/forum-replies/:id", controller.DeleteDoctorForumReplyController, m.CheckRole(constant.ROLE_DOCTOR))

	// DOCTOR ARTICLE ROUTE
	doctor.GET("/articles", controller.GetAllArticleDoctorsController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/articles/:id", controller.GetDoctorArticleByIDController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.POST("/articles", controller.CreateDoctorArticleController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.PUT("/articles/:id", controller.UpdateDoctorArticleController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.DELETE("/articles/:id", controller.DeleteDoctorArticleController, m.CheckRole(constant.ROLE_DOCTOR))

	// Endpoint baru untuk mengupdate status Published
	doctor.PUT("/articles/:id/publish", controller.UpdateArticlePublishedStatusController, m.CheckRole(constant.ROLE_DOCTOR))

	// DOCTOR DASHBOARD
	doctor.GET("/dashboard/data-count-one-month", controller.GetDataCountForDoctorControllerOneMonth, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/dashboard/data-count-one-week", controller.GetDataCountForDoctorControllerOneWeek, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/dashboard/data-count-one-day", controller.GetDataCountForDoctorControllerOneDay, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/dashboard/graph", controller.GetGraphController, m.CheckRole(constant.ROLE_DOCTOR))
	// doctor.GET("/dashboard/calendar", controller.GetCalendarController, m.CheckRole(constant.ROLE_DOCTOR))

	// DOCTOR AI
	doctorAIController := controller.NewDoctorAIController(repository.NewDoctorAIRepository())
	doctor.POST("/chatbot/health-recommendation", doctorAIController.DoctorGetHealthRecommendation, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/chatbot/health-recommendation/:doctor_id", doctorAIController.GetHealthRecommendationDoctorHistory, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/chatbot/health-recommendation/session/:session_id", doctorAIController.GetHealthRecommendationDoctorHistoryFromSession, m.CheckRole(constant.ROLE_DOCTOR))
	// per session +

	// DOCTOR APPOINTMENT
	doctor.GET("/appointments/details-transaction/:id", controller.DoctorGetDetailsTransactionController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/appointments/details-consultation/:id", controller.DoctorGetDetailsPatientController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.GET("/appointments/details-consultation", controller.DoctorGetAllConsultations, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.PUT("/appointments/confirm-consultation", controller.DoctorConfirmConsultationController, m.CheckRole(constant.ROLE_DOCTOR))
	doctor.PUT("/appointments/finish-consultation", controller.DoctorFinishedConsultationController, m.CheckRole(constant.ROLE_DOCTOR))

	// DOCTOR CHANGE PASSWORD
	doctor.PUT("/change-password", controller.ChangeDoctorPasswordController)

	// DOCTOR SCHEDULE
	doctor.GET("/schedule", controller.GetAllDoctorScheduleController, m.CheckRole(constant.ROLE_DOCTOR))

	return e
}