package routes

import (
	//"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/lokesh2201013/middleware"
	"github.com/lokesh2201013/controllers"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)

	app.Use(middleware.AuthMiddleware())

    
    //  routes for Admins
    app.Post("/admin/job", middleware.AdminOnly(controllers.CreateJob))
    app.Get("/admin/job/:job_id", middleware.AdminOnly(controllers.GetJob))
    app.Get("/admin/applicants", middleware.AdminOnly(controllers.GetAllApplicants))
    app.Get("/admin/applicant/:applicant_id", middleware.AdminOnly(controllers.GetApplicantDetails))

    // applicant-related routes
    app.Post("/uploadResume", middleware.ApplicantOnly(controllers.UploadResume))
    app.Get("/jobs", controllers.ListJobs)
    app.Get("/jobs/apply", middleware.ApplicantOnly(controllers.ApplyForJob))

}
