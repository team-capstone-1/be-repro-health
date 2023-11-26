package route

import (
	"capstone-project/config"
	"capstone-project/controller"

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
	e.POST("/users/change-password", controller.ChangeUserPasswordController)

	// user appointment route
	e.GET("/clinics", controller.GetClinicsController)
	e.GET("/specialists/:id/doctors", controller.GetDoctorsBySpecialistController)
	e.GET("/clinics/:id/doctors", controller.GetDoctorsByClinicController)
	e.GET("/doctors", controller.GetDoctorsController)
	e.GET("/doctors/:id", controller.GetDoctorController)
	r.POST("/consultations", controller.CreateConsultationController)

	// patient route
	r.GET("/patients", controller.GetPatientsController, m.CheckRole("user"))
	r.GET("/patients/:id", controller.GetPatientController, m.CheckRole("user"))
	r.POST("/patients", controller.CreatePatientController, m.CheckRole("user"))
	r.PUT("/patients/:id", controller.UpdatePatientController, m.CheckRole("user"))
	r.DELETE("/patients/:id", controller.DeletePatientController, m.CheckRole("user"))
	
	// user forum
	e.GET("/forums", controller.GetForumsController)
	r.POST("/forums", controller.CreateForumController)
	r.DELETE("/forums/:id", controller.DeleteForumController)
	
	// transaction
	r.GET("/transactions/:id", controller.GetTransactionController, m.CheckRole("user"))
	// davin

	// admin route
	e.POST("/admins/login", controller.AdminLoginController)
	adm := e.Group("/admins")
	adm.Use(middleware.JWT([]byte(config.JWT_KEY)))
	adm.POST("/doctors/signup", controller.SignUpDoctorController, m.CheckRole("admin"))

	// doctor route
	e.POST("/doctors/login", controller.DoctorLoginController)

	// doctor route
	e.POST("/doctors/login", controller.DoctorLoginController)
	doctor := e.Group("/doctors")
	doctor.Use(middleware.JWT([]byte(config.JWT_KEY)))
	doctor.GET("/profile", controller.GetDoctorProfileController, m.CheckRole("doctor"))
	// doctor work history
	doctor.GET("/profile/work-histories", controller.GetDoctorWorkHistoriesController, m.CheckRole("doctor"))
	doctor.POST("/profile/work-history", controller.CreateDoctorWorkHistoryController, m.CheckRole("doctor"))
	doctor.PUT("/profile/work-history/:id", controller.UpdateDoctorWorkHistoryController, m.CheckRole("doctor"))
	doctor.DELETE("/profile/work-history/:id", controller.DeleteDoctorWorkHistoryController, m.CheckRole("doctor"))
	// doctor education
	doctor.GET("/profile/educations", controller.GetDoctorEducationController, m.CheckRole("doctor"))
	doctor.POST("/profile/education", controller.CreateDoctorEducationController, m.CheckRole("doctor"))
	doctor.PUT("/profile/education/:id", controller.UpdateDoctorEducationController, m.CheckRole("doctor"))
	doctor.DELETE("/profile/education/:id", controller.DeleteDoctorEducationController, m.CheckRole("doctor"))
	// doctor certification
	doctor.GET("/profile/certifications", controller.GetDoctorCertificationController, m.CheckRole("doctor"))
	doctor.POST("/profile/certification", controller.CreateDoctorCertificationController, m.CheckRole("doctor"))
	doctor.PUT("/profile/certification/:id", controller.UpdateDoctorCertificationController, m.CheckRole("doctor"))
	doctor.DELETE("/profile/certification/:id", controller.DeleteDoctorCertificationController, m.CheckRole("doctor"))
	
	// doctor specialist
	doctor.GET("/specialists", controller.GetSpecialistsController, m.CheckRole("doctor"))
	doctor.POST("/specialist", controller.CreateSpecialistController, m.CheckRole("doctor"))
	doctor.PUT("/specialist/:id", controller.UpdateSpecialistController, m.CheckRole("doctor"))
	doctor.DELETE("/specialist/:id", controller.DeleteSpecialistController, m.CheckRole("doctor"))

	// doctor article route
	doctor.GET("/articles", controller.GetAllArticleDoctorsController, m.CheckRole("doctor"))
	doctor.POST("/articles", controller.CreateDoctorArticleController, m.CheckRole("doctor"))
	doctor.DELETE("/articles/:id", controller.DeleteDoctorArticleController, m.CheckRole("doctor"))

	// doctor dashboard
	doctor.GET("/consultations-dashboard", controller.GetConsultationSchedulesForDoctorDashboardController, m.CheckRole("doctor"))
	doctor.GET("/patients-dashboard", controller.GetPatientsForDoctorDashboardController)
	doctor.GET("/transactions-dashboard", controller.GetTransactionsForDoctorDashboardController)
	doctor.GET("/articles-dashboard", controller.GetArticleForDoctorDashboardController)

	return e
}
